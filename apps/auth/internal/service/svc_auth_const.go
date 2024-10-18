package service

import "errors"

const (
	ERROR_CODE_AUTH_ACCOUNT_TYPE_ERR                             int32 = 200001
	ERROR_CODE_AUTH_ACCOUNT_OR_PASSWORD_ERR                      int32 = 200002
	ERROR_CODE_AUTH_QUERY_DB_FAILED                              int32 = 200003
	ERROR_CODE_AUTH_REDIS_GET_FAILED                             int32 = 200004
	ERROR_CODE_AUTH_REGISTER_ERR                                 int32 = 200005
	ERROR_CODE_AUTH_INSERT_VALUE_FAILED                          int32 = 200006
	ERROR_CODE_AUTH_ACCOUNT_DOES_NOT_EXIST                       int32 = 200007
	ERROR_CODE_AUTH_MOBILE_HAS_BEEN_REGISTERED                   int32 = 200008
	ERROR_CODE_AUTH_LOGOUT_FAILED                                int32 = 200009
	ERROR_CODE_AUTH_UPDATE_VALUE_FAILED                          int32 = 200010
	ERROR_CODE_AUTH_GENERATE_TOKEN_FAILED                        int32 = 200011
	ERROR_CODE_AUTH_REDIS_SET_FAILED                             int32 = 200012
	ERROR_CODE_AUTH_JWT_TOKEN_ERR                                int32 = 200013
	ERROR_CODE_AUTH_JWT_SESSION_ID_ERR                           int32 = 200014
	ERROR_CODE_AUTH_THE_MOBILE_HAS_BEEN_BOUND_TO_ANOTHER_ACCOUNT int32 = 200015
	ERROR_CODE_AUTH_OAUTH_TOKEN_ACQUISITION_FAILED               int32 = 200016
	ERROR_CODE_AUTH_OAUTH_USER_INFO_ACQUISITION_FAILED           int32 = 200017
	ERROR_CODE_AUTH_OAUTH_USER_INFO_QUERY_FAILED                 int32 = 200018
	ERROR_CODE_AUTH_GRPC_SERVICE_FAILURE                         int32 = 200019
	ERROR_CODE_AUTH_UPDATE_USER_SERVER_ID_FAILED                 int32 = 200020
)

const (
	ERROR_AUTH_ACCOUNT_TYPE_ERR                             = "账户类型错误"
	ERROR_AUTH_ACCOUNT_OR_PASSWORD_ERR                      = "账户或密码错误"
	ERROR_AUTH_QUERY_DB_FAILED                              = "查询失败"
	ERROR_AUTH_REDIS_GET_FAILED                             = "读取redis缓存失败"
	ERROR_AUTH_REGISTER_TYPE_ERR                            = "注册失败"
	ERROR_AUTH_INSERT_VALUE_FAILED                          = "数据入库失败"
	ERROR_AUTH_ACCOUNT_DOES_NOT_EXIST                       = "账号不存在"
	ERROR_AUTH_MOBILE_HAS_BEEN_REGISTERED                   = "该手机已注册"
	ERROR_AUTH_LOGOUT_FAILED                                = "退出账号失败"
	ERROR_AUTH_UPDATE_VALUE_FAILED                          = "更新Value失败"
	ERROR_AUTH_GENERATE_TOKEN_FAILED                        = "生成token失败"
	ERROR_AUTH_REDIS_SET_FAILED                             = "设置redis缓存失败"
	ERROR_AUTH_JWT_TOKEN_ERR                                = "授权token错误"
	ERROR_AUTH_JWT_SESSION_ID_ERR                           = "会话ID错误"
	ERROR_AUTH_THE_MOBILE_HAS_BEEN_BOUND_TO_ANOTHER_ACCOUNT = "该手机号已绑定其他账号"
	ERROR_AUTH_OAUTH_TOKEN_ACQUISITION_FAILED               = "获取token失败"
	ERROR_AUTH_OAUTH_USER_INFO_ACQUISITION_FAILED           = "获取用户信息失败"
	ERROR_AUTH_OAUTH_USER_INFO_QUERY_FAILED                 = "查询用户信息失败"
	ERROR_AUTH_GRPC_SERVICE_FAILURE                         = "服务故障"
	ERROR_AUTH_UPDATE_USER_SERVER_ID_FAILED                 = "更新ServerId失败"
)

var (
	ERR_AUTH_THE_MOBILE_HAS_BEEN_BOUND_TO_ANOTHER_ACCOUNT = errors.New("该手机号已绑定其他账号")
	ERR_AUTH_OAUTH_TOKEN_ACQUISITION_FAILED               = errors.New("oauth token获取失败")
	ERR_AUTH_OAUTH_USER_INFO_ACQUISITION_FAILED           = errors.New("获取用户信息失败")
)

const (
	API_GITHUB_OAUTH_ACCESS_TOKEN = "https://github.com/login/oauth/access_token?client_id=%s&client_secret=%s&code=%s"
	API_GITHUB_USER               = "https://api.github.com/user"
)
const (
	API_GOOGLE_USERINFO = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="
)

const (
	DEFAULT_LOGIN_PASSWORD = "EA405B607DE5E4F6797640AB81F1767D" // 密码 12345678
)
