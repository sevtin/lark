package xkafka

import (
	"github.com/IBM/sarama"
	"lark/pkg/common/xlog"
	"sync"
)

type Consumer struct {
	addr          []string
	WG            sync.WaitGroup
	Topic         string
	PartitionList []int32
	Consumer      sarama.Consumer
}

func NewKafkaConsumer(addr []string, topic string) *Consumer {
	c := Consumer{}
	c.Topic = topic
	c.addr = addr

	consumer, err := sarama.NewConsumer(c.addr, nil)
	if err != nil {
		xlog.Error(err.Error())
		return nil
	}
	c.Consumer = consumer

	partitionList, err := consumer.Partitions(c.Topic)
	if err != nil {
		xlog.Error(err.Error())
		return nil
	}
	c.PartitionList = partitionList

	return &c
}
