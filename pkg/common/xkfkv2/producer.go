package xkafka

import (
	"errors"
	"github.com/IBM/sarama"
	"google.golang.org/protobuf/proto"
	"lark/pkg/common/xlog"
	"lark/pkg/conf"
	"lark/pkg/utils"
	"runtime/debug"
)

type Producer struct {
	cfg      *conf.KafkaProducer
	config   *sarama.Config
	producer sarama.SyncProducer
}

type PartitionerType int

const (
	NewHashPartitioner    PartitionerType = 0 // 哈希分区
	ManualPartitioner     PartitionerType = 1 // 手动分区
	RoundRobinPartitioner PartitionerType = 2 // 随机分区
)

func NewKafkaProducer(cfg *conf.KafkaProducer) *Producer {
	p := Producer{cfg: cfg}
	p.config = sarama.NewConfig()                      // Instantiate a sarama Config
	p.config.Producer.Return.Successes = true          // Whether to enable the successes channel to be notified after the message is sent successfully
	p.config.Producer.RequiredAcks = sarama.WaitForAll // Set producer Message Reply level 0 1 all
	// p.config.Producer.Retry.Max = 5                    // The total number of times to retry sending a message (default 3).
	// Set the hash-key automatic hash partition. When sending a message, you must specify the key value of the message. If there is no key, the partition will be selected randomly
	if cfg.Sasl != nil && cfg.Sasl.Enable == true {
		p.config.Net.SASL.Enable = true
		p.config.Net.SASL.User = cfg.Sasl.User
		p.config.Net.SASL.Password = cfg.Sasl.Password
	}

	switch PartitionerType(cfg.Partitioner) {
	case NewHashPartitioner:
		p.config.Producer.Partitioner = sarama.NewHashPartitioner
	case ManualPartitioner:
		p.config.Producer.Partitioner = sarama.NewManualPartitioner
	case RoundRobinPartitioner:
		p.config.Producer.Partitioner = sarama.NewRoundRobinPartitioner
	default:
		p.config.Producer.Partitioner = sarama.NewHashPartitioner
	}

	producer, err := sarama.NewSyncProducer(p.cfg.Address, p.config) // Initialize the client
	if err != nil {
		xlog.Errorf("create kafka producer error: %s", err.Error())
	}
	p.producer = producer
	return &p
}

func (p *Producer) EnQueue(m proto.Message, key string) (partition int32, offset int64, err error) {
	defer func() {
		if r := recover(); r != nil {
			xlog.Warn(r, string(debug.Stack()))
		}
	}()
	msg := &sarama.ProducerMessage{}
	msg.Topic = p.cfg.Topic
	msg.Key = sarama.StringEncoder(key)
	buf, err := proto.Marshal(m)
	if err != nil {
		return -1, -1, err
	}
	msg.Value = sarama.ByteEncoder(buf)
	partition, offset, err = p.producer.SendMessage(msg)
	if err != nil {
		xlog.Errorf("failed to send message. error: %s, partition: %d, offset: %d, message: %s", err.Error(), partition, offset, string(buf))
	}
	return
}

func (p *Producer) Push(m interface{}, key string) (partition int32, offset int64, err error) {
	defer func() {
		if r := recover(); r != nil {
			xlog.Warn(r, string(debug.Stack()))
		}
	}()
	if p.producer == nil {
		return -1, -1, errors.New("kafka producer is nil")
	}
	msg := &sarama.ProducerMessage{}
	msg.Topic = p.cfg.Topic
	msg.Key = sarama.StringEncoder(key)
	buf, err := utils.EncodeToBytes(m)
	if err != nil {
		return -1, -1, err
	}
	msg.Value = sarama.ByteEncoder(buf)
	partition, offset, err = p.producer.SendMessage(msg)
	if err != nil {
		xlog.Errorf("failed to send message. error: %s, partition: %d, offset: %d, message: %s", err.Error(), partition, offset, string(buf))
	}
	return
}

func (p *Producer) PushBuffer(buf []byte, key string) (partition int32, offset int64, err error) {
	defer func() {
		if r := recover(); r != nil {
			xlog.Warn(r, string(debug.Stack()))
		}
	}()
	msg := &sarama.ProducerMessage{}
	msg.Key = sarama.StringEncoder(key)
	msg.Topic = p.cfg.Topic
	msg.Value = sarama.ByteEncoder(buf)
	partition, offset, err = p.producer.SendMessage(msg)
	if err != nil {
		xlog.Errorf("failed to send message. error: %s, partition: %d, offset: %d, message: %s", err.Error(), partition, offset, string(buf))
	}
	return
}

func (p *Producer) Close() {
	err := p.producer.Close()
	if err != nil {
		xlog.Warnf("close kafka producer error: %s", err.Error())
	}
}

func (p *Producer) GetProducer() sarama.SyncProducer {
	return p.producer
}
