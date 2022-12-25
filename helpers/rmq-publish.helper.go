package helpers

import (
	"encoding/base64"
	"encoding/json"
	"lets-go-framework/libraries/rabbitmq"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/kataras/golog"
	"github.com/streadway/amqp"
)

type IRandNumber interface {
	GenerateNumber() int
}

type myRandNumber int

func (myRandNumber) GenerateNumber() int {
	return rand.Int()
}

var RandomNumber IRandNumber = myRandNumber(0)

type PayloadRequest struct {
	Sender string      `json:"sender"`
	Event  string      `json:"event"`
	Time   string      `json:"time"`
	Data   interface{} `json:"data"`
}

type PayloadResponse struct {
	Status  string      `json:"status"`  // SUCCESS, FAILED
	Message string      `json:"message"` // Readable message for human
	Data    interface{} `json:"data"`
}

type RmqPublish struct {
	credentials   rabbitmq.ICredentials
	rqConnection  *amqp.Connection
	rqChannel     *amqp.Channel
	event         string
	replyQueue    string
	replyTimeOut  string
	correlationId string
	waitTimeOut   time.Duration
	payload       PayloadRequest
	rbMq          rabbitmq.IRqAutoPubConsumer
}

func (r *RmqPublish) SetCredential(credentials rabbitmq.ICredentials) *RmqPublish {
	r.credentials = credentials
	return r
}

func (r *RmqPublish) SetEvent(event string) *RmqPublish {
	r.event = event
	return r
}

func (r *RmqPublish) SetData(data interface{}) *RmqPublish {
	r.payload = PayloadRequest{
		Sender: os.Getenv("SERVICE_ID"),
		Event:  r.event,
		Time:   time.Now().String(),
		Data:   data,
	}

	return r
}

func (r *RmqPublish) start() {
	// Aliasing
	r.rbMq = rabbitmq.CreateRqPubConsumer()

	var err error
	r.rqConnection, err = r.rbMq.StartConnection(
		r.credentials.GetUsername(),
		r.credentials.GetPassword(),
		r.credentials.GetHost(),
		r.credentials.GetPort(),
		r.credentials.GetVirtualHost())

	if err != nil {
		golog.Errorf("RabbitMQ: %s. [050000]", err)
	}

	r.rqChannel = r.rbMq.GetRqChannel()
}

func (r *RmqPublish) stop() {
	r.rbMq.Stop()
}

func randomByte(n int) string {
	b := make([]byte, n)
	_, _ = rand.Read(b)

	return base64.URLEncoding.EncodeToString(b)
}

func (r *RmqPublish) createReplyQueue() <-chan amqp.Delivery {
	r.correlationId = randomByte(8)
	r.replyQueue = "plus_wd_sync_reply-" + strconv.Itoa(int(time.Now().Unix())) + r.correlationId

	_, err := r.rqChannel.QueueDeclare(
		r.replyQueue,
		true,
		true,
		false,
		false,
		nil,
	)
	if err != nil {
		golog.Errorf("RabbitMQ: %s", err)
		r.replyQueue = ""
		return nil
	}

	deliver, err := r.rqChannel.Consume(
		r.replyQueue,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		golog.Errorf("RabbitMQ: %s", err)
		r.replyQueue = ""
		return nil
	}
	r.replyTimeOut = "60000"
	r.waitTimeOut = 60 * time.Second

	return deliver
}

func (r *RmqPublish) Publish(fReply bool) (request PayloadRequest, response PayloadResponse) {
	golog.Infof("RabbitMQ: Publish message: to \"%s\" ...", r.credentials.GetExchange())

	request = r.payload
	payload := ToJson(r.payload)
	golog.Debugf("RabbitMQ: Publish payload: \n%s", payload)

	r.start()
	defer r.stop()

	// Create delivery channel
	delivery := make(<-chan amqp.Delivery)
	if fReply {
		delivery = r.createReplyQueue()
		defer r.rqChannel.QueueDelete(r.replyQueue, false, false, true)
	}

	err := r.rqChannel.Publish(
		r.credentials.GetExchange(),
		r.credentials.GetRoutingKey(),
		false,
		false,
		amqp.Publishing{
			ContentType:   "application/json",
			Body:          []byte(payload),
			Expiration:    r.replyTimeOut,
			ReplyTo:       r.replyQueue,
			CorrelationId: r.correlationId,
		},
	)

	if err != nil {
		golog.Errorf("RabbitMQ: %s. [050001]", err)
		return
	}

	golog.Info("RabbitMQ: message published successfully.")

	if fReply {
		golog.Info("RabbitMQ: waiting synchronus reply.")
		waitTimeOut := time.After(r.waitTimeOut)

		for fReply {
			golog.Debug("waiting")
			select {
			case <-waitTimeOut:
				fReply = false
				response = PayloadResponse{
					Status:  "error",
					Message: "Server did not respond.",
				}

				return
			case data := <-delivery:
				golog.Info("RabbitMQ: got reply.")

				if r.correlationId == data.CorrelationId {
					if payload == string(data.Body) {
						response = PayloadResponse{
							Status:  "error",
							Message: "Server did not respond.",
						}
						return
					} else {
						var reply map[string]interface{}
						json.Unmarshal([]byte(string(data.Body)), &reply)

						response = PayloadResponse{
							Status:  reply["StatusMessage"].(string),
							Message: reply["ErrorCode"].(string),
							Data:    reply["Data"],
						}
					}
					fReply = false
				}
			}
		}
	}

	return
}
