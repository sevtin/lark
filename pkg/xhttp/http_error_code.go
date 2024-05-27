package xhttp

const (
	ERROR_CODE_HTTP_JWT_TOKEN_ERR     int32 = 401 // TOKEN 错误
	ERROR_CODE_HTTP_TOO_MANY_REQUESTS       = 429
)

const (
	ERROR_CODE_HTTP_USER_ID_DOESNOT_EXIST    int32 = 601 // 用户ID信息缺失
	ERROR_CODE_HTTP_PLATFORM_DOESNOT_EXIST   int32 = 602 // 平台信息缺失
	ERROR_CODE_HTTP_SESSION_ID_DOESNOT_EXIST int32 = 603 // SESSION ID 缺失
	ERROR_CODE_HTTP_READ_SESSION_ID_FAILED   int32 = 604 // 读取会话ID失败
	ERROR_CODE_HTTP_TOKEN_EXPIRES            int32 = 605 // Token 过期
	ERROR_CODE_HTTP_ACCOUNT_KICKED_OUT       int32 = 606 // 账号被踢出
	ERROR_CODE_HTTP_GITHUB_TOKEN_EXPIRES     int32 = 607 // Token 过期
)

const (
	ERROR_CODE_HTTP_REQ_DESERIALIZE_FAILED           int32 = 701 // 请求参数序列化错误
	ERROR_CODE_HTTP_REQ_PARAM_ERR                    int32 = 702 // 请求参数错误
	ERROR_CODE_HTTP_REQ_NOT_AUTHORIZED               int32 = 703 // 没有授权
	ERROR_CODE_HTTP_REGISTER_FAILED                  int32 = 704 // 注册失败
	ERROR_CODE_HTTP_TOKEN_FAILED                     int32 = 705 // 获取TOKEN失败
	ERROR_CODE_HTTP_JWT_TOKEN_UUID_DOESNOT_EXIST     int32 = 706 // Token 信息缺失
	ERROR_CODE_HTTP_GET_USER_INFO_FAILED             int32 = 707 // 获取用户信息失败
	ERROR_CODE_HTTP_ADD_FRIEND_FAILED                int32 = 708 // 添加好友失败
	ERROR_CODE_HTTP_SERVICE_FAILURE                  int32 = 709 // 服务故障
	ERROR_CODE_HTTP_MESSAGE_ENQUEUE_FAILED           int32 = 710 // 消息入队失败
	ERROR_CODE_HTTP_PRESIGNED_FAILED                 int32 = 711 // 上传文件预先签署失败
	ERROR_CODE_HTTP_READ_UPLOAD_FILE_FAILED          int32 = 712 // 读取上传文件失败
	ERROR_CODE_HTTP_OPEN_UPLOAD_FILE_FAILED          int32 = 713 // 打开上传文件失败
	ERROR_CODE_HTTP_CROP_PHOTO_FAILED                int32 = 714 // 裁剪图片失败
	ERROR_CODE_HTTP_PAGINATION_LIMIT_EXCEEDED        int32 = 715 // 超出分页限制
	ERROR_CODE_HTTP_ALIPAY_NOTIFY_VERIFY_SIGN_FAILED int32 = 716 // alipay通知验签失败
)
const (
	ERROR_CODE_HTTP_400          = 400
	ERROR_CODE_HTTP_UNAUTHORIZED = 401
)

