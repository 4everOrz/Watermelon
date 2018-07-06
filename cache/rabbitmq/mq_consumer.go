package cache

import (
	"fmt"
	"log"
	"time"

	"github.com/astaxie/beego"
	"github.com/streadway/amqp"
)

var Timercount = 0

type MqConsumer struct {
	channel     *amqp.Channel
	conn        *amqp.Connection
	connflag    chan int
	RabbitError chan *amqp.Error
}

var receiver = make(chan MqConsumer, 5)
var consumer *MqConsumer

func (mq *MqConsumer) mqConnect() error {
	var err error

	mq.conn, err = amqp.Dial(MMqurl)
	if err == nil {
		mq.channel, err = mq.conn.Channel()
	}
	return err
}
func (mq *MqConsumer) mqClose() {
	mq.channel.Close()
	mq.conn.Close()
}

func (c *MqConsumer) connectToRabbitMQ() {
	var err error
	for {
		c.conn, err = amqp.Dial(MMqurl)

		if err == nil {
			c.connflag <- 1
			break
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func (c *MqConsumer) rabbitConnector() {
	var rabbitErr *amqp.Error
	for {
		rabbitErr = <-c.RabbitError
		if rabbitErr != nil {
			c.connectToRabbitMQ()
			c.RabbitError = make(chan *amqp.Error)
			c.conn.NotifyClose(c.RabbitError)
		}
	}
}

//获取一个消费者对象，并建立连接
func GetConsumerConn(connFlag chan int) *MqConsumer {

	if consumer == nil {
		consumer = &MqConsumer{
			RabbitError: make(chan *amqp.Error),
			connflag:    connFlag,
		}
		consumer.connectToRabbitMQ()
		consumer.conn.NotifyClose(consumer.RabbitError)
		go consumer.rabbitConnector()
	}
	return consumer
}

//监听连接，中断重连
func consumerConnListener(consumer *MqConsumer) {
	cc := make(chan *amqp.Error)
	e := <-consumer.conn.NotifyClose(cc)
	log.Println("mqconsumer  connect error,try reconnect", e)
	consumer.mqClose()
	consumer.conn = nil
	consumer.channel = nil
	consumer.rabbitConnector()
	consumer.channel, _ = consumer.conn.Channel()
}

//接收数据
func Catchdata(*MqConsumer) {
	defer consumer.mqClose()
	var err error
	if consumer.channel == nil {
		consumer.channel, _ = consumer.conn.Channel()
	}
	msgs, err := consumer.channel.Consume(MQueueName, "", true, false, false, false, nil)
	if err != nil {
		beego.Error(err, "mqconsumer  connect error,try reconnect!")
		return
	}
	for d := range msgs {
		fmt.Println("msg", d.Body)
	}
}

//mq_main()
func MqConsumer_Run() {
	connSucessFlag := make(chan int, 1)
	mqconsumer := GetConsumerConn(connSucessFlag)
	go func() {
		for {
			<-mqconsumer.connflag
			go Catchdata(mqconsumer)
			go consumerConnListener(mqconsumer)
		}
	}()
}
