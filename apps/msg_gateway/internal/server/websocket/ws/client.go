package ws

import (
	"bytes"
	"context"
	"github.com/gorilla/websocket"
	"io"
	"lark/pkg/utils"
	"log"
	"runtime/debug"
	"sync"
	"time"
)

type Client struct {
	rwLock   sync.RWMutex
	hub      *Hub
	conn     *websocket.Conn
	uid      int64 // 用户ID
	platform int32 // 平台ID
	key      string
	onlineTs int64 // 上线时间戳（毫秒）
	sendChan chan []byte
	closed   bool
	ctx      context.Context
	cancel   context.CancelFunc
}

func newClient(hub *Hub, conn *websocket.Conn, uid int64, platform int32) *Client {
	cli := &Client{
		hub:      hub,
		conn:     conn,
		uid:      uid,
		platform: platform,
		key:      clientKey(uid, platform),
		onlineTs: time.Now().UnixNano() / 1e6,
		sendChan: make(chan []byte, WS_WRITE_MAX_MESSAGE_CHAN_SIZE),
	}
	cli.ctx, cli.cancel = context.WithCancel(context.Background())
	//cli.debug()
	return cli
}

func (c *Client) GetUid() int64 {
	return c.uid
}

func (c *Client) GetPlatform() int32 {
	return c.platform
}

func (c *Client) listen() {
	go c.writeLoop()
	go c.readLoop()
}

func (c *Client) debug() {
	go func() {
		var (
			ticker = time.NewTicker(time.Second * 30)
			ok     bool
		)
		defer ticker.Stop()
		for {
			select {
			case _, ok = <-ticker.C:
				if ok == false {
					return
				}
				log.Println(c.uid, c.closed, len(c.sendChan))
				if c.closed == true {
					ticker.Stop()
					return
				}
			}
		}
	}()
}

func (c *Client) closeConn() {
	c.rwLock.Lock()
	if c.closed == true {
		c.rwLock.Unlock()
		return
	}
	c.closed = true
	c.rwLock.Unlock()
	c.hub.unregisterClient(c)
	c.cancel()
	c.conn.Close()
}

func (c *Client) readLoop() {
	defer func() {
		c.closeConn()
		if r := recover(); r != nil {
			wsLog.Warn(r, string(debug.Stack()))
		}
	}()

	var (
		msgType int
		buf     []byte
		buffer  *bytes.Buffer
		message *Message
		err     error
	)

	c.conn.SetReadLimit(WS_READ_MAX_MESSAGE_BUFFER_SIZE)
	c.conn.SetReadDeadline(time.Now().Add(WS_PONG_WAIT))
	c.conn.SetPongHandler(c.pongHandler)
	//c.conn.SetPingHandler(c.pingHandler)
	c.conn.SetCloseHandler(c.closeHandler)

	for {
		if msgType, buf, err = c.conn.ReadMessage(); err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
			}
			switch err.(type) {
			case *websocket.CloseError:
			default:
				wsLog.Warn(err.Error())
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
		// 限流处理
		if len(c.hub.readChan) >= WS_CHAN_SERVER_READ_MESSAGE_THRESHOLD {
			// c.overloadReply(bufMsg)
			continue
		}
		buffer = bytes.NewBuffer(buf)
		message = &Message{
			Uid:      c.uid,
			Platform: c.platform,
			Hub:      c.hub,
			Packet:   utils.Decode(buffer, false),
		}
		//c.hub.messageHandler(message)
		c.hub.readChan <- message
	}
}

func (c *Client) pongHandler(appData string) (err error) {
	err = c.conn.SetReadDeadline(time.Now().Add(WS_PONG_WAIT))
	if err != nil {
		wsLog.Warn(err.Error())
		c.closeConn()
	}
	return
}

func (c *Client) pingHandler(appData string) (err error) {
	err = c.conn.WriteControl(websocket.PongMessage, WS_MSG_BUF_PONG, time.Now().Add(time.Second))
	/*
		if err == websocket.ErrCloseSent {
			return
		} else if e, ok := err.(net.Error); ok && e.Temporary() {
			return
		}
	*/
	if err != nil {
		wsLog.Warn(err.Error())
		c.closeConn()
	}
	return
}

func (c *Client) closeHandler(code int, text string) (err error) {
	c.closeConn()
	//message := websocket.FormatCloseMessage(code, "")
	//c.conn.WriteControl(websocket.CloseMessage, message, time.Now().Add(time.Second))
	return
}

func (c *Client) writeLoop() {
	pingTicker := time.NewTicker(WS_PING_PERIOD)
	defer func() {
		pingTicker.Stop()
		c.closeConn()
		if r := recover(); r != nil {
			wsLog.Warn(r, string(debug.Stack()))
		}
	}()

	var (
		message []byte
		ok      bool
		wc      io.WriteCloser
		chLen   int
		bufLen  int
		index   int
		merges  int
		err     error
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
			merges = 1
			if err = c.conn.SetWriteDeadline(time.Now().Add(WS_WRITE_WAIT)); err != nil {
				wsLog.Warn(err.Error())
				return
			}
			/*
				if err = c.conn.WriteMessage(websocket.BinaryMessage, message); err != nil {
					return
				}
			*/
			wc, err = c.conn.NextWriter(websocket.BinaryMessage)
			if err != nil {
				wsLog.Warn(err.Error())
				return
			}
			bufLen = len(message)
			_, err = wc.Write(message)
			if err != nil {
				wsLog.Warn(err.Error())
				return
			}
			if bufLen < WS_WRITE_MAX_MERGE_MESSAGE_BUFFER_SIZE {
				chLen = len(c.sendChan)
				for index = 0; index < chLen; index++ {
					if bufLen >= WS_WRITE_MAX_MESSAGE_BUFFER_SIZE {
						break
					}
					merges++
					if merges > WS_WRITE_MAX_MERGE_MESSAGE_SIZE {
						break
					}
					message = <-c.sendChan
					bufLen += len(message)
					_, err = wc.Write(message)
					if err != nil {
						wsLog.Warn(err.Error())
						return
					}
				}
			}

			if err = wc.Close(); err != nil {
				wsLog.Warn(err.Error())
				return
			}
		case _, ok = <-pingTicker.C:
			if ok == false {
				return
			}
			if c.closed == true {
				return
			}
			if err = c.conn.SetWriteDeadline(time.Now().Add(WS_WRITE_WAIT)); err != nil {
				wsLog.Warn(err.Error())
				return
			}
			if err = c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				wsLog.Warn(err.Error())
				return
			}
		case <-c.ctx.Done():
			return
		}
	}
}

func (c *Client) Send(message []byte) {
	defer func() {
		if r := recover(); r != nil {
			//wsLog.Warn(r, string(debug.Stack()))
		}
	}()
	c.rwLock.RLock()
	if c.closed == true {
		c.rwLock.RUnlock()
		return
	}
	c.rwLock.RUnlock()
	if len(c.sendChan) >= WS_WRITE_MESSAGE_THRESHOLD {
		return
	}
	c.sendChan <- message
}

func (c *Client) Close() (err error) {
	defer func() {
		if r := recover(); r != nil {
			//wsLog.Warn(r, string(debug.Stack()))
		}
	}()
	c.rwLock.RLock()
	if c.closed == true {
		c.rwLock.RUnlock()
		return
	}
	c.rwLock.RUnlock()
	err = c.conn.WriteMessage(websocket.CloseMessage, WS_MSG_BUF_CLOSE)
	if err != nil {
		wsLog.Warn(err.Error())
	}
	c.cancel()
	return
}
