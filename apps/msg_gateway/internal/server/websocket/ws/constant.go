package ws

import (
	"time"
)

const (
	// Time allowed to write a message to the peer.
	WS_WRITE_WAIT = 10 * time.Second
	// Time allowed to read the next pong message from the peer.
	WS_PONG_WAIT = 60 * time.Second
	// Send pings to peer with this period. Must be less than pongWait.
	WS_PING_PERIOD    = (WS_PONG_WAIT * 9) / 10
	WS_RITE_WAIT      = time.Second
	WS_WRITE_TIME_OUT = 5 * time.Second
	WS_RW_DEAD_LINE   = 200 * time.Millisecond
	// Maximum message size allowed from peer.
	WS_READ_MAX_MESSAGE_BUFFER_SIZE        = 2048
	WS_WRITE_MAX_MESSAGE_BUFFER_SIZE       = 2048
	WS_WRITE_MAX_MERGE_MESSAGE_BUFFER_SIZE = 1024
	WS_WRITE_MAX_MESSAGE_CHAN_SIZE         = 256
	WS_WRITE_MAX_MERGE_MESSAGE_SIZE        = 16
	WS_WRITE_MESSAGE_THRESHOLD             = WS_WRITE_MAX_MESSAGE_CHAN_SIZE - WS_WRITE_MAX_MERGE_MESSAGE_SIZE
	WS_READ_BUFFER_SIZE                    = 1024
	WS_WRITE_BUFFER_SIZE                   = 1024
	WS_CHAN_SERVER_READ_MESSAGE_SIZE       = 10000
	WS_CHAN_SERVER_READ_MESSAGE_TOLERANCE  = 256
	WS_CHAN_SERVER_READ_MESSAGE_THRESHOLD  = WS_CHAN_SERVER_READ_MESSAGE_SIZE - WS_CHAN_SERVER_READ_MESSAGE_TOLERANCE
	WS_CHAN_CLIENT_REGISTER_SIZE           = 10000
	WS_CHAN_CLIENT_UNREGISTER_SIZE         = 10000
	WS_MAX_CONNECTIONS                     = 1000000
	WS_MAX_CONSUMER_SIZE                   = 5120
	// TODO: 测试值
	WS_MINIMUM_TIME_INTERVAL = -1 //5000
	// 最小消息间隔时长(毫秒)
	WS_MINIMUM_INTERVAL = 200
)

const (
	WS_PING  = 3091
	WS_PONG  = 3092
	WS_CLOSE = 3093
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
)

const (
	ERROR_CODE_HTTP_UPGRADER_FAILED        = 3001
	ERROR_CODE_HTTP_UID_DOESNOT_EXIST      = 3002
	ERROR_CODE_HTTP_PLATFORM_DOESNOT_EXIST = 3003
	ERROR_CODE_HTTP_REQUEST_TOO_MUNDANE    = 3004
)
const (
	ERROR_HTTP_UID_DOESNOT_EXIST      = "uid 缺失"
	ERROR_HTTP_PLATFORM_DOESNOT_EXIST = "platform 缺失"
	ERROR_HTTP_REQUEST_TOO_MUNDANE    = "请求过于平凡"
)

const (
	WS_SEND_MSG_SUCCESS = 0
	WS_CLIENT_OFFLINE   = 3021
	WS_SEND_MSG_FAILED  = 3022
)

const (
	ERROR_CODE_WS_EXCEED_MAX_CONNECTIONS = 3031
)
const (
	ERROR_WS_EXCEED_MAX_CONNECTIONS = "超出最大连接数限制"
)

const (
	ERROR_CODE_WS_SERVER_OVERLOAD = 3061
)
const (
	ERROR_WS_SERVER_OVERLOAD = "服务器过载"
)
