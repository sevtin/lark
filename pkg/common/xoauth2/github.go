package xoauth2

import (
	"context"
	"github.com/google/go-github/v57/github"
	"golang.org/x/oauth2"
	oauth2github "golang.org/x/oauth2/github"
	"lark/pkg/common/xlog"
	"lark/pkg/conf"
)

func NewGithubOauthConfig(cfg *conf.GithubOAuth2) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     cfg.ClientId,
		ClientSecret: cfg.ClientSecret,
		RedirectURL:  cfg.RedirectUrl,
		Endpoint:     oauth2github.Endpoint,
	}
}

func GithubUserInfo(accessToken string, user string) (info *github.User, resp *github.Response, err error) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	info, resp, err = client.Users.Get(ctx, user)
	if err != nil {
		xlog.Warn(err.Error())
		return
	}
	return
}

func GithubUserFollowers(accessToken string, user string) (followers int, err error) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	var (
		info *github.User
		resp *github.Response
	)
	info, resp, err = client.Users.Get(ctx, user)
	if err != nil {
		xlog.Warn(err.Error())
		return
	}
	if resp == nil || resp.StatusCode != 200 {
		return
	}
	if info != nil {
		followers = info.GetFollowers()
	}
	return
}
