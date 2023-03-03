package dto_auth

import (
	"lark/apps/interfaces/internal/dto/dto_user"
	"lark/pkg/proto/pb_enum"
)

type SignInReq struct {
	AccountType      int32                 `json:"account_type" binding:"required,oneof=1 2"`   // 登录类型 1:手机号 2:lark账户
	Platform         pb_enum.PLATFORM_TYPE `json:"platform" binding:"required,oneof=1 2 3 4 5"` // 平台 1:iOS 2:安卓
	Account          string                `json:"account" binding:"required,min=5,max=20"`     // 手机号/lark账户
	Udid             string                `json:"udid" binding:"required,len=40"`              // UDID
	VerificationCode string                `json:"verification_code" binding:"omitempty"`       // 验证码
	Password         string                `json:"password" binding:"required,len=32"`          // 密码
}

type AuthResp struct {
	AccessToken  *Token             `json:"access_token"`
	RefreshToken *Token             `json:"refresh_token"`
	UserInfo     *dto_user.UserInfo `json:"user_info"`
	Server       *ServerInfo        `json:"server"`
}

type Token struct {
	Token  string `json:"token"`  // 用户token
	Expire int64  `json:"expire"` // token过期时间戳（秒）
}

type ServerInfo struct {
	ServerId int64  `json:"server_id"` // 服务器ID
	Name     string `json:"name"`      // 服务器名
	Port     int    `json:"port"`      // 端口号
}
