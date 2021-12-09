package token

import (
	"log"

	"github.com/OkyWiliarso/efishery-test/fetch/entity"
	"github.com/OkyWiliarso/efishery-test/fetch/repo/token"
)

type ITokenService interface {
	Verify(token string) (res entity.User, err error)
}

type TokenService struct {
	TokenRepo token.ITokenRepo
}

func NewTokenService() ITokenService {
	return &TokenService{
		TokenRepo: token.NewTokenRepo(),
	}
}

func (s *TokenService) Verify(token string) (res entity.User, err error) {
	res, err = s.TokenRepo.Claims(token)
	if err != nil {
		log.Println(err)
		return res, err
	}

	return res, nil
}
