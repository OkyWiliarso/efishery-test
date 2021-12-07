package main

import (
	"fmt"

	"github.com/OkyWiliarso/efishery-test/fetch/app"
	"github.com/OkyWiliarso/efishery-test/fetch/config"
)

func main() {
	config.InitConfig()

	app := &app.App{}
	app.Initialize()

	listenPort := fmt.Sprintf(":%s", config.PORT)
	fmt.Println("Server running on port", listenPort)

	app.Run(listenPort)
}
