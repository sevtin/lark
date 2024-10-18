package client

import (
	"fmt"
	"github.com/spf13/cast"
	"lark/pkg/common/xlog"
	"lark/pkg/common/xredis"
	"lark/pkg/common/xtimer"
	"lark/pkg/constant"
	"lark/pkg/obj"
	"lark/pkg/proto/pb_chat_member"
	"lark/pkg/proto/pb_enum"
	"lark/pkg/utils"
	"log"
	"strconv"
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
	FirstUid    int64 // First user id
	OnlineCount int64 // Online users
	SendCount   int64 // Number of messages send
	MemberCount int64 // Number of group members
	TestCount   int64 // Number of tests
	ServerCount int   // Number of Servers
	ChatId      int64
	Cluster     bool
}

func NewManager(firstUid int64, onlineCount int64, sendCount int64, memberCount int64, testCount int64, chatId int64, cluster bool, serverCount int) (mgr *Manager) {
	mgr = &Manager{
		unregister:  make(chan *Client, 1000),
		clients:     obj.NewRwMap(),
		FirstUid:    firstUid,
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
		key        string
		err        error
		memberInfo *pb_chat_member.ChatMemberInfo
		jsonStr    string
		start      = m.FirstUid
		end        = start + m.MemberCount
		status     = map[int64]map[string]string{}
		servers    = map[int64]map[string]string{}
		infos      = map[string]string{}
		sid        = utils.NewServerId(0, 1, 1)
	)

	m.debug()

	for uid = start; uid < end; uid++ {
		uidStr = cast.ToString(uid)
		// 0:ServerId, 1:Platform, 2:Uid, 3:Status
		memberInfo = &pb_chat_member.ChatMemberInfo{
			ChatId:       m.ChatId,
			ChatType:     pb_enum.CHAT_TYPE_GROUP,
			Uid:          uid,
			Alias:        "昵称:" + cast.ToString(uid),
			MemberAvatar: "https://photos.bucket.lark.com/b11883ba-f3d7-4164-a593-700c177c37c8",
			RoleId:       1,
		}
		jsonStr, _ = utils.Marshal(memberInfo)

		// 成员列表
		cs := uid % constant.MAX_CHAT_SLOT
		if _, ok := status[cs]; ok == false {
			status[cs] = make(map[string]string)
		}
		status[cs][uidStr] = "0"

		us := uid % constant.MAX_USER_SLOT
		if _, ok := servers[us]; ok == false {
			servers[us] = make(map[string]string)
		}
		servers[us][uidStr] = cast.ToString(sid)
		infos[uidStr] = jsonStr
	}

	for r, kvs := range status {
		key = constant.RK_SYNC_DIST_CHAT_MEMBER_HASH + utils.GetHashTagKey(m.ChatId) + ":" + strconv.Itoa(int(r))
		err = xredis.HMSet(key, kvs)
		if err != nil {
			xlog.Error(err.Error())
			return
		}
	}
	for r, kvs := range servers {
		key = constant.RK_SYNC_USER_SERVER + strconv.Itoa(int(r))
		err = xredis.HMSet(key, kvs)
		if err != nil {
			xlog.Error(err.Error())
			return
		}
	}
	key = constant.RK_SYNC_CHAT_MEMBER_INFO_HASH + utils.GetHashTagKey(m.ChatId)
	err = xredis.HMSet(key, infos)
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
		start  = m.FirstUid
		end    = start + count
	)
	ch := make(chan int, 1000)
	for i = start; i < end; i++ {
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
		m.newConnection(ch, i, server+":"+cast.ToString(port))
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
			start  = m.FirstUid
			end    = start + m.SendCount
		)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				if isStart == false {
					isStart = true
					msgTimer.Run()
				}
				for i = start; i < end; i++ {
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
