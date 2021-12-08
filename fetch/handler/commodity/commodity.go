package commodity

import (
	"net/http"

	"github.com/OkyWiliarso/efishery-test/fetch/pkg/common"
	"github.com/OkyWiliarso/efishery-test/fetch/service/commodity"
)

type ICommodity interface {
	CommodityList(w http.ResponseWriter, r *http.Request)
}

type CommodityService struct {
	service commodity.ICommodityService
}

func NewCommodityHandler() ICommodity {
	return &CommodityService{
		service: commodity.NewCommodityService(),
	}
}

func (s *CommodityService) CommodityList(w http.ResponseWriter, r *http.Request) {
	commodities, err := s.service.GetAllCommodity()
	if err != nil {
		common.RespondError(w, http.StatusBadRequest, err.Error())
	}

	common.RespondJSON(w, http.StatusOK, commodities)
}
