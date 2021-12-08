package commodity

import (
	"fmt"
	"log"
	"strconv"

	"github.com/OkyWiliarso/efishery-test/fetch/entity"
	"github.com/OkyWiliarso/efishery-test/fetch/pkg/currency"
	"github.com/OkyWiliarso/efishery-test/fetch/pkg/utils"
	"github.com/OkyWiliarso/efishery-test/fetch/repo/commodity"
)

type ICommodityService interface {
	GetAllCommodity() (res []entity.CommodityUSD, err error)
	GetAggCommodity() (res []entity.AggCommodity, err error)
}

type CommodityService struct {
	CommodityRepo commodity.ICommodityRepo
}

func NewCommodityService() ICommodityService {
	return &CommodityService{
		CommodityRepo: commodity.NewCommodityRepo(),
	}
}

type tempComm struct {
	provinsi   string
	minggu     string
	totalPrice int
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

	for _, v := range commodities {
		if v.Komoditas == "" || v.Price == "" {
			continue
		}
		price, _ := strconv.ParseFloat(v.Price, 64)

		usdPrice := price * rate

		res = append(res, entity.CommodityUSD{
			UUID:         v.UUID,
			Komoditas:    v.Komoditas,
			AreaProvinsi: v.AreaProvinsi,
			AreaKota:     v.AreaKota,
			Size:         v.Size,
			Price:        v.Price,
			TglParsed:    v.TglParsed,
			Timestamp:    v.Timestamp,
			USDPrice:     fmt.Sprintf("$%f", usdPrice),
		})
	}

	return res, nil
}

func (s *CommodityService) GetAggCommodity() (res []entity.AggCommodity, err error) {
	var commodities []entity.Commodity
	var newComm []tempComm
	var result []entity.AggCommodity

	commodities, err = s.CommodityRepo.FetchCommodity()
	if err != nil || commodities == nil {
		log.Println(err)
		return res, err
	}

	for _, v := range commodities {
		if v.Komoditas == "" || v.TglParsed == "" || v.Price == "" || v.Size == "" || v.AreaProvinsi == "" {
			continue
		}
		timestamp := utils.ParseTime(v.TglParsed)
		_, week := timestamp.ISOWeek()
		price, _ := strconv.Atoi(v.Price)
		size, _ := strconv.Atoi(v.Size)
		total := price * size

		temp := tempComm{
			provinsi:   v.AreaProvinsi,
			minggu:     strconv.Itoa(week),
			totalPrice: total,
		}

		newComm = append(newComm, temp)
	}

	mapComm := make(map[string]map[string]int)

	for _, v := range newComm {
		if provinsi, ok := mapComm[v.provinsi]; !ok {
			minggu := map[string]int{v.minggu: v.totalPrice}
			mapComm[v.provinsi] = minggu
		} else {
			if profit, ok := provinsi[v.minggu]; !ok {
				provinsi[v.minggu] = v.totalPrice
			} else {
				provinsi[v.minggu] += profit
			}
		}
	}

	for key, v := range mapComm {
		data := entity.AggCommodity{
			Provinsi: key,
			Minggu:   v,
			Max:      utils.Max(v),
			Min:      utils.Min(v),
			Avg:      utils.Average(v),
			Median:   utils.Median(v),
		}

		result = append(result, data)
	}

	return result, nil
}
