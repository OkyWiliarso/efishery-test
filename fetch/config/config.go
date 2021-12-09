package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	PORT    string
	ApiKey  string
	AuthURL string
	Secret  string
)

func InitConfig() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(fmt.Errorf("fatal error read config file: %w", err))
	}

	PORT = os.Getenv("PORT")
	ApiKey = os.Getenv("API_KEY")
	AuthURL = os.Getenv("AUTH_URL")
	Secret = os.Getenv("SECRET")
}
