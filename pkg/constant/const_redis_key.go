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
	RK_SYNC_RED_REMAIN_QUANTITY           = "RED_ENV:REMAIN_QUANTITY:" // 剩余红包数量
	RK_SYNC_RED_REMAIN_AMOUNT             = "RED_ENV:REMAIN_AMOUNT:"   // 剩余红包金额
	RK_SYNC_RED_ENV_STATUS                = "RED_ENV:STATUS:"          // 红包状态
	RK_SYNC_RED_ENV_RECORD                = "RED_ENV:RECORD:"          // 红包领取记录
	RK_SYNC_RED_ENV_INFO                  = "RED_ENV:INFO:"            // 红包信息
	RK_SYNC_RED_ENV_LOCK                  = "RED_ENV:LOCK:"            // 红包分布式锁
	RK_SYNC_WALLET_ACCOUNT_INFO           = "WALLET:ACCOUNT_INFO:"     // 钱包账户信息
	RK_SYNC_RED_ENV_KEY                   = "RED_ENV:KEY:"             // 红包Key
)

// 弃用key
const (
	RK_SYNC_GROUP_CHAT_DETAILS = "CHAT:GROUP_CHAT_DETAILS:"
)

const (
	RK_FIELD_RED_ENV_TOTAL_NUM    = "total_num"    // 红包总数
	RK_FIELD_RED_ENV_RECEIVED_NUM = "received_num" // 已领红包数
)
