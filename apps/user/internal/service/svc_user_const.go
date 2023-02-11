package service

var (
	userUpdateFields = map[string]string{
		"lark_id":   "lark_id",
		"nickname":  "nickname",
		"firstname": "firstname",
		"lastname":  "lastname",
		"gender":    "gender",
		"birth_ts":  "birth_ts",
		"email":     "email",
		"mobile":    "mobile",
		"city_id":   "city_id",
	}
)

const (
	ERROR_CODE_USER_ACCOUNT_TYPE_ERR              int32 = 1001
	ERROR_CODE_USER_ACCOUNT_OR_PASSWORD_ERR       int32 = 1002
	ERROR_CODE_USER_QUERY_DB_FAILED               int32 = 1003
	ERROR_CODE_USER_REDIS_GET_FAILED              int32 = 1004
	ERROR_CODE_USER_REDIS_SET_FAILED              int32 = 1005
	ERROR_CODE_USER_SET_AVATAR_FAILED             int32 = 1006
	ERROR_CODE_USER_UPDATE_VALUE_FAILED           int32 = 1007
	ERROR_CODE_USER_MOBILE_HAS_BEEN_OCCUPIED      int32 = 1008
	ERROR_CODE_USER_LARK_ID_HAS_BEEN_OCCUPIED     int32 = 1009
	ERROR_CODE_USER_MARSHAL_FAILED                int32 = 1010
	ERROR_CODE_USER_CACHE_CHAT_MEMBER_INFO_FAILED int32 = 1011
	ERROR_CODE_USER_UPDATE_USER_CACHE_FAILED      int32 = 1012
)

const (
	ERROR_USER_ACCOUNT_TYPE_ERR              = "登录类型错误"
	ERROR_USER_ACCOUNT_OR_PASSWORD_ERR       = "账户或密码错误"
	ERROR_USER_QUERY_DB_FAILED               = "查询失败"
	ERROR_USER_REDIS_GET_FAILED              = "读取redis缓存失败"
	ERROR_USER_REDIS_SET_FAILED              = "缓存数据失败"
	ERROR_USER_SET_AVATAR_FAILED             = "设置用户头像失败"
	ERROR_USER_UPDATE_VALUE_FAILED           = "更新Value失败"
	ERROR_USER_MOBILE_HAS_BEEN_OCCUPIED      = "该手机号已被占用"
	ERROR_USER_LARK_ID_HAS_BEEN_OCCUPIED     = "该 LARK ID 已被占用"
	ERROR_USER_MARSHAL_FAILED                = "序列化失败"
	ERROR_USER_CACHE_CHAT_MEMBER_INFO_FAILED = "缓存Chat Member信息失败"
	ERROR_USER_UPDATE_USER_CACHE_FAILED      = "更新用户缓存失败"
)

const (
	ERROR_CODE_USER_REGISTER_ERR = 21101
)

const (
	ERROR_USER_REGISTER_TYPE_ERR = "注册失败"
)
