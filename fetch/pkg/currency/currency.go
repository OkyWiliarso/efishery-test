package currency

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/OkyWiliarso/efishery-test/fetch/config"
	"github.com/OkyWiliarso/efishery-test/fetch/entity"
)

func GetRate(currency string) (float64, error) {
	rate := entity.Currency{}

	url := fmt.Sprintf("https://free.currconv.com/api/v7/convert?q=%s&compact=ultra&apiKey=%s", currency, config.ApiKey)
	res, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&rate)
	if err != nil {
		log.Println(err)
		return rate.Ratio, err
	}

	return rate.Ratio, nil
}
