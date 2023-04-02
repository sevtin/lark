package xhttp

const (
	ERROR_CODE_HTTP_REQ_DESERIALIZE_FAILED             int32 = 601 // 请求参数序列化错误
	ERROR_CODE_HTTP_REQ_PARAM_ERR                      int32 = 602 // 请求参数错误
	ERROR_CODE_HTTP_REQ_NOT_AUTHORIZED                 int32 = 603 // 没有授权
	ERROR_CODE_HTTP_REGISTER_FAILED                    int32 = 604 // 注册失败
	ERROR_CODE_HTTP_TOKEN_FAILED                       int32 = 605 // 获取TOKEN失败
	ERROR_CODE_HTTP_JWT_TOKEN_ERR                      int32 = 606 // TOKEN 错误
	ERROR_CODE_HTTP_USER_ID_DOESNOT_EXIST              int32 = 607 // 用户ID信息缺失
	ERROR_CODE_HTTP_JWT_TOKEN_UUID_DOESNOT_EXIST       int32 = 608 // Token 信息缺失
	ERROR_CODE_HTTP_TOKEN_AUTHENTICATION_FAILED        int32 = 609 // Token 失效
	ERROR_CODE_HTTP_PLATFORM_DOESNOT_EXIST             int32 = 610 // 平台信息缺失
	ERROR_CODE_HTTP_GET_USER_INFO_FAILED               int32 = 611 // 获取用户信息失败
	ERROR_CODE_HTTP_ADD_FRIEND_FAILED                  int32 = 612 // 添加好友失败
	ERROR_CODE_HTTP_SERVICE_FAILURE                    int32 = 613 // 服务故障
	ERROR_CODE_HTTP_MESSAGE_ENQUEUE_FAILED             int32 = 614 // 消息入队失败
	ERROR_CODE_HTTP_PRESIGNED_FAILED                   int32 = 615 // 上传文件预先签署失败
	ERROR_CODE_HTTP_READ_UPLOAD_FILE_FAILED            int32 = 616 // 读取上传文件失败
	ERROR_CODE_HTTP_OPEN_UPLOAD_FILE_FAILED            int32 = 617 // 打开上传文件失败
	ERROR_CODE_HTTP_CROP_PHOTO_FAILED                  int32 = 618 // 裁剪图片失败
	ERROR_CODE_HTTP_PAGINATION_LIMIT_EXCEEDED          int32 = 619 // 超出分页限制
	ERROR_CODE_HTTP_JWT_TOKEN_SESSION_ID_DOESNOT_EXIST int32 = 620 // SESSION ID 缺失
)

const (
	ERROR_CODE_HTTP_400 = 400
)
