package service

import (
	"context"
	"lark/apps/chat_member/internal/config"
	user_client "lark/apps/user/client"
	"lark/domain/cache"
	"lark/domain/repo"
	"lark/pkg/common/xkafka"
	"lark/pkg/proto/pb_chat_member"
)

type ChatMemberService interface {
	GetChatMemberInfo(ctx context.Context, req *pb_chat_member.GetChatMemberInfoReq) (resp *pb_chat_member.GetChatMemberInfoResp, err error)
	GetDistMemberList(ctx context.Context, req *pb_chat_member.GetDistMemberListReq) (resp *pb_chat_member.GetDistMemberListResp, err error)
	GetChatMemberList(ctx context.Context, req *pb_chat_member.GetChatMemberListReq) (resp *pb_chat_member.GetChatMemberListResp, err error)
	GetContactList(ctx context.Context, req *pb_chat_member.GetContactListReq) (resp *pb_chat_member.GetContactListResp, err error)
	GetGroupChatList(ctx context.Context, req *pb_chat_member.GetGroupChatListReq) (resp *pb_chat_member.GetGroupChatListResp, err error)
}

type chatMemberService struct {
	cfg             *config.Config
	chatMemberRepo  repo.ChatMemberRepository
	userRepo        repo.UserRepository
	userClient      user_client.UserClient
	chatMemberCache cache.ChatMemberCache
	userCache       cache.UserCache
	producer        *xkafka.Producer
}

func NewChatMemberService(
	cfg *config.Config,
	chatMemberRepo repo.ChatMemberRepository,
	userRepo repo.UserRepository,
	chatMemberCache cache.ChatMemberCache,
	userCache cache.UserCache) ChatMemberService {

	svc := &chatMemberService{
		cfg:             cfg,
		chatMemberRepo:  chatMemberRepo,
		userRepo:        userRepo,
		chatMemberCache: chatMemberCache,
		userCache:       userCache}
	svc.userClient = user_client.NewUserClient(cfg.Etcd, cfg.UserServer, cfg.GrpcServer.Jaeger, cfg.Name)
	svc.producer = xkafka.NewKafkaProducer(cfg.MsgProducer.Address, cfg.MsgProducer.Topic)
	return svc
}
