package protocol

type ChatMessage struct {
	SrvMsgId        int64  `json:"srv_msg_id"`                                // 服务端消息号
	CliMsgId        int64  `json:"cli_msg_id" validate:"required,gt=0"`       // 客户端消息号
	SenderId        int64  `json:"sender_id" validate:"required,gt=0"`        // 发送者uid
	SenderPlatform  int    `json:"sender_platform" validate:"required,gte=0"` // 发送者平台
	SenderName      string `json:"sender_name"`                               // 发送者姓名
	SenderAvatarKey string `json:"sender_avatar_key"`                         // 发送者头像
	ChatId          int64  `json:"chat_id" validate:"required,gt=0"`          // 会话ID
	ChatType        int    `json:"chat_type"`                                 // 会话类型
	SeqId           int    `json:"seq_id"`                                    // 消息唯一ID
	MsgFrom         int    `json:"msg_from"`                                  // 消息来源
	MsgType         int    `json:"msg_type" validate:"required,gt=0"`         // 消息类型
	Body            string `json:"body"`                                      // 消息本体
	Status          int    `json:"status"`                                    // 消息状态
	SentTs          int64  `json:"sent_ts" validate:"required,gt=0"`          // 客户端本地发送时间
	SrvTs           int64  `json:"srv_ts"`                                    // 服务端接收消息的时间
}

// 图片 image
type Image struct {
	ImageKey string `json:"image_key" validate:"required,min=32,max=50"`
}

// 文件 file
type File struct {
	FileKey  string `json:"file_key" validate:"required,min=32,max=50"`
	FileName string `json:"file_name" validate:"required"` // 文件名
}

// 音频 audio
type Audio struct {
	FileKey  string `json:"file_key" validate:"required,min=32,max=50"` // 文件key
	Duration int    `json:"duration" validate:"gt=500"`                 // 时长 毫秒级
}

// 视频 media
type Media struct {
	FileKey  string `json:"file_key" validate:"required,min=32,max=50"`  // 文件key
	ImageKey string `json:"image_key" validate:"required,min=32,max=50"` // 视频封面图片key
	FileName string `json:"file_name" validate:"required"`               // 文件名
	Duration int    `json:"duration" validate:"gt=500"`                  // 视频时长 毫秒级
}

// 表情包 sticker
type Sticker struct {
	FileKey string `json:"file_key" validate:"required,min=32,max=50"` // 文件key
}

type Result struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
}

func (r *Result) Set(code int32, msg string) {
	r.Code = code
	r.Msg = msg
}
