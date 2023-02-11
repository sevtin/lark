package xgin

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"lark/pkg/common/xjwt"
	"lark/pkg/utils"
)

func GetToken(ctx *gin.Context) (token string, expire int64) {
	var (
		tk    *jwt.Token
		key   string
		value interface{}
		exp   int
		err   error
	)
	token, err = ctx.Cookie("jwt")
	if err != nil {
		return
	}
	tk, err = xjwt.ParseFromCookie(ctx)
	if err != nil {
		return
	}
	for key, value = range tk.Claims.(jwt.MapClaims) {
		if key == "exp" {
			exp, _ = utils.ToInt(value)
		}
	}
	expire = int64(exp)
	return
}
