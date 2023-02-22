package constant

import "time"

const (
	CONST_DURATION_BASIC_USER_INFO_SECOND = 60 * 60 * 24 * 7 * time.Second //用户基础信息缓存时间
	CONST_DURATION_USER_INFO_SECOND       = 60 * 60 * 24 * 7 * time.Second //用户信息缓存时间
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
	CONST_DURATION_USER_SERVER_ID_SECOND = 60 * 60 * 24 * 7 * time.Second //用户 server id 缓存时间
)

const (
	CONST_DURATION_DIST_CHAT_MEMBER_HASH_SECOND = 60 * 60 * 24 * 7 * time.Second //消息分发缓存hash
	CONST_DURATION_CHAT_MEMBER_INFO_HASH_SECOND = 60 * 60 * 24 * 7 * time.Second //chat成员缓存hash
)

const (
	CONST_DURATION_JWT_ACCESS_TOKEN_EXPIRE_IN_SECOND = 60 * 60 * 24 * 7 * time.Second //ACCESS_TOKEN有效期
	//CONST_DURATION_JWT_REFRESH_TOKEN_EXPIRE_IN_SECOND = 3600 * 24 * 30 * time.Second //REFRESH_TOKEN有效期
)

const (
	CONST_DURATION_SERVER_MGR_SECOND    = 60 * 60 * 10 * time.Second
	CONST_DURATION_REDSYNC_MUTEX_SECOND = 5 * time.Second
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
