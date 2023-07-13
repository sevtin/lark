package conf

type GithubOAuth2 struct {
	ClientId     string `yaml:"client_id"`     // 对应 Client ID
	ClientSecret string `yaml:"client_secret"` // 对应 Client Secret
	RedirectUrl  string `yaml:"redirect_url"`  // 对应 Authorization callback URL
}
