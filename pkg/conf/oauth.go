package conf

type OAuth struct {
	ClientId     string // 对应 Client ID
	ClientSecret string // 对应 Client Secret
	RedirectUrl  string // 对应 Authorization callback URL
}
