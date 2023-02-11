package client

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
	"lark/domain/cache"
	"lark/pkg/common/xjwt"
	"lark/pkg/common/xlog"
	"lark/pkg/common/xsnowflake"
	"lark/pkg/constant"
	"lark/pkg/proto/pb_enum"
	"lark/pkg/proto/pb_msg"
	"lark/pkg/utils"
	"net/url"
	"runtime/debug"
	"sync"
	"sync/atomic"
	"time"
)

var (
	authCache = cache.NewAuthCache()
)

type Client struct {
	rwLock   sync.RWMutex
	mgr      *Manager
	conn     *websocket.Conn
	uid      int64 // 用户ID
	platform int32 // 平台ID
	key      string
	onlineAt int64 // 上线时间戳（毫秒）
	sendChan chan []byte
	closed   bool
	nickname string
}

func NewClient(uid int64, mgr *Manager, host string) (client *Client) {
	var (
		u url.URL
		//q      url.Values
		ts     int64
		conn   *websocket.Conn
		header map[string][]string
		token  *xjwt.JwtToken
		err    error
	)
	ts = time.Now().Unix()
	client = &Client{
		mgr:      mgr,
		conn:     nil,
		uid:      uid,
		platform: 1,
		onlineAt: ts,
		sendChan: make(chan []byte, WS_CHAN_CLIENT_WRITE_MESSAGE_SIZE),
		closed:   false,
		nickname: "昵称:" + utils.Int64ToStr(uid),
	}

	u = url.URL{Scheme: "ws", Host: host, Path: "/"}
	/*
		q := u.Query()
		q.Set("uid", uid)
		q.Set("platform", "1")
		u.RawQuery = q.Encode()
	*/

	token, _ = xjwt.CreateToken(client.uid, client.platform, true, constant.CONST_DURATION_SHA_JWT_ACCESS_TOKEN_EXPIRE_IN_SECOND)
	header = make(map[string][]string)
	header[WS_KEY_UID] = []string{utils.Int64ToStr(uid)}
	header[WS_KEY_PLATFORM] = []string{"1"}
	header[WS_KEY_COOKIE] = []string{token.Token}
	err = authCache.SetAccessTokenSessionId(client.uid, client.platform, token.SessionId)
	if err != nil {
		client.closed = true
		fmt.Println("缓存会话ID失败", err.Error())
		return
	}
	conn, _, err = websocket.DefaultDialer.Dial(u.String(), header)
	if err != nil {
		client.closed = true
		fmt.Println("创建连接失败", err.Error())
		return
	}
	client.conn = conn
	client.listen()
	return
}

func (c *Client) listen() {
	go c.write()
	go c.read()
}

func (c *Client) closeConn() {
	c.rwLock.Lock()
	if c.closed == true {
		c.rwLock.Unlock()
		return
	}
	c.closed = true
	close(c.sendChan)
	c.rwLock.Unlock()
	c.mgr.unregister <- c
	c.conn.Close()
}

func (c *Client) read() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r, string(debug.Stack()))
		}
		c.closeConn()
	}()

	var (
		msgType int
		buf     []byte
		err     error
	)

	c.conn.SetReadLimit(WS_MAX_MESSAGE_SIZE)
	c.conn.SetReadDeadline(time.Now().Add(WS_PONG_WAIT))
	c.conn.SetPongHandler(c.pongHandler)
	//c.conn.SetPingHandler(c.pingHandler)
	c.conn.SetCloseHandler(c.closeHandler)

	for {
		if msgType, buf, err = c.conn.ReadMessage(); err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
			}
			break
		}
		if c.closed == true {
			break
		}
		if msgType == websocket.CloseMessage {
			break
		}
		if msgType != websocket.BinaryMessage {
			continue
		}
		if len(buf) == 0 {
			continue
		}
		c.decode(buf)
	}
}

func (c *Client) pongHandler(appData string) (err error) {
	err = c.conn.SetReadDeadline(time.Now().Add(WS_PONG_WAIT))
	if err != nil {
		xlog.Warn(err.Error())
		c.closeConn()
	}
	return
}

func (c *Client) pingHandler(appData string) (err error) {
	err = c.conn.WriteControl(websocket.PongMessage, []byte(appData), time.Now().Add(time.Second))
	/*
		if err == websocket.ErrCloseSent {
			return
		} else if e, ok := err.(net.Error); ok && e.Temporary() {
			return
		}
	*/
	if err != nil {
		xlog.Warn(err.Error())
		c.closeConn()
	}
	return
}

func (c *Client) closeHandler(code int, text string) (err error) {
	c.closeConn()
	// message := websocket.FormatCloseMessage(code, "")
	// c.conn.WriteControl(websocket.CloseMessage, message, time.Now().Add(time.Second))
	return
}

