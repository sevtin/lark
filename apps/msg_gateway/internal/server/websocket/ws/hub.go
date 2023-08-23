package ws

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
	"lark/pkg/proto/pb_msg"
	"lark/pkg/utils"
	"log"
	"runtime/debug"
	"strconv"
	"time"
)

type Hub struct {
	serverId       int
	upgrader       websocket.Upgrader
	registerChan   chan *Client
	unregisterChan chan *Client    // 只在Client调用closeConn()函数时触发
	readChan       chan *Message   // 客户端发送的消息
	msgCallback    MessageCallback // 回调
	clients        *CliMap         // key:platform-uid
	now            time.Time
}

func NewHub(serverId int, msgCallback MessageCallback) *Hub {
	return &Hub{
		serverId: serverId,
		upgrader: websocket.Upgrader{
			ReadBufferSize:    WS_READ_BUFFER_SIZE,
			WriteBufferSize:   WS_WRITE_BUFFER_SIZE,
			EnableCompression: false, //关闭压缩
		},
		registerChan:   make(chan *Client, WS_CHAN_CLIENT_REGISTER_SIZE),
		unregisterChan: make(chan *Client, WS_CHAN_CLIENT_UNREGISTER_SIZE),
		readChan:       make(chan *Message, WS_CHAN_SERVER_READ_MESSAGE_SIZE),
		msgCallback:    msgCallback,
		clients:        NewCliMap(),
		now:            time.Now(),
	}
}

func (h *Hub) registerClient(client *Client) {
	var (
		ok  bool
		cli *Client
	)

	if cli, ok = h.clients.Get(client.key); ok == false {
		h.clients.Set(client.key, client)
		h.OnOffline(true)
		return
	}

	if client.onlineTs > cli.onlineTs {
		h.clients.Set(client.key, client)
		cli.Close()
		return
	}
	// TODO:踢出消息
	client.Close()
}

func (h *Hub) unregisterClient(client *Client) {
	var (
		ok  bool
		cli *Client
	)
	if cli, ok = h.clients.Get(client.key); ok == false {
		return
	}
	if client == cli {
		h.clients.Delete(client.key)
		h.OnOffline(false)
	}
}

func (h *Hub) OnOffline(onOff bool) {
	if onOff {
		log.Println("R在线用户数量:", h.clients.Len())
	} else {
		log.Println("U在线用户数量:", h.clients.Len())
	}
}

func (h *Hub) coroutine() {
	go func() {
		var (
			ticker = time.NewTicker(time.Millisecond * 500)
			now    time.Time
			ok     bool
		)
		defer ticker.Stop()
		for {
			select {
			case now, ok = <-ticker.C:
				if ok == false {
					return
				}
				h.now = now
			}
		}
	}()

	go func() {
		var (
			client *Client
		)
		for {
			select {
			case client = <-h.registerChan:
				h.registerClient(client)
			case client = <-h.unregisterChan:
				h.unregisterClient(client)
			}
		}
	}()
}

func (h *Hub) Run() {
	defer func() {
		if r := recover(); r != nil {
			wsLog.Error(r, string(debug.Stack()))
		}
	}()
	h.coroutine()
	// 调试用
	h.debug()

	var (
		index int
		loop  = WS_MAX_CONSUMER_SIZE
	)
	for index = 0; index < loop; index++ {
		go func() {
			var (
				msg *Message
			)
			for {
				select {
				case msg = <-h.readChan:
					h.msgCallback(msg)
				}
			}
		}()
	}
}

func (h *Hub) IsOnline(uid int64, platform int32) (ok bool) {
	_, ok = h.clients.Get(clientKey(uid, platform))
	return
}

func (h *Hub) SendMessage(uid int64, platform int32, message []byte) (result int32) {
	result = WS_CLIENT_OFFLINE
	var (
		cli *Client
		ok  bool
	)
	if cli, ok = h.clients.Get(clientKey(uid, platform)); ok == false {
		return
	}
	cli.Send(message)
	result = WS_SEND_MSG_SUCCESS
	return
}

func (h *Hub) NumberOfOnline() int {
	return h.clients.Len()
}

func (h *Hub) broadcastMessage() {
	var (
		msgBuf, _       = proto.Marshal(&pb_msg.SrvChatMessage{})
		broadcastBuf, _ = utils.Encode(1, 0, 0, msgBuf)
		uid             int64
	)
	for uid = 1; uid <= 10000; uid++ {
		h.SendMessage(uid, 1, broadcastBuf)
	}
}

func (h *Hub) debug() {
	go func() {
		ticker := time.NewTicker(time.Second * 30)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				log.Println("在线人数:", h.clients.Len())
			}
		}
	}()
}

func (h *Hub) wsHandler(c *gin.Context) {
	var (
		uidVal   interface{}
		pidVal   interface{}
		exists   bool
		uid      int64
		platform int32
		conn     *websocket.Conn
		client   *Client
		err      error
	)

	if h.clients.Len() >= WS_MAX_CONNECTIONS {
		httpError(c, ERROR_CODE_WS_EXCEED_MAX_CONNECTIONS, ERROR_WS_EXCEED_MAX_CONNECTIONS)
		return
	}
	uidVal, exists = c.Get(WS_KEY_UID)
	if exists == false {
		httpError(c, ERROR_CODE_HTTP_UID_DOESNOT_EXIST, ERROR_HTTP_UID_DOESNOT_EXIST)
		return
	}
	pidVal, exists = c.Get(WS_KEY_PLATFORM)
	if exists == false {
		httpError(c, ERROR_CODE_HTTP_PLATFORM_DOESNOT_EXIST, ERROR_HTTP_PLATFORM_DOESNOT_EXIST)
		return
	}
	uid, _ = strconv.ParseInt(uidVal.(string), 10, 64)
	if uid == 0 {
		httpError(c, ERROR_CODE_HTTP_UID_DOESNOT_EXIST, ERROR_HTTP_UID_DOESNOT_EXIST)
		return
	}
	platform = int32(pidVal.(float64))

	if conn, err = h.upgrader.Upgrade(c.Writer, c.Request, nil); err != nil {
		// 协议升级失败
		httpError(c, ERROR_CODE_HTTP_UPGRADER_FAILED, err.Error())
		wsLog.Warn(err.Error())
		return
	}
	client = newClient(h, conn, uid, platform)
	client.listen()
	h.registerChan <- client
}
