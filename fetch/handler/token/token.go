package token

import (
	"net/http"

	"github.com/OkyWiliarso/efishery-test/fetch/middleware"
	"github.com/OkyWiliarso/efishery-test/fetch/pkg/common"
	"github.com/OkyWiliarso/efishery-test/fetch/service/token"
)

type IToken interface {
	VerifyToken(w http.ResponseWriter, r *http.Request)
}

type TokenService struct {
	service token.ITokenService
}

func NewTokenHandler() IToken {
	return &TokenService{
		service: token.NewTokenService(),
	}
}

func (s *TokenService) VerifyToken(w http.ResponseWriter, r *http.Request) {
	bearerToken := r.Header.Get("Authorization")
	token := middleware.GetBearerToken(bearerToken)

	user, err := s.service.Verify(token)
	if err != nil {
		common.RespondError(w, http.StatusBadRequest, err.Error())
	}

	common.RespondJSON(w, http.StatusOK, user)
}
