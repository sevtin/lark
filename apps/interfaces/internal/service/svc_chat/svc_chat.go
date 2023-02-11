package svc_chat

import (
	chat_client "lark/apps/chat/client"
	"lark/apps/interfaces/internal/config"
	"lark/apps/interfaces/internal/dto/dto_chat"
	"lark/pkg/xhttp"
)

type ChatService interface {
	CreateGroupChat(params *dto_chat.CreateGroupChatReq, uid int64) (resp *xhttp.Resp)
	EditGroupChat(params *dto_chat.EditGroupChatReq, uid int64) (resp *xhttp.Resp)
	DeleteContact(params *dto_chat.DeleteContactReq, uid int64) (resp *xhttp.Resp)
	RemoveGroupChatMember(params *dto_chat.RemoveGroupChatMemberReq, uid int64) (resp *xhttp.Resp)
	QuitGroupChat(params *dto_chat.QuitGroupChatReq, uid int64) (resp *xhttp.Resp)
	GroupChatDetails(params *dto_chat.GroupChatDetailsReq) (resp *xhttp.Resp)
}

type chatService struct {
	chatClient chat_client.ChatClient
}

func NewChatService() ChatService {
	conf := config.GetConfig()
	chatClient := chat_client.NewChatClient(conf.Etcd, conf.ChatServer, conf.Jaeger, conf.Name)
	return &chatService{chatClient: chatClient}
}
