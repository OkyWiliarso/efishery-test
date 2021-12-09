package token

import (
	"log"

	"github.com/OkyWiliarso/efishery-test/fetch/config"
	"github.com/OkyWiliarso/efishery-test/fetch/entity"
	"github.com/golang-jwt/jwt"
	"github.com/mitchellh/mapstructure"
)

type ITokenRepo interface {
	Claims(token string) (res entity.User, err error)
}

type TokenRepo struct {
}

func NewTokenRepo() ITokenRepo {
	return &TokenRepo{}
}

func (repo *TokenRepo) Claims(tokenStr string) (res entity.User, err error) {
	hmacSecretString := config.Secret
	hmacSecret := []byte(hmacSecretString)
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return hmacSecret, nil
	})

	if err != nil {
		return res, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		err = mapstructure.Decode(claims, &res)
		if err != nil {
			log.Println(err)
			return res, err
		}
	}

	return res, nil
}
