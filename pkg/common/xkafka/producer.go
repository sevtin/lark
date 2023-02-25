package xkafka

import (
	"github.com/Shopify/sarama"
	"google.golang.org/protobuf/proto"
	"lark/pkg/common/xlog"
	"lark/pkg/utils"
	"runtime/debug"
)

type Producer struct {
	topic    string
	addr     []string
	config   *sarama.Config
	producer sarama.SyncProducer
}

func NewKafkaProducer(addr []string, topic string) *Producer {
	p := Producer{}
	p.config = sarama.NewConfig()                                   //Instantiate a sarama Config
	p.config.Producer.Return.Successes = true                       //Whether to enable the successes channel to be notified after the message is sent successfully
	p.config.Producer.RequiredAcks = sarama.WaitForAll              //Set producer Message Reply level 0 1 all
	p.config.Producer.Partitioner = sarama.NewRoundRobinPartitioner //Set the hash-key automatic hash partition. When sending a message, you must specify the key value of the message. If there is no key, the partition will be selected randomly

	p.addr = addr
	p.topic = topic

	producer, err := sarama.NewSyncProducer(p.addr, p.config) //Initialize the client
	if err != nil {
		xlog.Error(err.Error())
		return nil
	}
	p.producer = producer
	return &p
}

func (p *Producer) EnQueue(m proto.Message, key ...string) (int32, int64, error) {
	defer func() {
		if r := recover(); r != nil {
			xlog.Warn(r, string(debug.Stack()))
		}
	}()
	msg := &sarama.ProducerMessage{}
	msg.Topic = p.topic
	if len(key) == 1 {
		msg.Key = sarama.StringEncoder(key[0])
	}
	buf, err := proto.Marshal(m)
	if err != nil {
		return -1, -1, err
	}
	msg.Value = sarama.ByteEncoder(buf)
	return p.producer.SendMessage(msg)
}

func (p *Producer) Push(m interface{}, key ...string) (int32, int64, error) {
	defer func() {
		if r := recover(); r != nil {
			xlog.Warn(r, string(debug.Stack()))
		}
	}()
	msg := &sarama.ProducerMessage{}
	msg.Topic = p.topic
	if len(key) == 1 {
		msg.Key = sarama.StringEncoder(key[0])
	}
	buf, err := utils.ObjToByte(m)
	if err != nil {
		return -1, -1, err
	}
	msg.Value = sarama.ByteEncoder(buf)
	return p.producer.SendMessage(msg)
}

func (p *Producer) Close() {
	err := p.producer.Close()
	if err != nil {
		xlog.Warn(err.Error())
	}
}
