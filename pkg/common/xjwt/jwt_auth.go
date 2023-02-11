package xjwt

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/golang-jwt/jwt/v4/request"
	"lark/pkg/common/xredis"
	"lark/pkg/constant"
	"lark/pkg/utils"
	"time"
)

const (
	JWT_KEY_ISS = "iss"
	JWT_KEY_EXP = "exp"
	JWT_KEY_IAT = "iat"
)

func CreateToken(uid int64, platform int32, access bool, expiresIn int) (t *JwtToken, err error) {
	t = new(JwtToken)
	var (
		token     *jwt.Token
		claims    jwt.MapClaims
		now       = time.Now()
		sessionId string
		expire    int64
		tokenStr  string
	)
	token = jwt.New(jwt.SigningMethodHS256)
	claims = token.Claims.(jwt.MapClaims)
	sessionId = utils.MD5(utils.NewUUID())

	expire = now.Add(time.Duration(expiresIn) * time.Second).Unix()
	claims[JWT_KEY_ISS] = constant.JWT_ISSUER
	claims[JWT_KEY_EXP] = expire
	claims[JWT_KEY_IAT] = now.Unix()
	claims[constant.USER_JWT_SESSION_ID] = sessionId
	claims[constant.USER_UID] = utils.Int64ToStr(uid)
	claims[constant.USER_PLATFORM] = platform

	tokenStr, err = token.SignedString([]byte(constant.JWT_TOKEN_SECRET_KEY))
	if err != nil {
		return
	}
	if access == true {
		tokenStr = constant.JWT_PREFIX + tokenStr
		// TODO: 开发调试用!!!
		key := constant.RK_SYNC_USER_ACCESS_TOKEN + utils.Int64ToStr(uid) + ":" + utils.Int32ToStr(platform)
		xredis.Set(key, tokenStr, constant.CONST_DURATION_JWT_ACCESS_TOKEN_EXPIRE_IN_SECOND)
	}
	t.SessionId = sessionId
	t.Expire = expire
	t.Token = tokenStr
	t.Uid = uid
	t.Platform = platform
	return
}

func ParseFromHeader(ctx *gin.Context) (res *jwt.Token, err error) {
	res, err = request.ParseFromRequest(ctx.Request, request.AuthorizationHeaderExtractor,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(constant.JWT_TOKEN_SECRET_KEY), nil
		})
	if err == request.ErrNoTokenInRequest {
		token := ctx.Query("token")
		res, err = jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
			return []byte(constant.JWT_TOKEN_SECRET_KEY), nil
		})
	}
	return
}

func ParseFromCookie(ctx *gin.Context) (*jwt.Token, error) {
	token, _ := ctx.Cookie("jwt")
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(constant.JWT_TOKEN_SECRET_KEY), nil
	})
}

func ParseFromToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(constant.JWT_TOKEN_SECRET_KEY), nil
	})
}

func Decode(in string) (t *JwtToken, err error) {
	t = new(JwtToken)
	var (
		token *jwt.Token
	)
	token, err = jwt.Parse(in, func(t *jwt.Token) (interface{}, error) {
		return []byte(constant.JWT_TOKEN_SECRET_KEY), nil
	})
	if err != nil {
		return
	}
	for key, value := range token.Claims.(jwt.MapClaims) {
		switch key {
		case constant.USER_JWT_SESSION_ID:
			switch value.(type) {
			case string:
				t.SessionId = value.(string)
			}
		case JWT_KEY_EXP:
			switch value.(type) {
			case float64:
				t.Expire = int64(value.(float64))
			}
		case constant.USER_UID:
			switch value.(type) {
			case string:
				t.Uid, _ = utils.ToInt64(value)
			}
		case constant.USER_PLATFORM:
			switch value.(type) {
			case float64:
				t.Platform = int32(value.(float64))
			}
		}
	}
	return
}

/*
Payload载荷：
jti：该jwt的唯一标识
iss：该jwt的签发者
iat：该jwt的签发时间
aud：该jwt的接收者
sub：该jwt的面向的用户
nbf：该jwt的生效时间,可不设置,若设置,一定要大于当前Unix UTC,否则token将会延迟生效
exp：该jwt的过期时间
*/
