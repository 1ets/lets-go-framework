package rabbitmq

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/kataras/golog"
	"github.com/streadway/amqp"
)

// IRqAutoConsumer is interface defining method of rabbit mq auto connect for consumer
type IRqAutoConsumer interface {
	IRqAutoConnect
	GetMessageChanel(queue string) <-chan amqp.Delivery
	SetReadQueue(queue ...string)
	ConsumeMessage() (err error)
}

// Consumer information.
type rMqAutoConsumer struct {
	rMqAutoConnect
	deliveryCh     map[string]<-chan amqp.Delivery
	msgCh          map[string]chan amqp.Delivery
	ctxConsumeMsg  map[string]context.Context
	stopConsumeMsg map[string]context.CancelFunc
	isBroken       bool
	readQueue      []string
}

// Stop consuming message.
func (r *rMqAutoConsumer) Stop() {
	for _, stopConsumeMsg := range r.stopConsumeMsg {
		stopConsumeMsg()
	}
	r.stop()
}

// Setup read queue.
func (r *rMqAutoConsumer) SetReadQueue(queue ...string) {
	r.readQueue = queue
}

// Make a connection into RabbitMQ server.
func (r *rMqAutoConsumer) StartConnection(username, password, host, port, vhost string) (c *amqp.Connection, err error) {
	err = r.startConnection(username, password, host, port, vhost)
	if err != nil {
		log.Panicln(err.Error())
	}
	c = r.conn
	return
}

// Listening queue on selected channel.
func (r *rMqAutoConsumer) listenQueueOnChannel() (err error) {
	// make sure that only one message at one time
	err = r.GetRqChannel().Qos(
		1,     //prefetch count
		0,     //prefetch size
		false, //global
	)
	if err != nil {
		r.ch.Close()
		r.conn.Close()
		golog.Fatal(err.Error())
	}

	for _, readQueue := range r.readQueue {
		golog.Info("RabbitMQ: ", fmt.Sprintf("listening queue: '%s'", readQueue))

		r.deliveryCh[readQueue], err = r.ch.Consume(
			readQueue, //queue
			//config.RqNotifQueue(), //name
			"",    //consumer
			false, //auto ack
			false, //exclusive
			false, //no local
			false, //no wait
			nil,   //args
		)
		if err != nil {
			golog.Fatal("RabbitMQ: ", err)
		}
	}
	r.isBroken = false

	return
}

// Consuming message / retrieve mesage and process stream reading.
func (r *rMqAutoConsumer) ConsumeMessage() (err error) {
	r.deliveryCh = map[string]<-chan amqp.Delivery{}
	r.msgCh = map[string]chan amqp.Delivery{}
	r.ctxConsumeMsg = map[string]context.Context{}
	r.stopConsumeMsg = map[string]context.CancelFunc{}
	r.listenQueueOnChannel()
	for _, readQueue := range r.readQueue {
		golog.Debug("RabbitMQ: ", fmt.Sprintf("create context with cancel, delivery chanel, message chanel on queue : %s", readQueue))

		// prepare context
		r.ctxConsumeMsg[readQueue], r.stopConsumeMsg[readQueue] = context.WithCancel(context.Background())
		// prepare chanel
		r.msgCh[readQueue] = make(chan amqp.Delivery)
		go func(readQueue string) {
			//defer close(r.msgCh[readQueue])
			for {
				if r.isBroken {
					<-time.After(time.Duration(1) * time.Second)
					continue
				}
				select {
				case <-r.ctxConsumeMsg[readQueue].Done():
					close(r.msgCh[readQueue])
					golog.Info("RabbitMQ: ", fmt.Sprintf("stop consuming message on %s\n", readQueue))
					return
				case <-time.After(time.Duration(1) * time.Second):
				case delivery := <-r.deliveryCh[readQueue]:
					if delivery.Body == nil {
						golog.Error("RabbitMQ: ", "queue may be not exist or server down")
						// //r.reconnect()

						<-time.After(time.Duration(1) * time.Second)
					} else {
						golog.Info("RabbitMQ: ", fmt.Sprintf("receiving %s", readQueue))
						golog.Debug("RabbitMQ: ", "received body \n", string(delivery.Body))
					}
					r.msgCh[readQueue] <- delivery
				}
			}
		}(readQueue)
	}
	return
}

// Get message by channel.
func (r *rMqAutoConsumer) GetMessageChanel(queue string) <-chan amqp.Delivery {
	return r.msgCh[queue]
}

func (r *rMqAutoConsumer) beforeReconnect() { // implement template pattern
	r.isBroken = true
}

func (r *rMqAutoConsumer) afterReconnect() { // implement template pattern
	r.listenQueueOnChannel()
}

// CreateRqConsumer is function to create rabbit mq auto connect for consumer
func CreateRqConsumer() (r IRqAutoConsumer) {
	rmq := new(rMqAutoConsumer)
	rmq.rq = rmq
	return rmq
}