func (c *Client) decode(buf []byte) {
	msgTimer.UpdateEndTime()
	var (
		msg    *pb_msg.Packet
		buffer = bytes.NewBuffer(buf)
	)
	for {
		atomic.AddInt64(&receivedMessageCount, 1)
		msg = utils.Decode(buffer, true)
		if buffer.Len() < utils.SrvMessageHeaderLen {
			break
		}
		if msg.MsgType == pb_enum.MESSAGE_TYPE_NEW {

		}
	}
}

func (c *Client) oldDecode(buf []byte) {
	msgTimer.UpdateEndTime()
	var (
		lengthBuff   []byte
		length       uint32
		topicBuff    []byte
		topic        uint32
		subtopicBuff []byte
		subtopic     uint32
		msgTypeBuff  []byte
		msgType      uint32
		body         []byte
		totalLength  uint32
	)
	var (
	// resp      *pb_msg.MessageResp
	// msg       *pb_msg.SrvChatMessage
	)
	for {
		totalLength = uint32(len(buf))
		if totalLength < MessageType {
			break
		}
		lengthBuff = buf[:MessageLength]
		length = binary.LittleEndian.Uint32(lengthBuff)
		if totalLength < MessageType+length {
			break
		}
		topicBuff = buf[MessageLength:MessageTopic]
		topic = binary.LittleEndian.Uint32(topicBuff)

		subtopicBuff = buf[MessageTopic:MessageSubtopic]
		subtopic = binary.LittleEndian.Uint32(subtopicBuff)

		msgTypeBuff = buf[MessageSubtopic:MessageType]
		msgType = binary.LittleEndian.Uint32(msgTypeBuff)

		body = buf[MessageType : MessageType+length]
		if length > 0 && topic > 0 && subtopic > 0 && msgType > 0 && len(body) > 0 {

		}
		atomic.AddInt64(&receivedMessageCount, 1)
		/*
			switch pb_enum.MESSAGE_TYPE(msgType) {
			case pb_enum.MESSAGE_TYPE_RESP:
				resp = new(pb_msg.MessageResp)
				proto.Unmarshal(body, resp)
				xlog.Info("应答消息:", topic, subtopic, resp.Code, resp.Msg)
			case pb_enum.MESSAGE_TYPE_NEW:
				msg = new(pb_msg.SrvChatMessage)
				proto.Unmarshal(body, msg)
				xlog.Info("新消息:", topic, subtopic, msg.SeqId, string(msg.Body))
			}
		*/
		if totalLength < MessageType+length+MessageType {
			break
		}
		buf = buf[MessageType+length:]
	}
}

func (c *Client) write() {
	pingTicker := time.NewTicker(WS_PING_PERIOD)
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r, string(debug.Stack()))
		}
		pingTicker.Stop()
		c.closeConn()
	}()

	var (
		err     error
		message []byte
		ok      bool
	)

	for {
		select {
		case message, ok = <-c.sendChan:
			if ok == false {
				return
			}
			if c.closed == true {
				return
			}
			if err = c.conn.SetWriteDeadline(time.Now().Add(WS_WRITE_WAIT)); err != nil {
				return
			}
			if err = c.conn.WriteMessage(websocket.BinaryMessage, message); err != nil {
				return
			}
		case <-pingTicker.C:
			if err = c.conn.SetWriteDeadline(time.Now().Add(WS_WRITE_WAIT)); err != nil {
				return
			}
			if err = c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func (c *Client) SendMsg(chatId int64) (err error) {
	var (
		ts      int64
		msgBuf  []byte
		msgBody *pb_msg.CliChatMessage
	)
	if c.conn == nil {
		return
	}
	if c.closed == true {
		return
	}
	ts = utils.MillisFromTime(time.Now())
	msgBody = &pb_msg.CliChatMessage{
		CliMsgId: xsnowflake.NewSnowflakeID(), //客户端唯一消息号
		ChatId:   chatId,
		MsgType:  1,
		Body:     utils.Str2Bytes("文本聊天消息" + utils.Int64ToStr(c.uid)),
		SentTs:   ts,
	}
	msgBuf, _ = proto.Marshal(msgBody)
	msgBuf, _ = utils.EncodeCliMessage(int32(pb_enum.TOPIC_CHAT), int32(pb_enum.SUB_TOPIC_CHAT_MSG), msgBody.CliMsgId, msgBuf)
	c.Send(msgBuf)
	return
}

func (c *Client) Send(message []byte) {
	defer func() {
		if r := recover(); r != nil {
			xlog.Warn(r, string(debug.Stack()))
		}
	}()
	c.rwLock.RLock()
	if c.closed == true {
		c.rwLock.RUnlock()
		return
	}
	c.rwLock.RUnlock()
	if len(c.sendChan) >= WS_CHAN_CLIENT_WRITE_MESSAGE_THRESHOLD {
		return
	}
	c.sendChan <- message
}

func (c *Client) Close() (err error) {
	c.rwLock.RLock()
	if c.closed == true {
		c.rwLock.RUnlock()
		return
	}
	c.rwLock.RUnlock()
	err = c.conn.WriteMessage(websocket.CloseMessage, WS_MSG_BUF_CLOSE)
	return
}
