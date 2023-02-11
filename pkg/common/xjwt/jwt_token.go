package xjwt

type JwtToken struct {
	Token     string
	SessionId string
	Expire    int64
	Uid       int64
	Platform  int32
}
