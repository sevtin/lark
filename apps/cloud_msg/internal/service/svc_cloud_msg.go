package service

import (
	"github.com/Shopify/sarama"
	"lark/apps/cloud_msg/internal/config"
	"lark/pkg/common/xkafka"
	"lark/pkg/obj"
)

type CloudMessageService interface {
}

type cloudMessageService struct {
	cfg           *config.Config
	consumerGroup *xkafka.MConsumerGroup
	msgHandle     map[string]obj.KafkaMessageHandler
	msgCount      int64
}

func NewCloudMessageService(cfg *config.Config) CloudMessageService {
	svc := &cloudMessageService{cfg: cfg}
	svc.msgHandle = make(map[string]obj.KafkaMessageHandler)
	svc.msgHandle[cfg.MsgConsumer.Topic[0]] = svc.MessageHandler

	svc.consumerGroup = xkafka.NewMConsumerGroup(&xkafka.MConsumerGroupConfig{KafkaVersion: sarama.V3_2_1_0, OffsetsInitial: sarama.OffsetNewest, IsReturnErr: false},
		cfg.MsgConsumer.Topic,
		cfg.MsgConsumer.Address,
		cfg.MsgConsumer.GroupID)
	svc.consumerGroup.RegisterHandler(svc)

	return svc
}

func (s *cloudMessageService) Setup(_ sarama.ConsumerGroupSession) error {
	close(s.consumerGroup.Ready)
	return nil
}
func (s *cloudMessageService) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (s *cloudMessageService) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	var (
		msg *sarama.ConsumerMessage
		err error
	)
	for {
		select {
		case msg = <-claim.Messages():
			if msg == nil {
				continue
			}
			if err = s.msgHandle[msg.Topic](msg.Value, string(msg.Key)); err != nil {
				continue
			}
			session.MarkMessage(msg, "")
		case <-session.Context().Done():
			return nil
		}
	}
	return nil
}
