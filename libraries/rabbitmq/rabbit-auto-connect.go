package rabbitmq

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/kataras/golog"
	"github.com/streadway/amqp"
)

// IRqAutoConnect is interface defining method of rabbit mq auto connect
type IRqAutoConnect interface {
	StartConnection(username, password, host, port, vhost string) (c *amqp.Connection, err error)
	DeclareQueues(queues ...string) (err error)
	GetRqChannel() *amqp.Channel
	Stop()
	beforeReconnect() // implement template pattern
	afterReconnect()  // implement template pattern
}

type rMqAutoConnect struct {
	conn           *amqp.Connection
	ch             *amqp.Channel
	uriConnection  string
	notifCloseCh   chan *amqp.Error
	ctxReconnect   context.Context
	stopReconnect  context.CancelFunc
	rq             IRqAutoConnect // implement template pattern
	declaredQueues []string       // queues
}

func (r *rMqAutoConnect) reset() {
	r.ch.Close()
	r.conn.Close()
}

func (r *rMqAutoConnect) connect(uri string) (c *amqp.Connection, err error) {
	const (
		maxTrialSecond = 3 // 60 second
		maxTrialMinute = 7 // 10 minute
	)

	// connect to rabbit mq
	golog.Info("RabbitMQ: ", "connecting to "+uri+" ...")
	trial := 0
	for {
		trial++
		r.conn, err = amqp.Dial(uri)
		if err != nil {

			golog.Error("RabbitMQ: ", err)
			switch {
			case trial <= maxTrialSecond:
				golog.Info("RabbitMQ: ", "try to reconnect in 5 seconds ...")
				<-time.After(time.Duration(5) * time.Second)
			case trial <= maxTrialMinute:
				golog.Info("RabbitMQ: ", "try to reconnect in 30 seconds ...")
				<-time.After(time.Duration(10) * time.Second)
			default:
				golog.Info("RabbitMQ: ", "try to reconnect in 1 minute ...")
				<-time.After(time.Duration(1) * time.Minute)
			}
			continue
		}
		break
	}

	// keep a live
	r.conn.Config.Heartbeat = time.Duration(5) * time.Second

	//declare channel
	golog.Debug("RabbitMQ: ", "opening channel ...")

	r.ch, err = r.conn.Channel()
	if err != nil {
		r.conn.Close()
		golog.Fatal("RabbitMQ: ", err)
	}

	golog.Debug("RabbitMQ: ", "channel opened sucessfully")
	return r.conn, nil
}

func (r *rMqAutoConnect) DeclareQueues(queues ...string) (err error) {
	r.declaredQueues = queues
	//declare queues
	for _, queue := range queues {

		golog.Info("RabbitMQ: ", fmt.Sprintf("declare %s queue ...\n", queue))
		_, err = r.ch.QueueDeclare(
			queue, //name
			//true,  //durable
			false, //durable
			false, //auto delte
			false, //exclusive
			false, //no wait
			func() (out amqp.Table) {
				return
			}(), //args
		)
		if err != nil {
			log.Println(err.Error())
			return
		}

		golog.Info("RabbitMQ: ", fmt.Sprintf("queue %s is successfully declared\n", queue))
	}
	return
}

func (r *rMqAutoConnect) stop() {
	defer func() {
		if it := recover(); it != nil {
			golog.Debug("RabbitMQ: ", fmt.Sprintf("panic : %v\n", it))
		}
	}()
	if r.stopReconnect != nil {
		r.stopReconnect()
	}
	r.reset()
}

func (r *rMqAutoConnect) GetRqChannel() *amqp.Channel {
	return r.ch
}

func (r *rMqAutoConnect) beforeReconnect() {
	r.rq.beforeReconnect()
}

func (r *rMqAutoConnect) afterReconnect() {
	r.rq.afterReconnect()
}

func (r *rMqAutoConnect) startConnection(username, password, host, port, vhost string) (err error) {
	// set uri parameter to connect to rabbit mq
	r.uriConnection = fmt.Sprintf("amqp://%s:%s@%s:%s/%s", username, password, host, port, vhost)
	r.conn, err = r.connect(r.uriConnection)
	if err != nil {
		golog.Fatal(err)
	}

	// try to reconnect
	r.reconnect()

	return
}

func (r *rMqAutoConnect) getConnection() *amqp.Connection {
	return r.conn
}

func (r *rMqAutoConnect) reconnect() {
	r.ctxReconnect, r.stopReconnect = context.WithCancel(context.Background()) // prepare context
	r.notifCloseCh = make(chan *amqp.Error)
	go func() {
		for {
			select {
			case <-r.ctxReconnect.Done():
				return
			case <-r.getConnection().NotifyClose(r.notifCloseCh):
				r.beforeReconnect()

				golog.Debug("RabbitMQ: ", "connection is closed, reconnecting ...")

				r.reset()
				r.connect(r.uriConnection)
				r.DeclareQueues(r.declaredQueues...)
				r.afterReconnect()
				r.notifCloseCh = make(chan *amqp.Error)
			}
		}
	}()
}
