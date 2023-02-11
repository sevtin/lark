package service

import (
	"github.com/Shopify/sarama"
	"lark/apps/msg_hot/internal/config"
	"lark/domain/mrepo"
	"lark/pkg/common/xkafka"
	"lark/pkg/obj"
)

type MessageHotService interface {
}

type messageHotService struct {
	conf           *config.Config
	messageHotRepo mrepo.MessageHotRepository
	consumerGroup  *xkafka.MConsumerGroup
	msgHandle      map[string]obj.KafkaMessageHandler
}

func NewMessageHotService(conf *config.Config, messageHotRepo mrepo.MessageHotRepository) MessageHotService {
	svc := &messageHotService{conf: conf, messageHotRepo: messageHotRepo}
	svc.msgHandle = make(map[string]obj.KafkaMessageHandler)
	svc.msgHandle[conf.MsgConsumer.Topic[0]] = svc.MessageHandler

	svc.consumerGroup = xkafka.NewMConsumerGroup(&xkafka.MConsumerGroupConfig{KafkaVersion: sarama.V3_2_1_0, OffsetsInitial: sarama.OffsetNewest, IsReturnErr: false},
		conf.MsgConsumer.Topic,
		conf.MsgConsumer.Address,
		conf.MsgConsumer.GroupID)
	svc.consumerGroup.RegisterHandler(svc)

	return svc
}

func (s *messageHotService) Setup(_ sarama.ConsumerGroupSession) error {
	close(s.consumerGroup.Ready)
	return nil
}
func (s *messageHotService) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (s *messageHotService) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
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
