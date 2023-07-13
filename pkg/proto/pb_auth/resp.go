package pb_auth

func (r *SignUpResp) Set(code int32, msg string) {
	r.Code = code
	r.Msg = msg
}

func (r *SignInResp) Set(code int32, msg string) {
	r.Code = code
	r.Msg = msg
}

func (r *SignOutResp) Set(code int32, msg string) {
	r.Code = code
	r.Msg = msg
}

func (r *RefreshTokenResp) Set(code int32, msg string) {
	r.Code = code
	r.Msg = msg
}

func (r *GithubOAuth2CallbackResp) Set(code int32, msg string) {
	r.Code = code
	r.Msg = msg
}