// HTTP status codes as registered with IANA.
// See: https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml
const (
	StatusContinue           = 100 // RFC 9110, 15.2.1
	StatusSwitchingProtocols = 101 // RFC 9110, 15.2.2
	StatusProcessing         = 102 // RFC 2518, 10.1
	StatusEarlyHints         = 103 // RFC 8297

	StatusOK                   = 200 // RFC 9110, 15.3.1
	StatusCreated              = 201 // RFC 9110, 15.3.2
	StatusAccepted             = 202 // RFC 9110, 15.3.3
	StatusNonAuthoritativeInfo = 203 // RFC 9110, 15.3.4
	StatusNoContent            = 204 // RFC 9110, 15.3.5
	StatusResetContent         = 205 // RFC 9110, 15.3.6
	StatusPartialContent       = 206 // RFC 9110, 15.3.7
	StatusMultiStatus          = 207 // RFC 4918, 11.1
	StatusAlreadyReported      = 208 // RFC 5842, 7.1
	StatusIMUsed               = 226 // RFC 3229, 10.4.1

	StatusMultipleChoices   = 300 // RFC 9110, 15.4.1
	StatusMovedPermanently  = 301 // RFC 9110, 15.4.2
	StatusFound             = 302 // RFC 9110, 15.4.3
	StatusSeeOther          = 303 // RFC 9110, 15.4.4
	StatusNotModified       = 304 // RFC 9110, 15.4.5
	StatusUseProxy          = 305 // RFC 9110, 15.4.6
	_                       = 306 // RFC 9110, 15.4.7 (Unused)
	StatusTemporaryRedirect = 307 // RFC 9110, 15.4.8
	StatusPermanentRedirect = 308 // RFC 9110, 15.4.9

	StatusBadRequest                   = 400 // RFC 9110, 15.5.1
	StatusUnauthorized                 = 401 // RFC 9110, 15.5.2
	StatusPaymentRequired              = 402 // RFC 9110, 15.5.3
	StatusForbidden                    = 403 // RFC 9110, 15.5.4
	StatusNotFound                     = 404 // RFC 9110, 15.5.5
	StatusMethodNotAllowed             = 405 // RFC 9110, 15.5.6
	StatusNotAcceptable                = 406 // RFC 9110, 15.5.7
	StatusProxyAuthRequired            = 407 // RFC 9110, 15.5.8
	StatusRequestTimeout               = 408 // RFC 9110, 15.5.9
	StatusConflict                     = 409 // RFC 9110, 15.5.10
	StatusGone                         = 410 // RFC 9110, 15.5.11
	StatusLengthRequired               = 411 // RFC 9110, 15.5.12
	StatusPreconditionFailed           = 412 // RFC 9110, 15.5.13
	StatusRequestEntityTooLarge        = 413 // RFC 9110, 15.5.14
	StatusRequestURITooLong            = 414 // RFC 9110, 15.5.15
	StatusUnsupportedMediaType         = 415 // RFC 9110, 15.5.16
	StatusRequestedRangeNotSatisfiable = 416 // RFC 9110, 15.5.17
	StatusExpectationFailed            = 417 // RFC 9110, 15.5.18
	StatusTeapot                       = 418 // RFC 9110, 15.5.19 (Unused)
	StatusMisdirectedRequest           = 421 // RFC 9110, 15.5.20
	StatusUnprocessableEntity          = 422 // RFC 9110, 15.5.21
	StatusLocked                       = 423 // RFC 4918, 11.3
	StatusFailedDependency             = 424 // RFC 4918, 11.4
	StatusTooEarly                     = 425 // RFC 8470, 5.2.
	StatusUpgradeRequired              = 426 // RFC 9110, 15.5.22
	StatusPreconditionRequired         = 428 // RFC 6585, 3
	StatusTooManyRequests              = 429 // RFC 6585, 4
	StatusRequestHeaderFieldsTooLarge  = 431 // RFC 6585, 5
	StatusUnavailableForLegalReasons   = 451 // RFC 7725, 3

	StatusInternalServerError           = 500 // RFC 9110, 15.6.1
	StatusNotImplemented                = 501 // RFC 9110, 15.6.2
	StatusBadGateway                    = 502 // RFC 9110, 15.6.3
	StatusServiceUnavailable            = 503 // RFC 9110, 15.6.4
	StatusGatewayTimeout                = 504 // RFC 9110, 15.6.5
	StatusHTTPVersionNotSupported       = 505 // RFC 9110, 15.6.6
	StatusVariantAlsoNegotiates         = 506 // RFC 2295, 8.1
	StatusInsufficientStorage           = 507 // RFC 4918, 11.5
	StatusLoopDetected                  = 508 // RFC 5842, 7.2
	StatusNotExtended                   = 510 // RFC 2774, 7
	StatusNetworkAuthenticationRequired = 511 // RFC 6585, 6
)
