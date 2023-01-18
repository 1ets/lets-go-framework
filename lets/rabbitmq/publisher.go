package rabbitmq

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"lets-go-framework/lets/types"
	"log"
	"os"
	"time"

	"github.com/rabbitmq/amqp091-go"
)

var RabbitMQDsn types.IRabbitMQDsn
var Publisher types.IRabbitMQPublisher

// HTTP service struct
type RabbitClient struct {
	Dsn        string
	Publisher  types.IRabbitMQPublisher
	Engine     Engine
	Config     amqp091.Config
	Connection *amqp091.Connection
	Channel    *amqp091.Channel
	Client     func(*RabbitClient)
}

func (r *RabbitClient) Init() {
	fmt.Println("ServiceRabbit.Init()")

	r.Publisher = Publisher
	r.Dsn = fmt.Sprintf("amqp://%s:%s@%s:%s/%s", RabbitMQDsn.GetUsername(), RabbitMQDsn.GetPassword(), RabbitMQDsn.GetHost(), RabbitMQDsn.GetPort(), RabbitMQDsn.GetVirtualHost())
	r.Config = amqp091.Config{Properties: amqp091.NewConnectionProperties()}
	r.Config.Properties.SetClientConnectionName(Publisher.GetName())
}

func (rabbit *RabbitClient) Connect() {
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

}

// Publish event
func (rabbit *RabbitClient) Publish(event Event) error {
	var (
		Log = log.New(os.Stdout, "[INFO] ", log.LstdFlags|log.Lmsgprefix)
	)

	flag.Parse()

	// Log.Println("declared Exchange, publishing messages")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Encode object to json string
	body, _ := json.Marshal(types.RabbitBody{Event: event.Name, Data: event.Data})

	// for {
	// seqNo := rabbit.channel.GetNextPublishSeqNo()
	Log.Printf("publishing %dB body (%q)", len(body), body)

	if err := rabbit.Channel.PublishWithContext(ctx,
		event.Exchange,   // Exchange
		event.RoutingKey, // RoutingKey or queues
		false,            // Mandatory
		false,            // Immediate
		amqp091.Publishing{
			Headers:         amqp091.Table{},
			ContentType:     "application/json",
			ContentEncoding: "",
			Body:            []byte(body),
			DeliveryMode:    amqp091.Transient, // 1=non-persistent, 2=persistent
			Priority:        0,                 // 0-9
			// a bunch of application/implementation-specific fields
			ReplyTo:       event.ReplyTo.GetJson(),
			CorrelationId: event.CorrelationId,
		},
	); err != nil {
		return fmt.Errorf("exchange Publish: %s", err)
	}

	Log.Printf("published %dB OK", len(body))
	// if reliable {
	// 	publishes <- seqNo
	// }
	// }
	return nil
}

// Publish event
func (r *RabbitClient) GetDst() types.IRabbitMQPublisher {

	return r.Publisher
}

// Define rabbit service host and port
func SetupPublisher() {
	fmt.Println("SetupPublisher()")

	var rabbitClient RabbitClient

	rabbitClient.Init()
	rabbitClient.Connect()
	rabbitClient.Client(&rabbitClient)
}
