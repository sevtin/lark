package po

import "lark/pkg/entity"

type RedEnvelope struct {
	entity.GormEntityTs
	EnvId          int64  `gorm:"column:env_id;primary_key" json:"env_id"`                          // 红包ID
	EnvType        int    `gorm:"column:env_type;default:0;NOT NULL" json:"env_type"`               // 红包类型 1-均分红包 2-碰运气红包
	WalletId       int64  `gorm:"column:wallet_id;default:0;NOT NULL" json:"wallet_id"`             // 红包支出钱包ID
	ReceiverType   int    `gorm:"column:receiver_type;default:0;NOT NULL" json:"receiver_type"`     // 接收者类型 1-私聊对方 2-群聊所有人 3-群聊指定人
	TradeNo        string `gorm:"column:trade_no;NOT NULL" json:"trade_no"`                         // 交易编号
	ChatId         int64  `gorm:"column:chat_id;default:0;NOT NULL" json:"chat_id"`                 // ChatID
	SenderUid      int64  `gorm:"column:sender_uid;default:0;NOT NULL" json:"sender_uid"`           // 发红包用户ID
	SenderPlatform int    `gorm:"column:sender_platform;default:0;NOT NULL" json:"sender_platform"` // 发送平台
	Total          int64  `gorm:"column:total;default:0;NOT NULL" json:"total"`                     // 红包总金额(分)
	Quantity       int32  `gorm:"column:quantity;default:0;NOT NULL" json:"quantity"`               // 红包数量
	RemainQuantity int32  `gorm:"column:remain_quantity;default:0;NOT NULL" json:"remain_quantity"` // 剩余红包数量
	RemainAmount   int64  `gorm:"column:remain_amount;default:0;NOT NULL" json:"remain_amount"`     // 剩余红包金额(分)
	Message        string `gorm:"column:message;default:恭喜发财;NOT NULL" json:"message"`              // 祝福语
	EnvStatus      int    `gorm:"column:env_status;default:0;NOT NULL" json:"env_status"`           // 状态 0-创建 1-已发放 2-已领完 3-已过期且退还剩余红包
	PayStatus      int    `gorm:"column:pay_status;default:0;NOT NULL" json:"pay_status"`           // 支付状态 0-未支付 1-支付中 2-已支付 3-支付失败
	ExpiredTs      int64  `gorm:"column:expired_ts;default:0;NOT NULL" json:"expired_ts"`           // 过期时间
	FinishedTs     int64  `gorm:"column:finished_ts;default:0;NOT NULL" json:"finished_ts"`         // 红包领完时间
	Receivers      string `gorm:"column:receivers;NOT NULL" json:"receivers"`                       // 接收人IDs 逗号分隔
}

type FundFlow struct {
	entity.GormEntityTs
	FlowId      int64  `gorm:"column:flow_id;primary_key" json:"flow_id"`                    // 流水ID
	Uid         int64  `gorm:"column:uid;default:0;NOT NULL" json:"uid"`                     // 用户UID
	WalletId    int64  `gorm:"column:wallet_id;default:0;NOT NULL" json:"wallet_id"`         // 收/支钱包ID
	TradeNo     string `gorm:"column:trade_no;NOT NULL" json:"trade_no"`                     // 自编唯一交易编号
	AssocId     int64  `gorm:"column:assoc_id;default:0;NOT NULL" json:"assoc_id"`           // 关联ID
	TradeType   int    `gorm:"column:trade_type;default:0;NOT NULL" json:"trade_type"`       // 收支类型 1-收入 2-支出
	TradeTypeId int    `gorm:"column:trade_type_id;default:0;NOT NULL" json:"trade_type_id"` // 交易类型ID
	TradeAmount int64  `gorm:"column:trade_amount;default:0;NOT NULL" json:"trade_amount"`   // 交易金额
	Balance     int64  `gorm:"column:balance;default:0;NOT NULL" json:"balance"`             // 交易前账户余额
	PayStatus   int    `gorm:"column:pay_status;default:0;NOT NULL" json:"pay_status"`       // 支付状态 0-未支付 1-支付中 2-已支付 3-支付失败
	Description string `gorm:"column:description;NOT NULL" json:"description"`               // 描述信息
}

type RedEnvelopeReceiver struct {
	entity.GormEntityTs
	ReceiverId  int64 `gorm:"column:receiver_id;primary_key" json:"receiver_id"`          // 领取ID
	EnvId       int64 `gorm:"column:env_id;default:0;NOT NULL" json:"env_id"`             // 红包ID
	ReceiverUid int64 `gorm:"column:receiver_uid;default:0;NOT NULL" json:"receiver_uid"` // 领取用户ID
}

type RedEnvelopeRecord struct {
	entity.GormEntityTs
	RecordId       int64  `gorm:"column:record_id;primary_key" json:"record_id"`                    // 红包领取记录ID
	ReceiverUid    int64  `gorm:"column:receiver_uid;default:0;NOT NULL" json:"receiver_uid"`       // 领取用户ID
	EnvId          int64  `gorm:"column:env_id;default:0;NOT NULL" json:"env_id"`                   // 红包ID
	TradeNo        string `gorm:"column:trade_no;NOT NULL" json:"trade_no"`                         // 交易编号
	ReceiveAmount  int64  `gorm:"column:receive_amount;default:0;NOT NULL" json:"receive_amount"`   // 领取金额(分)
	RemainAmount   int64  `gorm:"column:remain_amount;default:0;NOT NULL" json:"remain_amount"`     // 红包剩余金额(分)
	RemainQuantity int64  `gorm:"column:remain_quantity;default:0;NOT NULL" json:"remain_quantity"` // 红包剩余数量
	ReceiveStatus  int    `gorm:"column:receive_status;default:0;NOT NULL" json:"receive_status"`   // 领取状态 0-领取中 1-成功领取 2-领取失败
}
