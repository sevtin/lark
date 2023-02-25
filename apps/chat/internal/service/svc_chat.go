package service

import (
	"context"
	chat_client "lark/apps/chat/client"
	"lark/apps/chat/internal/config"
	dist_client "lark/apps/dist/client"
	user_client "lark/apps/user/client"
	"lark/domain/cache"
	"lark/domain/repo"
	"lark/pkg/common/xkafka"
	"lark/pkg/proto/pb_chat"
)

type ChatService interface {
	CreateGroupChat(ctx context.Context, req *pb_chat.CreateGroupChatReq) (resp *pb_chat.CreateGroupChatResp, err error)
	EditGroupChat(ctx context.Context, req *pb_chat.EditGroupChatReq) (resp *pb_chat.EditGroupChatResp, err error)
	GroupChatDetails(ctx context.Context, req *pb_chat.GroupChatDetailsReq) (resp *pb_chat.GroupChatDetailsResp, err error)
	RemoveGroupChatMember(ctx context.Context, req *pb_chat.RemoveGroupChatMemberReq) (resp *pb_chat.RemoveGroupChatMemberResp, err error)
	QuitGroupChat(ctx context.Context, req *pb_chat.QuitGroupChatReq) (resp *pb_chat.QuitGroupChatResp, err error)
	DeleteContact(ctx context.Context, req *pb_chat.DeleteContactReq) (resp *pb_chat.DeleteContactResp, err error)
	UploadAvatar(ctx context.Context, req *pb_chat.UploadAvatarReq) (resp *pb_chat.UploadAvatarResp, err error)
	GetChatInfo(ctx context.Context, req *pb_chat.GetChatInfoReq) (resp *pb_chat.GetChatInfoResp, err error)
}

type chatService struct {
	cfg              *config.Config
	chatRepo         repo.ChatRepository
	chatInviteRepo   repo.ChatInviteRepository
	chatMemberRepo   repo.ChatMemberRepository
	userRepo         repo.UserRepository
	avatarRepo       repo.AvatarRepository
	chatCache        cache.ChatCache
	chatMessageCache cache.ChatMessageCache
	chatMemberCache  cache.ChatMemberCache
	userCache        cache.UserCache
	chatClient       chat_client.ChatClient
	userClient       user_client.UserClient
	distClient       dist_client.DistClient
	producer         *xkafka.Producer
	cacheProducer    *xkafka.Producer
}

func NewChatService(conf *config.Config,
	chatRepo repo.ChatRepository,
	chatInviteRepo repo.ChatInviteRepository,
	chatMemberRepo repo.ChatMemberRepository,
	userRepo repo.UserRepository,
	avatarRepo repo.AvatarRepository,
	chatCache cache.ChatCache,
	chatMessageCache cache.ChatMessageCache,
	chatMemberCache cache.ChatMemberCache,
	userCache cache.UserCache) ChatService {
	svc := &chatService{cfg: conf,
		chatRepo:         chatRepo,
		chatInviteRepo:   chatInviteRepo,
		chatMemberRepo:   chatMemberRepo,
		userRepo:         userRepo,
		avatarRepo:       avatarRepo,
		chatCache:        chatCache,
		chatMessageCache: chatMessageCache,
		chatMemberCache:  chatMemberCache,
		userCache:        userCache,
	}
	svc.chatClient = chat_client.NewChatClient(conf.Etcd, conf.ChatServer, conf.Jaeger, conf.Name)
	svc.userClient = user_client.NewUserClient(conf.Etcd, conf.UserServer, conf.Jaeger, conf.Name)
	svc.distClient = dist_client.NewDistClient(conf.Etcd, conf.DistServer, conf.Jaeger, conf.Name)
	svc.producer = xkafka.NewKafkaProducer(conf.MsgProducer.Address, conf.MsgProducer.Topic)
	svc.cacheProducer = xkafka.NewKafkaProducer(conf.CacheProducer.Address, conf.CacheProducer.Topic)
	return svc
}
