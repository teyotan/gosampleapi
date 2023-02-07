package authhandler

import (
	"gosampleapi/service/authservice"
)

type AuthHandler struct {
	AuthService *authservice.AuthService
}

func NewAuthHandler(svc *authservice.AuthService) *AuthHandler {
	return &AuthHandler{
		AuthService: svc,
	}
}
