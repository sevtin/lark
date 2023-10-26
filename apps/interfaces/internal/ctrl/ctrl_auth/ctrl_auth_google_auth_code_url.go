package ctrl_auth

import (
	"github.com/gin-gonic/gin"
	"lark/pkg/constant"
	"net/http"
)

// http://localhost:8088/open/auth/google/auth_code_url
func (ctrl *AuthCtrl) GoogleAuthCodeURL(ctx *gin.Context) {
	url := ctrl.googleOauthConfig.AuthCodeURL(constant.GOOGLE_OAUTH_STATE)
	ctx.Redirect(http.StatusTemporaryRedirect, url)
}

func (ctrl *AuthCtrl) GithubAuthCodeURL(ctx *gin.Context) {
	url := ctrl.githubOauthConfig.AuthCodeURL(constant.GITHUB_OAUTH_STATE)
	ctx.Redirect(http.StatusTemporaryRedirect, url)
}
