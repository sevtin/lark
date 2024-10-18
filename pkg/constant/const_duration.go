package constant

import "time"

const (
	CONST_DURATION_BASIC_USER_INFO_SECOND  = 60 * 60 * 24 * 7 * time.Second //用户基础信息缓存时间
	CONST_DURATION_USER_INFO_SECOND        = 60 * 60 * 24 * 7 * time.Second //用户信息缓存时间
	CONST_DURATION_OAUTH_USER_TOKEN_SECOND = 60 * 60 * 24 * 7 * time.Second

	CONST_DURATION_USER_OAUTH_USER_INFO_SECOND = 60 * 60 * 24 * 7 * time.Second //OAUTH_USER基础信息缓存时
)

const (
	CONST_DURATION_GROUP_CHAT_INFO_SECOND = 60 * 60 * 24 * 7 * time.Second //群信息缓存时间
)

const (
	CONST_DURATION_MSG_CACHE_SECOND = 60 * 60 * 24 * time.Second //单个消息缓存时间
)

const (
	CONST_DURATION_NUMBER_OF_CHAT_MEMBER_ONLINE_SECOND = 5 * time.Second //更新在线人数
)

const (
	CONST_DURATION_USER_SERVER_ID_SECOND = 60 * 60 * 24 * 60 * time.Second //用户 server id 缓存时间
)

const (
	CONST_DURATION_DIST_CHAT_MEMBER_HASH_SECOND = 60 * 60 * 24 * 120 * time.Second //消息分发缓存hash
	CONST_DURATION_CHAT_MEMBER_INFO_HASH_SECOND = 60 * 60 * 24 * 7 * time.Second   //chat成员缓存hash
)

const (
	CONST_DURATION_JWT_ACCESS_TOKEN_EXPIRE_IN_SECOND  = 60 * 60 * 24 * 7 * time.Second  //ACCESS_TOKEN有效期
	CONST_DURATION_JWT_REFRESH_TOKEN_EXPIRE_IN_SECOND = 60 * 60 * 24 * 30 * time.Second //REFRESH_TOKEN有效期
)

const (
	CONST_DURATION_RED_ENVELOPE_EXPIRE_SECOND     = 60 * 60 * 48 * time.Second //红包缓存时间
	CONST_DURATION_RED_ENVELOPE_KEY_EXPIRE_SECOND = 60 * 10 * time.Second      //红包Key缓存时间
)

const (
	CONST_DURATION_24H_SECOND = 60 * 60 * 24
)

const (
	CONST_DURATION_SERVER_MGR_SECOND    = 60 * 60 * 10 * time.Second
	CONST_DURATION_REDSYNC_MUTEX_SECOND = 5 * time.Second
	CONST_DURATION_REDIS_LOCK_EXPIRY    = 10 * time.Second
)

const (
	CONST_DURATION_WALLET_ACCOUNT_INFO_SECOND = 60 * 60 * 24 * time.Second
	CONST_DURATION_USER_WALLETS_SECOND        = 60 * 60 * 24 * time.Second
)

const (
	CONST_DURATION_SHA_CONVO_MESSAGE_SECOND               = 60 * 60 * 24      //会话消息缓存时间
	CONST_DURATION_SHA_BASIC_USER_INFO_SECOND             = 60 * 60 * 24 * 7  //用户基础信息缓存时间
	CONST_DURATION_SHA_USER_INFO_SECOND                   = 60 * 60 * 24 * 7  //用户信息缓存时间
	CONST_DURATION_SHA_MSG_ID_SECOND                      = 60 * 60 * 12      //消息ID缓存时间,用于判断是否是重复消息
	CONST_DURATION_SHA_USER_SERVER_ID_SECOND              = 60 * 60 * 24 * 7  //用户 server id 缓存时间
	CONST_DURATION_SHA_JWT_ACCESS_TOKEN_EXPIRE_IN_SECOND  = 60 * 60 * 24 * 7  //ACCESS_TOKEN有效期
	CONST_DURATION_SHA_JWT_REFRESH_TOKEN_EXPIRE_IN_SECOND = 60 * 60 * 24 * 30 //REFRESH_TOKEN有效期
	CONST_DURATION_SHA_CHAT_MEMBER_INFO_HASH_SECOND       = 60 * 60 * 24 * 7  //chat成员缓存hash
)

const (
	CONST_DURATION_LBS_QUERY_LAST_ONLINE_SECOND = 60 * 60 * 24 * 14
)

const (
	CONST_DURATION_AWS_S3_EXPIRE_MINUTE = 10 * time.Minute
)

const (
	CONST_DURATION_RED_ENVELOPE_DISTRIBUTED_LOCK_EXPIRY_IN_SECOND = 15 * time.Second
)

const (
	CONST_DURATION_ORDER_STATUS_SECOND = 60 * 60 * time.Second
	CONST_DURATION_REPO_INFO_SECOND    = 60 * 5 * time.Second
	CONST_DURATION_TASK_REWARD_SECOND  = 60 * 60 * 24 * time.Second
	CONST_DURATION_TASK_STATUS_SECOND  = 60 * 60 * time.Second
)

const (
	CONST_DURATION_INVITATION_CODE_SECOND = 60 * 60 * 24 * 30 * time.Second
)

const (
	CONST_DURATION_WALLET_RESET_PWD_CODE_SECOND = 60 * 30 * time.Second
)
