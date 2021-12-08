package entity

type Commodity struct {
	UUID         string `json:"uuid"`
	Komoditas    string `json:"komoditas"`
	AreaProvinsi string `json:"area_provinsi"`
	AreaKota     string `json:"area_kota"`
	Size         string `json:"size"`
	Price        string `json:"price"`
	TglParsed    string `json:"tgl_parsed"`
	Timestamp    string `json:"timestamp"`
}

type CommodityUSD struct {
	UUID         string `json:"uuid"`
	Komoditas    string `json:"komoditas"`
	AreaProvinsi string `json:"area_provinsi"`
	AreaKota     string `json:"area_kota"`
	Size         string `json:"size"`
	Price        string `json:"price"`
	TglParsed    string `json:"tgl_parsed"`
	Timestamp    string `json:"timestamp"`
	USDPrice     string `json:"price_usd"`
}

type Currency struct {
	Ratio float64 `json:"IDR_USD"`
}
