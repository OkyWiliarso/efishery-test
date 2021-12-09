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

type AggCommodity struct {
	Provinsi string `json:"area_provinsi"`
	Minggu   map[string]int
	Max      float64 `json:"max"`
	Min      float64 `json:"min"`
	Avg      float64 `json:"average"`
	Median   float64 `json:"median"`
}

type AuthBody struct {
	Token string `json:"token"`
}

type AuthResp struct {
	Data User `json:"data"`
}

type User struct {
	Name      string `json:"name"`
	Role      string `json:"role"`
	Phone     string `json:"phone"`
	Timestamp string `json:"timestamp"`
}
