package frameworks

import (
	"context"
	"encoding/json"
	"fmt"
	"lets-go-framework/lets"
	"lets-go-framework/lets/rabbitmq"
	"lets-go-framework/lets/types"
	"os"
	"time"

	"github.com/rabbitmq/amqp091-go"
)

// Initialize RabbitMQ Configuration.
var RabbitMQConfig types.IRabbitMQConfig

// RabbitMQ server defirinitions.
type rabbitServer struct {
	dsn        string
	config     amqp091.Config
	connection *amqp091.Connection
	channel    *amqp091.Channel
}

// Initialize RabbitMQ server.
func (r *rabbitServer) init(config types.IRabbitMQServer) {
	r.dsn = fmt.Sprintf("amqp://%s:%s@%s:%s/%s", config.GetUsername(), config.GetPassword(), config.GetHost(), config.GetPort(), config.GetVHost())
	r.config = amqp091.Config{Properties: amqp091.NewConnectionProperties()}
	r.config.Properties.SetClientConnectionName(os.Getenv("SERVICE_ID"))
}

// Start consuming.
func (r *rabbitServer) connect() {
	var err error
	r.connection, err = amqp091.DialConfig(r.dsn, r.config)
	if err != nil {
		fmt.Printf("dial: %s", err.Error())
		return
	}

	// Listen for error on connection
	go func() {
		lets.LogE("RabbitMQ Server: %s", <-r.connection.NotifyClose(make(chan *amqp091.Error)))
	}()

	// Create channel connection.
	if r.channel, err = r.connection.Channel(); err != nil {
		lets.LogE("RabbitMQ Server: %s", err.Error())
		return
	}
}

// RabbitMQ consumer definitions.
type rabbitConsumer struct {
	queue      amqp091.Queue
	deliveries <-chan amqp091.Delivery
	done       chan error
	engine     types.Engine
}

// Start consuming.
func (r *rabbitConsumer) consume(server *rabbitServer, consumer types.IRabbitMQConsumer) {
	var err error

	// Create channel connection.
	if server.channel, err = server.connection.Channel(); err != nil {
		lets.LogE("RabbitMQ Server: %s", err.Error())
		return
	}

	// Declare (or using existing) queue.
	if r.queue, err = server.channel.QueueDeclare(
		consumer.GetQueue(), // name of the queue
		true,                // durable
		false,               // delete when unused
		false,               // exclusive
		false,               // noWait
		nil,                 // arguments
	); err != nil {
		lets.LogE("RabbitMQ Server: %s", err.Error())
		return
	}

	// Bind queue to exchange.
	if err = server.channel.QueueBind(
		r.queue.Name,             // name of the queue
		consumer.GetRoutingKey(), // routing key
		consumer.GetExchange(),   // sourceExchange
		false,                    // noWait
		nil,                      // arguments
	); err != nil {
		lets.LogE("RabbitMQ Server: %s", err.Error())
		return
	}

	// Consume message
	if r.deliveries, err = server.channel.Consume(
		r.queue.Name,       // name
		consumer.GetName(), // consumerTag,
		false,              // autoAck
		false,              // exclusive
		false,              // noLocal
		false,              // noWait
		nil,                // arguments
	); err != nil {
		fmt.Printf("consume: %s", err.Error())
		return
	}

	cleanup := func() {
		lets.LogE("RabbitMQ Server: %s", "Delivery channel is closed.")
		r.done <- nil
	}
	defer cleanup()

	var deliveryCount uint64 = 0

	// Waiting message
	for delivery := range r.deliveries {
		if consumer.GetDebug() {
			deliveryCount++
			lets.LogD("RabbitMQ Server: %d Byte delivery: [%v] \n%q", len(delivery.Body), delivery.DeliveryTag, delivery.Body)
		}

		// Bind body into types.RabbitBody.
		var body types.RabbitBody
		err := json.Unmarshal(delivery.Body, &body)
		if err != nil {
			lets.LogE("RabbitMQ Server: ", err.Error())
			continue
		}

		// Read reply to
		var replyTo types.ReplyTo
		json.Unmarshal([]byte(delivery.ReplyTo), &replyTo)

		// Create event data.
		event := types.Event{
			Name:          body.Event,
			Data:          body.Data,
			ReplyTo:       &replyTo,
			CorrelationId: delivery.CorrelationId,
			Exchange:      delivery.Exchange,
			RoutingKey:    delivery.RoutingKey,
		}

		// Call event handler.
		r.engine.Call(event.Name, &event)

		delivery.Ack(false)
	}
}

// RabbitMQ consumer definitions.
type RabbitPublisher struct {
	channel *amqp091.Channel
	name    string
}

func (r *RabbitPublisher) init(server *rabbitServer, publisher types.IRabbitMQPublisher) {
	r.channel = server.channel
	r.name = publisher.GetName()
}

func (r *RabbitPublisher) Publish(event types.IEvent) (err error) {
	var body = event.GetBody()
	// Encode object to json string
	if event.GetDebug() {
		seqNo := r.channel.GetNextPublishSeqNo()
		lets.LogD("RabbitMQ Publisher: to: exchange '%s'; key: '%s'", event.GetExchange(), event.GetRoutingKey())
		lets.LogD("RabbitMQ Publisher: sequence no: %d; %d Bytes; Body: \n%s", seqNo, len(body), string(body))
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err = r.channel.PublishWithContext(ctx,
		event.GetExchange(),   // Exchange
		event.GetRoutingKey(), // RoutingKey or queues
		false,                 // Mandatory
		false,                 // Immediate
		amqp091.Publishing{
			Headers:         amqp091.Table{},
			ContentType:     "application/json",
			ContentEncoding: "",
			Body:            []byte(body),
			DeliveryMode:    amqp091.Transient, // 1=non-persistent, 2=persistent
			Priority:        0,                 // 0-9
			// a bunch of application/implementation-specific fields
			ReplyTo:       event.GetReplyTo().GetJson(),
			CorrelationId: event.GetCorrelationId(),
		},
	); err != nil {
		lets.LogE("RabbitMQ Publisher: %s", err.Error())
		return
	}

	return
}

// Define rabbit service host and port
func RabbitMQ() {
	if RabbitMQConfig == nil {
		return
	}

	// Running RabbitMQ
	if servers := RabbitMQConfig.GetServers(); len(servers) != 0 {
		lets.LogI("RabbitMQ Starting ...")

		for _, server := range servers {
			var rs rabbitServer
			rs.init(server)
			rs.connect()

			// Consuming RabbitMQ.
			if consumers := server.GetConsumers(); len(consumers) != 0 {
				lets.LogI("RabbitMQ Consumer Starting ...")

				for _, consumer := range consumers {
					var rc = rabbitConsumer{
						engine: &rabbitmq.Engine{
							Debug: consumer.GetDebug(),
						},
					}

					consumer.GetListener()(rc.engine)
					go rc.consume(&rs, consumer)
				}
			}

			if publishers := server.GetPublishers(); len(publishers) != 0 {
				lets.LogI("RabbitMQ Publisher Starting ...")
				for _, publisher := range publishers {
					var rp RabbitPublisher
					rp.init(&rs, publisher)

					lets.LogI("RabbitMQ Publisher: %s", publisher.GetName())
					for _, client := range publisher.GetClients() {
						client.SetConnection(&rp)
					}
				}
			}
		}
	}
}
