package client

import (
	"time"
)

const (
	// Time allowed to write a message to the peer.
	WS_WRITE_WAIT = 10 * time.Second
	// Time allowed to read the next pong message from the peer.
	WS_PONG_WAIT = 60 * time.Second
	// Send pings to peer with this period. Must be less than pongWait.
	WS_PING_PERIOD = (WS_PONG_WAIT * 9) / 10
	WS_RITE_WAIT   = time.Second
	// Maximum message size allowed from peer.
	WS_MAX_MESSAGE_SIZE                    = 4096
	WS_READ_BUFFER_SIZE                    = 4096
	WS_WRITE_BUFFER_SIZE                   = 2048
	WS_HEADER_LENGTH                       = 4
	WS_CHAN_CLIENT_WRITE_MESSAGE_SIZE      = 1024
	WS_CHAN_SERVER_WRITE_MESSAGE_TOLERANCE = 24
	WS_CHAN_CLIENT_WRITE_MESSAGE_THRESHOLD = WS_CHAN_CLIENT_WRITE_MESSAGE_SIZE - WS_CHAN_SERVER_WRITE_MESSAGE_TOLERANCE
	WS_CHAN_CLOSE_SIZE                     = 2
	WS_CHAN_CLOSE_THRESHOLD                = WS_CHAN_CLOSE_SIZE - 1
	WS_CHAN_SERVER_READ_MESSAGE_SIZE       = 10000
	WS_CHAN_CLIENT_REGISTER_SIZE           = 10000
	WS_CHAN_CLIENT_UNREGISTER_SIZE         = 10000
	WS_MAX_CONNECTIONS                     = 20000
	// TODO: 测试值
	WS_MINIMUM_TIME_INTERVAL = -1 //5000
	// 最小消息间隔时长(毫秒)
	WS_MINIMUM_INTERVAL = 250
)

const (
	WS_PING  = 61001
	WS_PONG  = 61002
	WS_CLOSE = 61003
)

const (
	WS_SEND_MSG_SUCCESS = 0
	WS_CLIENT_OFFLINE   = 61004
	WS_SEND_MSG_FAILED  = 61005
)

var (
	WS_MSG_BUF_NEWLINE = []byte{'\n'}
	WS_MSG_BUF_SPACE   = []byte{' '}
	WS_MSG_BUF_PING    = []byte{}
	WS_MSG_BUF_PONG    = []byte{}
	WS_MSG_BUF_CLOSE   = []byte{}
)

const (
	WS_KEY_UID      = "uid"
	WS_KEY_PLATFORM = "platform"
	WS_KEY_COOKIE   = "Cookie"
)

const (
	ERROR_CODE_HTTP_UPGRADER_FAILED        = 60001
	ERROR_CODE_HTTP_UID_DOESNOT_EXIST      = 60002
	ERROR_CODE_HTTP_PLATFORM_DOESNOT_EXIST = 60003
	ERROR_CODE_HTTP_REQUEST_TOO_MUNDANE    = 60004
)

const (
	ERROR_HTTP_UID_DOESNOT_EXIST      = "uid 缺失"
	ERROR_HTTP_PLATFORM_DOESNOT_EXIST = "platform 缺失"
	ERROR_HTTP_REQUEST_TOO_MUNDANE    = "请求过于平凡"
)

const (
	ERROR_CODE_WS_EXCEED_MAX_CONNECTIONS = 60005
)

const (
	ERROR_WS_EXCEED_MAX_CONNECTIONS = "超出最大连接数限制"
)

const (
	MessageLength   uint32 = 4
	MessageTopic    uint32 = 8
	MessageSubtopic uint32 = 12
	MessageType     uint32 = 16
)
