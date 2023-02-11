package ctrl_auth

import (
	"lark/apps/interfaces/internal/service/svc_auth"
)

type AuthCtrl struct {
	authService svc_auth.AuthService
}

func NewAuthCtrl(authService svc_auth.AuthService) *AuthCtrl {
	return &AuthCtrl{authService: authService}
}
