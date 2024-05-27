package cache

type GithubEmailCache interface {
}

type githubEmailCache struct {
}

func NewGithubEmailCache() GithubEmailCache {
	return &githubEmailCache{}
}
