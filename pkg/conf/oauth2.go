package conf

type GithubOAuth2 struct {
	ClientId     string `yaml:"client_id"`     // 对应 Client ID
	ClientSecret string `yaml:"client_secret"` // 对应 Client Secret
	RedirectUrl  string `yaml:"redirect_url"`  // 对应 Authorization callback URL
}

type GoogleOAuth2 struct {
	ClientId     string   `yaml:"client_id"`     // 对应 Client ID
	ClientSecret string   `yaml:"client_secret"` // 对应 Client Secret
	RedirectUrl  string   `yaml:"redirect_url"`  // 对应 Authorization callback URL
	Scopes       []string `yaml:"scopes"`
	AuthURL      string   `yaml:"auth_url"`
	TokenURL     string   `yaml:"token_url"`
}
