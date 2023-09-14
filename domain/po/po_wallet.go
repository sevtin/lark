package po

import "lark/pkg/entity"

type Wallet struct {
	entity.GormEntityTs
	WalletId     int64  `gorm:"column:wallet_id;primary_key" json:"wallet_id"`                // 钱包唯一ID
	WalletType   int    `gorm:"column:wallet_type;default:0;NOT NULL" json:"wallet_type"`     // 钱包类型
	Uid          int64  `gorm:"column:uid;default:0;NOT NULL" json:"uid"`                     // 用户UID
	Balance      int64  `gorm:"column:balance;default:0;NOT NULL" json:"balance"`             // 可用余额(balance+frozen_amount=总额)(分)
	FrozenAmount int64  `gorm:"column:frozen_amount;default:0;NOT NULL" json:"frozen_amount"` // 冻结金额(分)
	Status       int    `gorm:"column:status;default:0;NOT NULL" json:"status"`               // 钱包状态
	PayPassword  string `gorm:"column:pay_password;NOT NULL" json:"pay_password"`             // 支付密码
}

/*
type TransRecord struct {
	entity.GormEntityTs
	Tsn          int64  `gorm:"column:tsn;primary_key;default:0" json:"tsn"`                // 交易序列号 Transaction Number
	Uid          int64  `gorm:"column:uid;default:0;NOT NULL" json:"uid"`                   // 交易者UID
	FromWid      int64  `gorm:"column:from_wid;default:0;NOT NULL" json:"from_wid"`         // from 钱包ID
	ToWid        int64  `gorm:"column:to_wid;default:0;NOT NULL" json:"to_wid"`             // to 钱包ID
	ToAccount    string `gorm:"column:to_account;NOT NULL" json:"to_account"`               // to 账户
	AccountType  int    `gorm:"column:account_type;default:0;NOT NULL" json:"account_type"` // to_account账户类型 1:lark钱包 2:银行卡 3:支付宝 4:微信
	Amount       int64  `gorm:"column:amount;default:0;NOT NULL" json:"amount"`             // 交易额
	ExcRate      string `gorm:"column:exc_rate;default:0.00;NOT NULL" json:"exc_rate"`      // 对换率
	ExcAmount    int64  `gorm:"column:exc_amount;default:0;NOT NULL" json:"exc_amount"`     // 对换目标额
	Status       int    `gorm:"column:status;default:0;NOT NULL" json:"status"`             // 交易状态 0:待支付 1:已支付/已完成 2:已取消 3:失败
	TraderRole   int    `gorm:"column:trader_role;default:0;NOT NULL" json:"trader_role"`   // 交易者的角色 0:付款人 1:收款人
	TransType    int64  `gorm:"column:trans_type;default:0;NOT NULL" json:"trans_type"`     // 交易类型 1:转账支出 2:转账收入 3:兑换支出 4:兑换收入 5:交易支付 6:交易收款 7:提现
	Appid        string `gorm:"column:appid;NOT NULL" json:"appid"`                         // APPID
	MchId        string `gorm:"column:mchId;NOT NULL" json:"mchId"`                         // 商户ID(不对前端开放)
	OutTradeNo   string `gorm:"column:out_trade_no;NOT NULL" json:"out_trade_no"`           // 系统内部唯一订单号，只能是数字、大小写字母_-*
	TradeNo      string `gorm:"column:trade_no;NOT NULL" json:"trade_no"`                   // 该交易在第三方支付系统中的交易流水号
	Description  string `gorm:"column:description;NOT NULL" json:"description"`             // 交易内容描述
	Attach       string `gorm:"column:attach;NOT NULL" json:"attach"`                       // 自定义数据
	NotifyUrl    string `gorm:"column:notify_url;NOT NULL" json:"notify_url"`               // 回调通知url
	NotifyResult string `gorm:"column:notify_result;NOT NULL" json:"notify_result"`         // 回调内容
	NotifyTs     int64  `gorm:"column:notify_ts;default:0;NOT NULL" json:"notify_ts"`       // 回调时间
}
*/
