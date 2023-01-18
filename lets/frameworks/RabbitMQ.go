package frameworks

import (
	"encoding/json"
	"flag"
	"fmt"
	"lets-go-framework/lets/rabbitmq"
	"lets-go-framework/lets/types"
	"log"

	"github.com/rabbitmq/amqp091-go"
)

var RabbitMQDsn types.IRabbitMQDsn
var RabbitMQConsumer types.IRabbitMQConsumer

// HTTP service struct
type rabbitService struct {
	Dsn        string
	Engine     rabbitmq.Engine
	Config     amqp091.Config
	Connection *amqp091.Connection
	Channel    *amqp091.Channel
	deliveries <-chan amqp091.Delivery
	done       chan error
}

func (rabbit *rabbitService) Init() {
	fmt.Println("ServiceRabbit.Init()")
	rabbit.Dsn = fmt.Sprintf("amqp://%s:%s@%s:%s/%s", RabbitMQDsn.GetUsername(), RabbitMQDsn.GetPassword(), RabbitMQDsn.GetHost(), RabbitMQDsn.GetPort(), RabbitMQDsn.GetVirtualHost())

	rabbit.Config = amqp091.Config{Properties: amqp091.NewConnectionProperties()}
	rabbit.Config.Properties.SetClientConnectionName(RabbitMQConsumer.GetName())
}

func (rabbit *rabbitService) Serve() {
	fmt.Println("ServiceRabbit.Connect()")

	var err error
	rabbit.Connection, err = amqp091.DialConfig(rabbit.Dsn, rabbit.Config)
	if err != nil {
		fmt.Printf("dial: %s", err.Error())

		return
	}

	// Listen for error on connection
	go func() {
		log.Printf("closing: %s", <-rabbit.Connection.NotifyClose(make(chan *amqp091.Error)))
	}()

	/////////////
	// CHANNEL //
	/////////////

	// Create channel connection
	rabbit.Channel, err = rabbit.Connection.Channel()
	if err != nil {
		fmt.Printf("channel: %s", err.Error())
		return
	}

	// log.Println("got Channel")

	// log.Printf("declaring Exchange (%q)", RabbitMConsumer.GetExchange())

	if err := rabbit.Channel.ExchangeDeclare(
		RabbitMQConsumer.GetExchange(),     // name of the exchange
		RabbitMQConsumer.GetExchangeType(), // type
		true,                               // durable
		false,                              // delete when complete
		false,                              // internal
		false,                              // noWait
		nil,                                // arguments
	); err != nil {
		fmt.Printf("exchange: %s", err.Error())
		return
	}

	// log.Printf("declared Exchange, declaring Queue %q", RabbitMConsumer.GetQueue())

	queue, err := rabbit.Channel.QueueDeclare(
		RabbitMQConsumer.GetQueue(), // name of the queue
		true,                        // durable
		false,                       // delete when unused
		false,                       // exclusive
		false,                       // noWait
		nil,                         // arguments
	)
	if err != nil {
		fmt.Printf("queue: %s", err.Error())
		return
	}

	log.Printf("declared Queue (%q %d messages, %d consumers), binding to Exchange (key %q)",
		queue.Name, queue.Messages, queue.Consumers, RabbitMQConsumer.GetRoutingKey())

	if err = rabbit.Channel.QueueBind(
		queue.Name,                       // name of the queue
		RabbitMQConsumer.GetRoutingKey(), // bindingKey
		RabbitMQConsumer.GetExchange(),   // sourceExchange
		false,                            // noWait
		nil,                              // arguments
	); err != nil {
		fmt.Printf("bind: %s", err.Error())
		return
	}

	// log.Printf("Queue bound to Exchange, starting Consume (consumer tag %q)", c.GetName())
	rabbit.deliveries, err = rabbit.Channel.Consume(
		queue.Name,                 // name
		RabbitMQConsumer.GetName(), // consumerTag,
		false,                      // autoAck
		false,                      // exclusive
		false,                      // noLocal
		false,                      // noWait
		nil,                        // arguments
	)
	if err != nil {
		fmt.Printf("consume: %s", err.Error())
		return
	}

	cleanup := func() {
		log.Printf("handle: deliveries channel closed")
		rabbit.done <- nil
	}

	defer cleanup()

	var deliveryCount int = 0
	var verbose = flag.Bool("verbose", true, "enable verbose output of message data")
	var autoAck = flag.Bool("auto_ack", false, "enable message auto-ack")

	for d := range rabbit.deliveries {
		deliveryCount++
		if *verbose {
			log.Printf(
				"got %d Byte delivery: [%v] %q",
				len(d.Body),
				d.DeliveryTag,
				d.Body,
			)
		} else {
			if deliveryCount%65536 == 0 {
				log.Printf("delivery count %d", deliveryCount)
			}
		}

		// rabbit.onMessage(&d)

		var body types.RabbitBody
		err := json.Unmarshal(d.Body, &body)
		if err != nil {
			log.Default().Println("Json Body Error: ", err.Error())
			return
		}

		data, err := json.Marshal(body.Data)
		if err != nil {
			log.Default().Println("Invalid data format", err.Error())
			return
		}

		var replyTo rabbitmq.ReplyTo
		json.Unmarshal([]byte(d.ReplyTo), &replyTo)

		event := rabbitmq.Event{
			Name:          body.Event,
			Data:          data,
			ReplyTo:       replyTo,
			CorrelationId: d.CorrelationId,
			Exchange:      d.Exchange,
			RoutingKey:    d.RoutingKey,
		}

		rabbit.Engine.Call(event.Name, &event)

		if !*autoAck {
			d.Ack(false)
		}
	}
}

// Define rabbit service host and port
func RabbitMQ() {
	fmt.Println("LoadRabbitFramework()")

	var rabbitService rabbitService

	rabbitService.Init()
	// services.RabbitEventHandler(&rabbitService.Engine)
	rabbitService.Serve()
}
