package service

import (
	"context"
	"github.com/Shopify/sarama"
	"lark/apps/chat_member/client"
	"lark/apps/dist/internal/config"
	gw_client "lark/apps/msg_gateway/client"
	"lark/domain/cache"
	"lark/pkg/common/xetcd"
	"lark/pkg/common/xkafka"
	"lark/pkg/obj"
	"lark/pkg/proto/pb_dist"
	"sync"
)

type DistService interface {
	ChatInviteNotification(ctx context.Context, req *pb_dist.ChatInviteNotificationReq) (resp *pb_dist.ChatInviteNotificationResp, err error)
}

type distService struct {
	cfg              *config.Config
	mutex            sync.RWMutex
	msgHandle        map[string]obj.KafkaMessageHandler
	consumerGroup    *xkafka.MConsumerGroup
	producers        map[int64]*xkafka.Producer
	clients          map[int64]gw_client.MessageGatewayClient
	chatMemberClient chat_member_client.ChatMemberClient
	watcher          *xetcd.Watcher
	serverMgrCache   cache.ServerMgrCache
	chatMemberCache  cache.ChatMemberCache
	queues           chan struct{}
}

func NewDistService(cfg *config.Config, serverMgrCache cache.ServerMgrCache, chatMemberCache cache.ChatMemberCache) DistService {
	chatMemberClient := chat_member_client.NewChatMemberClient(cfg.Etcd, cfg.ChatMemberServer, cfg.GrpcServer.Jaeger, cfg.Name)
	svc := &distService{
		cfg:              cfg,
		msgHandle:        make(map[string]obj.KafkaMessageHandler),
		clients:          make(map[int64]gw_client.MessageGatewayClient),
		producers:        map[int64]*xkafka.Producer{},
		chatMemberClient: chatMemberClient,
		serverMgrCache:   serverMgrCache,
		chatMemberCache:  chatMemberCache,
		queues:           make(chan struct{}, 8),
	}

	svc.msgHandle[cfg.MsgConsumer.Topic[0]] = svc.MessageHandler

	svc.consumerGroup = xkafka.NewMConsumerGroup(&xkafka.MConsumerGroupConfig{KafkaVersion: sarama.V3_2_1_0, OffsetsInitial: sarama.OffsetNewest, IsReturnErr: false},
		cfg.MsgConsumer.Topic,
		cfg.MsgConsumer.Address,
		cfg.MsgConsumer.GroupID)
	svc.consumerGroup.RegisterHandler(svc)

	svc.watchMsgGatewayServer()

	return svc
}

func (s *distService) Setup(_ sarama.ConsumerGroupSession) error {
	close(s.consumerGroup.Ready)
	return nil
}
func (s *distService) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (s *distService) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
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
