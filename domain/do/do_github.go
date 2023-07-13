package do

type GithubToken struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
}

type GithubUser struct {
	AvatarURL string `json:"avatar_url"`
	Email     string `json:"email"`
	ID        int    `json:"id"`
	Login     string `json:"login"`
	Name      string `json:"name"`
}
