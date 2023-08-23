package xtrade

import (
	"fmt"
	"lark/pkg/common/xsnowflake"
	"strings"
	"time"
)

const chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-_"
const base = 64

func GenerateTradeNo(channel string) string {
	builder := &strings.Builder{}
	// 渠道前缀
	builder.WriteString(channel)
	builder.WriteString("-")
	// 日期
	now := time.Now()
	builder.WriteString(now.Format("200601021504"))
	builder.WriteString("-")
	// 唯一符
	snowflakeID := xsnowflake.NewSnowflakeID()
	for snowflakeID > 0 {
		remainder := snowflakeID % base
		builder.WriteByte(chars[remainder])
		snowflakeID /= base
	}
	builder.WriteString("-")
	// 序号
	builder.WriteString(fmt.Sprintf("%06d", now.UnixMicro()%100000))
	return builder.String()
}

//func decimalToBase64(builder *strings.Builder, decimal int64) string {
//	for decimal > 0 {
//		remainder := decimal % base
//		builder.WriteByte(base64Charset[remainder])
//		decimal /= base
//	}
//	// 未反转字符串
//	return builder.String()
//}
