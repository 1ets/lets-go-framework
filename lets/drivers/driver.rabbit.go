package drivers

import (
	"encoding/json"
	"flag"
	"fmt"
	"lets-go-framework/lets"
	"lets-go-framework/lets/types"
	"log"

	"github.com/rabbitmq/amqp091-go"
)

type Body struct {
	Event string      `json:"event"`
	Data  interface{} `json:"data"`
}

type Message struct {
	EventName     string
	Data          []byte
	ReplyTo       string
	CorrelationId string
	Exchange      string
	RoutingKey    string
}

func (m *Message) GetEventName() string {
	return m.EventName
}

func (m *Message) GetData() []byte {
	return m.Data
}

func (m *Message) GetReplyTo() string {
	return m.ReplyTo
}

func (m *Message) GetCorrelationId() string {
	return m.CorrelationId
}

func (m *Message) GetExchange() string {
	return m.Exchange
}

func (m *Message) GetRoutingKey() string {
	return m.RoutingKey
}

type ServiceRabbit struct {
	Dsn        types.IRabbitDsn
	Consumer   types.IRabbitConsumer
	Engine     lets.MessageEngine
	name       string
	dsn        string
	conn       *amqp091.Connection
	channel    *amqp091.Channel
	deliveries <-chan amqp091.Delivery
	done       chan error
}

func (rabbit *ServiceRabbit) Init() {
	fmt.Println("ServiceRabbit.Init()")

	server := rabbit.Dsn
	rabbit.dsn = fmt.Sprintf("amqp://%s:%s@%s:%s/%s", server.GetUsername(), server.GetPassword(), server.GetHost(), server.GetPort(), server.GetVirtualHost())
}

func (rabbit *ServiceRabbit) Connect() error {
	fmt.Println("ServiceRabbit.Connect()")

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

func (rabbit *ServiceRabbit) Register() error {
	fmt.Println("ServiceRabbit.Register()")

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
		return fmt.Errorf("exchange Declare: %s", err)
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
		return fmt.Errorf("queue declare: %s", err)
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
		return fmt.Errorf("queue Bind: %s", err)
	}

	log.Printf("Queue bound to Exchange, starting Consume (consumer tag %q)", c.GetName())
	rabbit.deliveries, err = rabbit.channel.Consume(
		queue.Name,  // name
		c.GetName(), // consumerTag,
		false,       // autoAck
		false,       // exclusive
		false,       // noLocal
		false,       // noWait
		nil,         // arguments
	)
	if err != nil {
		return fmt.Errorf("queue Consume: %s", err)
	}

	go rabbit.listen()

	return nil
}

func (rabbit *ServiceRabbit) listen() {
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

		rabbit.onMessage(&d)

		if !*autoAck {
			d.Ack(false)
		}
	}
}

func (rabbit *ServiceRabbit) onMessage(d *amqp091.Delivery) {
	var body Body
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

	m := Message{
		EventName:     body.Event,
		Data:          data,
		ReplyTo:       d.ReplyTo,
		CorrelationId: d.CorrelationId,
		Exchange:      d.Exchange,
		RoutingKey:    d.RoutingKey,
	}

	rabbit.Engine.Call(m.EventName, &m)
}
