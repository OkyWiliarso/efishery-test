package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/OkyWiliarso/efishery-test/fetch/config"
	"github.com/OkyWiliarso/efishery-test/fetch/handler/commodity"
	"github.com/OkyWiliarso/efishery-test/fetch/middleware"
	"github.com/gorilla/mux"
)

func main() {
	config.InitConfig()

	r := mux.NewRouter()

	commodity := commodity.NewCommodityHandler()

	r.HandleFunc("/commodity/list", commodity.CommodityList).Methods("Get")
	r.HandleFunc("/commodity/aggregate", commodity.AggCommodityList).Methods("Get")
	r.Use(middleware.ValidateToken)

	listenPort := fmt.Sprintf(":%s", config.PORT)
	fmt.Println("Server running on port", listenPort)

	log.Fatal(http.ListenAndServe(listenPort, r))
}
