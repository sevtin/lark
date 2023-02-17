package client

import (
	"fmt"
	"lark/pkg/common/xlog"
	"lark/pkg/common/xredis"
	"lark/pkg/common/xtimer"
	"lark/pkg/constant"
	"lark/pkg/obj"
	"lark/pkg/proto/pb_chat_member"
	"lark/pkg/proto/pb_enum"
	"lark/pkg/utils"
	"log"
	"time"
)

var (
	msgTimer             = xtimer.NewTimer(2 * time.Second)
	isStart              bool
	receivedMessageCount int64
)

type Manager struct {
	unregister  chan *Client
	clients     *obj.RwMap
	OnlineCount int64 // Online users
	SendCount   int64 // Number of messages send
	MemberCount int64 // Number of group members
	TestCount   int64 // Number of tests
	ServerCount int   // Number of Servers
	ChatId      int64
	Cluster     bool
}

func NewManager(onlineCount int64, sendCount int64, memberCount int64, testCount int64, chatId int64, cluster bool, serverCount int) (mgr *Manager) {
	mgr = &Manager{
		unregister:  make(chan *Client, 1000),
		clients:     obj.NewRwMap(),
		OnlineCount: onlineCount,
		SendCount:   sendCount,
		MemberCount: memberCount,
		TestCount:   testCount,
		ChatId:      chatId,
		Cluster:     cluster,
		ServerCount: serverCount,
	}
	return
}

func (m *Manager) Run() {
	var (
		uid        int64
		uidStr     string
		kv1        = map[string]interface{}{}
		kv2        = map[string]interface{}{}
		key        string
		err        error
		sid        int64
		memberInfo *pb_chat_member.ChatMemberInfo
		jsonStr    string
	)

	m.debug()

	for uid = 1; uid <= m.MemberCount; uid++ {
		uidStr = utils.Int64ToStr(uid)
		// 0:ServerId, 1:Platform, 2:Uid, 3:Status
		if m.Cluster == true {
			sid = uid % int64(m.ServerCount)
			switch sid {
			case 1:
				sid = 20000
			case 2:
				sid = 30000
			default:
				sid = 10000
			}
			kv1[uidStr] = fmt.Sprintf("%d,%d,%d", sid, uid, 0)
		} else {
			kv1[uidStr] = fmt.Sprintf("%d,%d,%d", 10000, uid, 0)
		}

		memberInfo = &pb_chat_member.ChatMemberInfo{
			ChatId:          m.ChatId,
			ChatType:        pb_enum.CHAT_TYPE_GROUP,
			Uid:             uid,
			Alias:           "昵称:" + utils.Int64ToStr(uid),
			MemberAvatarKey: "b11883ba-f3d7-4164-a593-700c177c37c8.jpeg",
			RoleId:          1,
		}
		jsonStr, _ = utils.Marshal(memberInfo)
		kv2[uidStr] = jsonStr
	}
	key = constant.RK_SYNC_DIST_CHAT_MEMBER_HASH + utils.Int64ToStr(m.ChatId)
	err = xredis.HMSet(key, kv1)
	if err != nil {
		xlog.Error(err.Error())
		return
	}
	key = constant.RK_SYNC_CHAT_MEMBER_INFO_HASH + utils.Int64ToStr(m.ChatId)
	err = xredis.HMSet(key, kv2)
	if err != nil {
		xlog.Error(err.Error())
		return
	}

	m.runLoop()
	m.batchCreate(m.OnlineCount)
}

func (m *Manager) unregisterClient(client *Client) {
	m.clients.Delete(client.uid)
}

func (m *Manager) runLoop() {
	go func() {
		var (
			client *Client
		)
		for {
			select {
			case client = <-m.unregister:
				m.unregisterClient(client)
			}
		}
	}()
}

func (m *Manager) batchCreate(count int64) {
	var (
		i      int64
		server = "127.0.0.1"
		port   = 7301
		sid    int64
	)
	ch := make(chan int, 1000)
	for i = 1; i <= count; i++ {
		ch <- 0
		if m.Cluster == true {
			sid = i % int64(m.ServerCount)
			switch sid {
			case 1:
				sid = 2
				port = 7311
			case 2:
				sid = 3
				port = 7321
			default:
				sid = 1
				port = 7301
			}
		}
		m.newConnection(ch, i, server+":"+utils.IntToStr(port))
		//server, _, port = m.getServer()
	}
	close(ch)

	fmt.Println("准备发送消息:", time.Now())
	time.Sleep(5 * time.Second)
	fmt.Println("开始发送消息:", time.Now())
	m.loopSend()
}

func (m *Manager) getServer() (server string, serverId int64, port int) {
	list := xredis.ZRevRange(constant.RK_SYNC_SERVER_MSG_GATEWAY, 0, 0)
	if len(list) == 0 {
		return
	}
	member := list[0]
	server, serverId, _ = utils.GetMsgGatewayServer(member)
	return
}

func (m *Manager) loopSend() {
	go func() {
		var (
			i      int64
			count  int64
			ticker = time.NewTicker(time.Second * 1)
		)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				if isStart == false {
					isStart = true
					msgTimer.Run()
				}
				for i = 1; i <= m.SendCount; i++ {
					client, ok := m.clients.Get(i)
					if ok == false {
						continue
					}
					client.(*Client).SendMsg(m.ChatId)
				}
				count++
				if count >= m.TestCount {
					return
				}
			}
		}
	}()
}

func (m *Manager) newConnection(ch chan int, uid int64, host string) {
	go func() {
		var (
			client *Client
		)
		client = NewClient(uid, m, host)
		if client.conn != nil {
			m.clients.Set(uid, client)
		}
		<-ch
	}()
}

func (m *Manager) debug() {
	go func() {
		allTicker := time.NewTicker(time.Second * 5)
		defer allTicker.Stop()
		for {
			select {
			case <-allTicker.C:
				log.Println("在线人数:", m.clients.Len(), " 接收消息数量:", receivedMessageCount)
			}
		}
	}()
}
