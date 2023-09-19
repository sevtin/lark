package service

import (
	"github.com/IBM/sarama"
	"lark/apps/cache/internal/config"
	"lark/domain/cache"
	"lark/pkg/common/xkafka"
	"lark/pkg/obj"
)

type CacheService interface {
}

type cacheService struct {
	cfg             *config.Config
	consumerGroup   *xkafka.MConsumerGroup
	msgHandle       map[string]obj.KafkaMessageHandler
	chatMemberCache cache.ChatMemberCache
}

func NewCacheService(cfg *config.Config, chatMemberCache cache.ChatMemberCache) CacheService {
	svc := &cacheService{
		cfg:             cfg,
		chatMemberCache: chatMemberCache}
	svc.msgHandle = make(map[string]obj.KafkaMessageHandler)
	svc.msgHandle[cfg.MsgConsumer.Topic[0]] = svc.MessageHandler
	svc.consumerGroup = xkafka.NewMConsumerGroup(&xkafka.MConsumerGroupConfig{KafkaVersion: sarama.V3_2_1_0, OffsetsInitial: sarama.OffsetNewest, IsReturnErr: false},
		cfg.MsgConsumer.Topic,
		cfg.MsgConsumer.Address,
		cfg.MsgConsumer.GroupID)
	svc.consumerGroup.RegisterHandler(svc)
	return svc
}
