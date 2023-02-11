package service

import (
	"context"
	"github.com/go-playground/validator/v10"
	chat_member_client "lark/apps/chat_member/client"
	"lark/apps/message/internal/config"
	"lark/domain/cache"
	"lark/pkg/common/xkafka"
	"lark/pkg/proto/pb_msg"
)

type MessageService interface {
	SendChatMessage(ctx context.Context, req *pb_msg.SendChatMessageReq) (resp *pb_msg.SendChatMessageResp, err error)
}

type messageService struct {
	cfg              *config.Config
	validate         *validator.Validate
	producer         *xkafka.Producer
	chatMemberClient chat_member_client.ChatMemberClient
	chatMemberCache  cache.ChatMemberCache
	chatMessageCache cache.ChatMessageCache
}

func NewMessageService(
	cfg *config.Config,
	chatMemberCache cache.ChatMemberCache,
	chatMessageCache cache.ChatMessageCache) MessageService {
	chatMemberClient := chat_member_client.NewChatMemberClient(cfg.Etcd, cfg.ChatMemberServer, cfg.GrpcServer.Jaeger, cfg.Name)
	svc := &messageService{
		cfg:              cfg,
		validate:         validator.New(),
		chatMemberClient: chatMemberClient,
		chatMemberCache:  chatMemberCache,
		chatMessageCache: chatMessageCache}
	svc.producer = xkafka.NewKafkaProducer(cfg.MsgProducer.Address, cfg.MsgProducer.Topic)
	return svc
}
