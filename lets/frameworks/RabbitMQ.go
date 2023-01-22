package frameworks

import (
	"encoding/json"
	"fmt"
	"lets-go-framework/lets"
	"lets-go-framework/lets/rabbitmq"
	"lets-go-framework/lets/types"
	"os"

	"github.com/rabbitmq/amqp091-go"
)

// Initialize RabbitMQ Configuration.
var RabbitMQConfig types.IRabbitMQConfig

// RabbitMQ server defirinitions.
type rabbitServer struct {
	dsn    string
	config amqp091.Config
}

// Initialize RabbitMQ server.
func (r *rabbitServer) init(config types.IRabbitMQServer) {
	r.dsn = fmt.Sprintf("amqp://%s:%s@%s:%s", config.GetUsername(), config.GetPassword(), config.GetHost(), config.GetPort())
	r.config = amqp091.Config{Properties: amqp091.NewConnectionProperties()}
	r.config.Properties.SetClientConnectionName(os.Getenv("SERVICE_ID"))
}

// RabbitMQ consumer definitions.
type rabbitConsumer struct {
	connection *amqp091.Connection
	channel    *amqp091.Channel
	queue      amqp091.Queue
	deliveries <-chan amqp091.Delivery
	done       chan error
	engine     types.Engine
}

// Start consuming.
func (r *rabbitConsumer) consume(server *rabbitServer, consumer types.IRabbitMQConsumer) {
	var err error
	var dsn = fmt.Sprintf("%s/%s", server.dsn, consumer.GetVHost())

	r.connection, err = amqp091.DialConfig(dsn, server.config)
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

	// Declare (or using existing) exchange.
	if err = r.channel.ExchangeDeclare(
		consumer.GetExchange(),     // name of the exchange
		consumer.GetExchangeType(), // type
		true,                       // durable
		false,                      // delete when complete
		false,                      // internal
		false,                      // noWait
		nil,                        // arguments
	); err != nil {
		lets.LogE("RabbitMQ Server: %s", err.Error())
		return
	}

	// Declare (or using existing) queue.
	if r.queue, err = r.channel.QueueDeclare(
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
	if err = r.channel.QueueBind(
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
	if r.deliveries, err = r.channel.Consume(
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
	// var verbose = flag.Bool("verbose", true, "enable verbose output of message data")
	// var autoAck = flag.Bool("auto_ack", false, "enable message auto-ack")

	// Waiting message
	for delivery := range r.deliveries {
		if consumer.GetDebug() {
			deliveryCount++
			lets.LogD("RabbitMQ Server: message no.: %d", deliveryCount)
			lets.LogD("RabbitMQ Server: %d Byte delivery: [%v] %q", len(delivery.Body), delivery.DeliveryTag, delivery.Body)
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
			ReplyTo:       replyTo,
			CorrelationId: delivery.CorrelationId,
			Exchange:      delivery.Exchange,
			RoutingKey:    delivery.RoutingKey,
		}

		// Call event handler.
		r.engine.Call(event.Name, &event)

		delivery.Ack(false)
	}
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
					rc.consume(&rs, consumer)
				}
			}
		}
	}
}
