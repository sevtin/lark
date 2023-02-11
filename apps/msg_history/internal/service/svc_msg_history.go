package service

import (
	"github.com/Shopify/sarama"
	"lark/apps/msg_history/internal/config"
	"lark/domain/cache"
	"lark/domain/repo"
	"lark/pkg/common/xkafka"
	"lark/pkg/obj"
)

type MessageHistoryService interface {
}

type messageHistoryService struct {
	cfg              *config.Config
	chatMessageRepo  repo.ChatMessageRepository
	consumerGroup    *xkafka.MConsumerGroup
	msgHandle        map[string]obj.KafkaMessageHandler
	chatMessageCache cache.ChatMessageCache
}

func NewMessageHistoryService(
	cfg *config.Config,
	chatMessageRepo repo.ChatMessageRepository,
	chatMessageCache cache.ChatMessageCache) MessageHistoryService {
	svc := &messageHistoryService{
		cfg:              cfg,
		chatMessageRepo:  chatMessageRepo,
		chatMessageCache: chatMessageCache}

	svc.msgHandle = make(map[string]obj.KafkaMessageHandler)
	svc.msgHandle[cfg.MsgConsumer.Topic[0]] = svc.MessageHandler

	svc.consumerGroup = xkafka.NewMConsumerGroup(&xkafka.MConsumerGroupConfig{KafkaVersion: sarama.V3_2_1_0, OffsetsInitial: sarama.OffsetNewest, IsReturnErr: false},
		cfg.MsgConsumer.Topic,
		cfg.MsgConsumer.Address,
		cfg.MsgConsumer.GroupID)
	svc.consumerGroup.RegisterHandler(svc)

	return svc
}

// Setup is run at the beginning of a new session, before ConsumeClaim
func (s *messageHistoryService) Setup(_ sarama.ConsumerGroupSession) error {
	// Mark the consumer as ready
	close(s.consumerGroup.Ready)
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (s *messageHistoryService) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
func (s *messageHistoryService) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
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
