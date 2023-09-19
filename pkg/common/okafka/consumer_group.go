package okafka

import (
	"context"
	"github.com/IBM/sarama"
	"lark/pkg/common/xlog"
)

type MConsumerGroup struct {
	sarama.ConsumerGroup
	groupID string
	topics  []string
}

type MConsumerGroupConfig struct {
	KafkaVersion   sarama.KafkaVersion
	OffsetsInitial int64
	IsReturnErr    bool
}

func NewMConsumerGroup(consumerConfig *MConsumerGroupConfig, topics, addr []string, groupID string) (group *MConsumerGroup) {
	group = &MConsumerGroup{
		ConsumerGroup: nil,
		groupID:       groupID,
		topics:        topics,
	}
	var (
		config        = sarama.NewConfig()
		client        sarama.Client
		consumerGroup sarama.ConsumerGroup
		err           error
	)
	config.Version = consumerConfig.KafkaVersion
	config.Consumer.Offsets.Initial = consumerConfig.OffsetsInitial
	config.Consumer.Return.Errors = consumerConfig.IsReturnErr

	client, err = sarama.NewClient(addr, config)
	if err != nil {
		xlog.Error(err.Error())
		return
	}
	consumerGroup, err = sarama.NewConsumerGroupFromClient(groupID, client)
	if err != nil {
		xlog.Error(err.Error())
		return
	}
	group.ConsumerGroup = consumerGroup
	return
}

func (mc *MConsumerGroup) RegisterHandleAndConsumer(handler sarama.ConsumerGroupHandler) {
	var (
		ctx = context.Background()
		err error
	)
	for {
		if mc.ConsumerGroup == nil {
			xlog.Error("consumer group is null")
			break
		}
		err = mc.ConsumerGroup.Consume(ctx, mc.topics, handler)
		if err != nil {
			xlog.Error(err.Error())
			break
		}
	}
}
