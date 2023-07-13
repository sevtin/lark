package ws

import (
	"bytes"
	"github.com/gorilla/websocket"
	"io"
	"lark/pkg/utils"
	"log"
	"runtime/debug"
	"sync"
	"time"
)

type Client struct {
	lock      sync.RWMutex
	hub       *Hub
	conn      *websocket.Conn
	uid       int64 // 用户ID
	platform  int32 // 平台ID
	key       string
	onlineTs  int64 // 上线时间戳（毫秒）
	sendChan  chan []byte
	closeChan chan struct{}
	closed    bool
}

func newClient(hub *Hub, conn *websocket.Conn, uid int64, platform int32) *Client {
	cli := &Client{
		hub:       hub,
		conn:      conn,
		uid:       uid,
		platform:  platform,
		key:       clientKey(uid, platform),
		onlineTs:  time.Now().UnixNano() / 1e6,
		sendChan:  make(chan []byte, WS_WRITE_MAX_MESSAGE_CHAN_SIZE),
		closeChan: make(chan struct{}),
	}
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
	c.lock.Lock()
	if c.closed == true {
		c.lock.Unlock()
		return
	}
	c.closed = true
	c.lock.Unlock()
	close(c.closeChan)
	c.hub.unregisterChan <- c
	nowAt := time.Now()
	c.conn.SetWriteDeadline(nowAt.Add(WS_RW_DEAD_LINE))
	c.conn.SetReadDeadline(nowAt.Add(WS_RW_DEAD_LINE))
	// 耗时操作!!!
	c.conn.Close()
}

func (c *Client) readLoop() {
	defer func() {
		if r := recover(); r != nil {
			wsLog.Warn(r, string(debug.Stack()))
		}
		c.closeConn()
	}()

	var (
		msgType int
		buf     []byte
		buffer  *bytes.Buffer
		message *Message
		err     error
	)

	c.conn.SetReadLimit(WS_READ_MAX_MESSAGE_BUFFER_SIZE)
	c.conn.SetReadDeadline(c.hub.now.Add(WS_PONG_WAIT))
	c.conn.SetPongHandler(c.pongHandler)
	//c.conn.SetPingHandler(c.pingHandler)
	c.conn.SetCloseHandler(c.closeHandler)

	for {
		if msgType, buf, err = c.conn.ReadMessage(); err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
			}
			/*
				switch err.(type) {
				case *websocket.CloseError:
				default:
					wsLog.Warn(err.Error())
				}
			*/
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
	err = c.conn.SetReadDeadline(c.hub.now.Add(WS_PONG_WAIT))
	if err != nil {
		wsLog.Warn(err.Error())
		c.closeConn()
	}
	return
}

func (c *Client) pingHandler(appData string) (err error) {
	err = c.conn.WriteControl(websocket.PongMessage, WS_MSG_BUF_PONG, c.hub.now.Add(time.Second))
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
	//c.conn.WriteControl(websocket.CloseMessage, message, c.hub.now.Add(time.Second))
	return
}

func (c *Client) writeLoop() {
	pingTicker := time.NewTicker(WS_PING_PERIOD)
	defer func() {
		if r := recover(); r != nil {
			wsLog.Warn(r, string(debug.Stack()))
		}
		pingTicker.Stop()
		c.closeConn()
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
			if err = c.conn.SetWriteDeadline(c.hub.now.Add(WS_WRITE_WAIT)); err != nil {
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
					if bufLen >= WS_WRITE_MAX_MERGE_MESSAGE_BUFFER_SIZE {
						break
					}
					merges++
					if merges > WS_WRITE_MAX_MERGE_MESSAGE_SIZE {
						break
					}
					message, ok = <-c.sendChan
					if ok == false {
						break
					}
					bufLen += len(message)
					if c.closed == true {
						break
					}
					_, err = wc.Write(message)
					if err != nil {
						//wsLog.Warn(err.Error())
						break
					}
				}
			}
			if err = wc.Close(); err != nil {
				//wsLog.Warn(err.Error())
				return
			}
		case _, ok = <-pingTicker.C:
			if ok == false {
				return
			}
			if c.closed == true {
				return
			}
			if err = c.conn.SetWriteDeadline(c.hub.now.Add(WS_WRITE_WAIT)); err != nil {
				wsLog.Warn(err.Error())
				return
			}
			if err = c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				wsLog.Warn(err.Error())
				return
			}
		case <-c.closeChan:
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
	c.lock.RLock()
	if c.closed == true {
		c.lock.RUnlock()
		return
	}
	c.lock.RUnlock()
	if len(c.sendChan) >= WS_WRITE_MESSAGE_THRESHOLD {
		return
	}
	c.sendChan <- message
}

func (c *Client) Close() (err error) {
	c.lock.Lock()
	if c.closed == true {
		c.lock.Unlock()
		return
	}
	c.lock.Unlock()
	err = c.conn.WriteMessage(websocket.CloseMessage, WS_MSG_BUF_CLOSE)
	if err != nil {
		wsLog.Warn(err.Error())
	}
	c.closeConn()
	return
}
