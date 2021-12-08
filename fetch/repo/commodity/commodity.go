package commodity

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/OkyWiliarso/efishery-test/fetch/entity"
)

type ICommodityRepo interface {
	FetchCommodity() (res []entity.Commodity, err error)
}

type CommodityRepo struct {
}

func NewCommodityRepo() ICommodityRepo {
	return &CommodityRepo{}
}

func (repo *CommodityRepo) FetchCommodity() (res []entity.Commodity, err error) {
	r, err := http.Get("https://stein.efishery.com/v1/storages/5e1edf521073e315924ceab4/list")
	if err != nil {
		log.Println(err)
		return res, err
	}
	defer r.Body.Close()

	err = json.NewDecoder(r.Body).Decode(&res)
	if err != nil {
		log.Println(err)
		return res, err
	}

	return res, nil
}
