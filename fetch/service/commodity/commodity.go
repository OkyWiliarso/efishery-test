package commodity

import (
	"fmt"
	"log"
	"strconv"

	"github.com/OkyWiliarso/efishery-test/fetch/entity"
	"github.com/OkyWiliarso/efishery-test/fetch/pkg/currency"
	"github.com/OkyWiliarso/efishery-test/fetch/repo/commodity"
)

type ICommodityService interface {
	GetAllCommodity() (res []entity.CommodityUSD, err error)
}

type CommodityService struct {
	CommodityRepo commodity.ICommodityRepo
}

func NewCommodityService() ICommodityService {
	return &CommodityService{
		CommodityRepo: commodity.NewCommodityRepo(),
	}
}

func (s *CommodityService) GetAllCommodity() (res []entity.CommodityUSD, err error) {
	var commodities []entity.Commodity

	commodities, err = s.CommodityRepo.FetchCommodity()
	if err != nil || commodities == nil {
		log.Println(err)
		return res, err
	}

	rate, err := currency.GetRate("IDR_USD")
	if err != nil {
		log.Println(err)
		return res, err
	}

	for _, val := range commodities {
		if val.Price == "" {
			continue
		}
		price, _ := strconv.ParseFloat(val.Price, 64)

		usdPrice := price * rate

		res = append(res, entity.CommodityUSD{
			UUID:         val.UUID,
			Komoditas:    val.Komoditas,
			AreaProvinsi: val.AreaProvinsi,
			AreaKota:     val.AreaKota,
			Size:         val.Size,
			Price:        val.Price,
			TglParsed:    val.TglParsed,
			Timestamp:    val.Timestamp,
			USDPrice:     fmt.Sprintf("$%f", usdPrice),
		})
	}

	return res, nil
}
