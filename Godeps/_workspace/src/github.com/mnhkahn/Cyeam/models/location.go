package models

import (
	"cyeam/Godeps/_workspace/src/github.com/astaxie/beego/httplib"
	"fmt"
)

type BaiduLocation struct {
	Result struct {
		AddressComponent struct {
			City         string `json:"city"`
			Country      string `json:"country"`
			CountryCode  int    `json:"country_code"`
			Direction    string `json:"direction"`
			Distance     string `json:"distance"`
			District     string `json:"district"`
			Province     string `json:"province"`
			Street       string `json:"street"`
			StreetNumber string `json:"street_number"`
		} `json:"addressComponent"`
		Business         string `json:"business"`
		CityCode         int    `json:"cityCode"`
		FormattedAddress string `json:"formatted_address"`
		Location         struct {
			Lat float64 `json:"lat"`
			Lng float64 `json:"lng"`
		} `json:"location"`
		PoiRegions []interface{} `json:"poiRegions"`
		Pois       []struct {
			Addr      string `json:"addr"`
			Cp        string `json:"cp"`
			Direction string `json:"direction"`
			Distance  string `json:"distance"`
			Name      string `json:"name"`
			PoiType   string `json:"poiType"`
			Point     struct {
				X float64 `json:"x"`
				Y float64 `json:"y"`
			} `json:"point"`
			Tag string `json:"tag"`
			Tel string `json:"tel"`
			UID string `json:"uid"`
			Zip string `json:"zip"`
		} `json:"pois"`
		SematicDescription string `json:"sematic_description"`
	} `json:"result"`
	Status int `json:"status"`
}

const (
	BAIDU_AK           = "43E57D0f47ca6382344892b9b65ba0ad"
	BAIDU_LOCATION_URL = "http://api.map.baidu.com/geocoder/v2/?ak=%s&location=%s,%s&output=json&pois=1"
)

func NewBaiduLocation(latitude, longitude string) *BaiduLocation {
	loc := new(BaiduLocation)

	req := httplib.Get(fmt.Sprintf(BAIDU_LOCATION_URL, BAIDU_AK, latitude, longitude))
	req.ToJson(loc)

	return loc
}
