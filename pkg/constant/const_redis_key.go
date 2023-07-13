package constant

// 无状态
const (
	RK_MSG_CLI_MSG_ID   = "MSG:CLI_MSG_ID:" // 缓存客户端消息ID + ChatId + CliMsgId {}
	RK_MSG_CONVO_MSG    = "MSG:CONVO:"      // 最新一条会话消息 {}
	RK_MSG_SEQ_TS       = "MSG:SEQ_TS:"     // 最新一条会话消息 SEQ + TIMESTAMP {}
	RK_MSG_SEQ_ID       = "MSG:SEQ_ID:"     // {}
	RK_USER_LOCK_MOBILE = "USER_LOCK:MOBILE:"
)

// 有状态
const (
	RK_SYNC_USER_ACCESS_TOKEN_SESSION_ID  = "USER:ACCESS_TOKEN_SESSION_ID:"  // {}
	RK_SYNC_USER_REFRESH_TOKEN_SESSION_ID = "USER:REFRESH_TOKEN_SESSION_ID:" // {}
	RK_SYNC_USER_ACCESS_TOKEN             = "USER:ACCESS_TOKEN:"             // {}
	RK_SYNC_USER_INFO                     = "USER:INFO:"                     // 用户信息 {}
	RK_SYNC_BASIC_USER_INFO               = "USER:BASIC_INFO:"               // 基础用户信息 {}
	RK_SYNC_USER_SERVER                   = "SRV:"                           // WS服务器ID {}
	RK_SYNC_MSG_CACHE                     = "MSG:CACHE:"                     // 消息缓存 + ChatId + seqId {}
	RK_SYNC_DIST_CHAT_MEMBER_HASH         = "CHAT:DIST_MEMBER_HASH:"         // 消息派发Chat成员列表 {}
	RK_SYNC_CHAT_MEMBER_INFO_HASH         = "CHAT:MEMBER_INFO_HASH:"         // Chat成员信息列表 {}
	RK_SYNC_GROUP_CHAT_INFO               = "CHAT:GROUP_CHAT_INFO:"          // {}
	RK_SYNC_SERVER_MSG_GATEWAY            = "SERVER:MSG_GATEWAY"
	RK_SYNC_CONVO_LIST                    = "CONVO:LIST:"
	RK_SYNC_SERVER_MGR                    = "SERVER_MGR:"
	RK_SYNC_SERVER_REDSYNC_MUTEX          = "SERVER_REDSYNC_MUTEX:"
)

// 弃用key
const (
	RK_SYNC_GROUP_CHAT_DETAILS = "CHAT:GROUP_CHAT_DETAILS:"
)
