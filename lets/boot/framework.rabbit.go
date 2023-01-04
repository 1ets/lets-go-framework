package boot

import (
	"flag"
	"fmt"
	"lets-go-framework/configs"
	"lets-go-framework/lets/types"
	"log"

	"github.com/rabbitmq/amqp091-go"
)

type serviceRabbit struct {
	Dsn      types.IRabbitDsn
	Consumer types.IRabbitConsumer
	name     string
	dsn      string
	conn     *amqp091.Connection
	channel  *amqp091.Channel
	tag      string
	done     chan error
}

func (rabbit *serviceRabbit) Init() {
	fmt.Println("serviceRabbit.Init()")

	server := rabbit.Dsn
	rabbit.dsn = fmt.Sprintf("amqp://%s:%s@%s:%s/%s", server.GetUsername(), server.GetPassword(), server.GetHost(), server.GetPort(), server.GetVirtualHost())
}

func (rabbit *serviceRabbit) Connect() error {
	fmt.Println("serviceRabbit.Connect()")

	config := amqp091.Config{Properties: amqp091.NewConnectionProperties()}
	config.Properties.SetClientConnectionName(rabbit.name)

	log.Printf("dialing %q", rabbit.dsn)
	var err error
	rabbit.conn, err = amqp091.DialConfig(rabbit.dsn, config)
	if err != nil {
		return fmt.Errorf("dial: %s", err)
	}

	go func() {
		log.Printf("closing: %s", <-rabbit.conn.NotifyClose(make(chan *amqp091.Error)))
	}()

	log.Printf("got Connection, getting Channel")

	rabbit.channel, err = rabbit.conn.Channel()
	if err != nil {
		return fmt.Errorf("channel: %s", err)
	}

	log.Println("got Channel")

	return nil
}

func (rabbit *serviceRabbit) Register() error {
	fmt.Println("serviceRabbit.Register()")

	c := rabbit.Consumer
	log.Printf("declaring Exchange (%q)", c.GetExchange())

	if err := rabbit.channel.ExchangeDeclare(
		c.GetExchange(),     // name of the exchange
		c.GetExchangeType(), // type
		true,                // durable
		false,               // delete when complete
		false,               // internal
		false,               // noWait
		nil,                 // arguments
	); err != nil {
		return fmt.Errorf("Exchange Declare: %s", err)
	}

	log.Printf("declared Exchange, declaring Queue %q", c.GetQueue())
	queue, err := rabbit.channel.QueueDeclare(
		c.GetQueue(), // name of the queue
		true,         // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // noWait
		nil,          // arguments
	)
	if err != nil {
		return fmt.Errorf("Queue Declare: %s", err)
	}

	log.Printf("declared Queue (%q %d messages, %d consumers), binding to Exchange (key %q)",
		queue.Name, queue.Messages, queue.Consumers, c.GetRoutingKey())

	if err = rabbit.channel.QueueBind(
		queue.Name,        // name of the queue
		c.GetRoutingKey(), // bindingKey
		c.GetExchange(),   // sourceExchange
		false,             // noWait
		nil,               // arguments
	); err != nil {
		return fmt.Errorf("Queue Bind: %s", err)
	}

	log.Printf("Queue bound to Exchange, starting Consume (consumer tag %q)", c.GetName())
	deliveries, err := rabbit.channel.Consume(
		queue.Name,  // name
		c.GetName(), // consumerTag,
		false,       // autoAck
		false,       // exclusive
		false,       // noLocal
		false,       // noWait
		nil,         // arguments
	)
	if err != nil {
		return fmt.Errorf("Queue Consume: %s", err)
	}

	go handle(deliveries, rabbit.done)

	return nil
}

// Define rabbit service host and port
func LoadRabbitFramework() {
	fmt.Println("serviceRabbit.LoadRabbitFramework()")

	rabbit := serviceRabbit{
		Dsn:      configs.RabbitDsn,
		Consumer: configs.RabbitConsumer,
	}

	rabbit.Init()

	// services.MiddlewareserviceRabbit(rabbit.Server)
	// services.RouteserviceRabbit(rabbit.Server)
	var err error
	err = rabbit.Connect()
	if err != nil {
		fmt.Printf("ERROR rabbit.Serve(): %s\n", err.Error())
		return
	}

	err = rabbit.Register()
	if err != nil {
		fmt.Printf("ERROR rabbit.Register(): %s\n", err.Error())
		return
	}
}
func handle(deliveries <-chan amqp091.Delivery, done chan error) {
	cleanup := func() {
		log.Printf("handle: deliveries channel closed")
		done <- nil
	}

	defer cleanup()

	var deliveryCount int = 0
	var verbose = flag.Bool("verbose", true, "enable verbose output of message data")
	var autoAck = flag.Bool("auto_ack", false, "enable message auto-ack")

	for d := range deliveries {
		deliveryCount++
		if *verbose == true {
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
		if *autoAck == false {
			d.Ack(false)
		}
	}
}
