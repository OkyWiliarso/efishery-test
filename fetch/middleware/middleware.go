package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/OkyWiliarso/efishery-test/fetch/config"
	"github.com/OkyWiliarso/efishery-test/fetch/entity"
	"github.com/OkyWiliarso/efishery-test/fetch/pkg/common"
)

func ValidateToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bearerToken := r.Header.Get("Authorization")

		if bearerToken == "" {
			common.RespondError(w, http.StatusBadRequest, "Token Required")
			return
		}

		token := GetBearerToken(bearerToken)

		payload := entity.AuthBody{
			Token: token,
		}

		payloadJSON, _ := json.Marshal(payload)

		url := fmt.Sprintf("%s%s", config.AuthURL, "/token/verify")
		res, err := http.Post(url, "application/json", bytes.NewBuffer(payloadJSON))
		if err != nil {
			log.Println(err)
			common.RespondError(w, http.StatusBadRequest, "Auth service unreachable")
			return
		}
		defer res.Body.Close()

		if res.StatusCode == 200 {
			next.ServeHTTP(w, r)
		} else {
			common.RespondError(w, http.StatusUnauthorized, "Unauthorized")
			return
		}
	})
}

func CheckRole(bearerToken string, role string) bool {
	token := GetBearerToken(bearerToken)
	auth := entity.AuthResp{}

	payload := entity.AuthBody{
		Token: token,
	}

	payloadJSON, _ := json.Marshal(payload)

	url := fmt.Sprintf("%s%s", config.AuthURL, "/token/verify")
	res, err := http.Post(url, "application/json", bytes.NewBuffer(payloadJSON))
	if err != nil {
		log.Println(err)
		return false
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&auth)
	if err != nil {
		log.Println(err)
		return false
	}

	return auth.Data.Role == role
}

func GetBearerToken(bearerToken string) string {
	var token string

	if authHeader := bearerToken; authHeader != "" {
		token = strings.Replace(authHeader, "Bearer ", "", 1)
	}

	return token
}
