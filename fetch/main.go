package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/OkyWiliarso/efishery-test/fetch/config"
	"github.com/OkyWiliarso/efishery-test/fetch/handler/commodity"
	"github.com/gorilla/mux"
)

func main() {
	config.InitConfig()

	r := mux.NewRouter()

	r.HandleFunc("/ping", Ping).Methods("GET")

	commodity := commodity.NewCommodityHandler()
	r.HandleFunc("/commodity/list", commodity.CommodityList).Methods("Get")

	listenPort := fmt.Sprintf(":%s", config.PORT)
	fmt.Println("Server running on port", listenPort)

	log.Fatal(http.ListenAndServe(listenPort, r))
}

func Ping(w http.ResponseWriter, r *http.Request) {
	response, _ := json.Marshal("Pong!")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(string(response)))
}
