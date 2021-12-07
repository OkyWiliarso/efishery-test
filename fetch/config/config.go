package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	PORT string
)

func InitConfig() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(fmt.Errorf("fatal error read config file: %w", err))
	}

	PORT = os.Getenv("PORT")
}
