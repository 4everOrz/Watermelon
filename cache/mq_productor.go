package cache

import (
	"Watermelon/config"
	"log"
	"time"

	"github.com/streadway/amqp"
)

type MqProducer struct {
	text     chan []byte
	channel  *amqp.Channel
	conn     *amqp.Connection
	connflag chan int
	mqerror  chan *amqp.Error
}

var (
	ProducerConnTag = true //生产者连接状态标志，使用过程中通过此标志判断是否发送数据
	mqproducer      *MqProducer
	MQueueName      string
	MExchange       string
	MExchangeType   string
	MMqurl          string
	MKey            string
)

func init() {
	conf := config.GetConf()
	MQueueName = conf.MQueueName
	MExchange = conf.MExchange
	MExchangeType = conf.MExchangeType
	MMqurl = conf.MMqurl
	MKey = conf.MKey
}

func (mq *MqProducer) pingmq() {
	var err error

	for {
		mq.conn, err = amqp.Dial(MMqurl)

		if err == nil {
			mq.connflag <- 1
			break
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func (mq *MqProducer) mqConnect() error {
	var err error
	mq.conn, err = amqp.Dial(MMqurl)
	if err == nil {
		mq.channel, err = mq.conn.Channel()
	}
	return err
}
func (mq *MqProducer) mqClose() {
	mq.channel.Close()
	mq.conn.Close()
}

func (mq *MqProducer) push(msgContent []byte) {
	if mq.channel == nil {
		mq.channel, _ = mq.conn.Channel()
	}

	err := mq.channel.Publish(MExchange, MKey, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        msgContent,
	})
	if err != nil {
		log.Println("发送失败")
	}
}

/*****************************************/

func (c *MqProducer) connectToRabbitMQ() {
	var err error
	for {
		c.conn, err = amqp.Dial(MMqurl)

		if err == nil {
			c.connflag <- 1
			ProducerConnTag = true
			break
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func getProducerConn(connFlag chan int) *MqProducer {
	if mqproducer == nil {
		mqproducer = &MqProducer{
			mqerror:  make(chan *amqp.Error),
			connflag: connFlag,
			text:     make(chan []byte),
		}
		mqproducer.connectToRabbitMQ()
		mqproducer.conn.NotifyClose(mqproducer.mqerror)
		go mqproducer.rabbitConnector()
	}
	return mqproducer
}

func (c *MqProducer) rabbitConnector() {
	var rabbitErr *amqp.Error
	for {
		rabbitErr = <-c.mqerror
		if rabbitErr != nil {
			c.connectToRabbitMQ()
			c.mqerror = make(chan *amqp.Error)
			c.conn.NotifyClose(c.mqerror)
		}
	}
}

func producerConnListener() {
	cc := make(chan *amqp.Error)
	e := <-mqproducer.conn.NotifyClose(cc)
	ProducerConnTag = false
	log.Println("mqproducer connect error,try reconnect", e)
	mqproducer.conn = nil
	mqproducer.channel = nil
	mqproducer.rabbitConnector()
	mqproducer.channel, _ = mqproducer.conn.Channel()
}

func (mq *MqProducer) translate() {
	var saveData []byte
	for {
		select {
		case saveData = <-mq.text:
			mq.push(saveData)
		}
	}
	defer mq.mqClose()
}

//mq_main()
func GetMqProducter() *MqProducer {
	connSucessFlag := make(chan int, 1)
	mqproducer := getProducerConn(connSucessFlag)
	go func() {
		for {
			<-mqproducer.connflag
			go mqproducer.translate() //数据发送
			go producerConnListener() //开启异常监听线程
		}
	}()
	return mqproducer
}

//生产者使用举例
/*func example() {
	var msg []byte
	msg=?                                                          //先定义一个消息字段,然后赋值，此处略去30字
	if ProducerConnTag == true {                                  //判断生者连接标志是否正常
		GetMqProducter().text <- msg                              //将消息放入消息通道
	} else {
		fmt.Println("The connection of rabbitmq has been cut down. Wait a moment until the connection is successful") //我都开始嫉妒自己的英语水平了
	}
}*/
