package models

import (
	"fmt"

	"github.com/mnhkahn/gogogo/logger"
	"github.com/mnhkahn/gogogo/util"
)

type Weather struct {
	Dataseries  []int     `json:"dataseries"`
	DescriptNow string    `json:"descript_now"`
	Location    []float64 `json:"location"`
	NearestRain struct {
		Direction float64 `json:"direction"`
		Distance  float64 `json:"distance"`
		Intensity int     `json:"intensity"`
	} `json:"nearest_rain"`
	Previous   []int   `json:"previous"`
	Product    string  `json:"product"`
	ServerTime float64 `json:"server_time"`
	Skycon     string  `json:"skycon"`
	Source     string  `json:"source"`
	Station    string  `json:"station"`
	Status     string  `json:"status"`
	Summary    string  `json:"summary"`
	Temp       int     `json:"temp"`
}

const (
	CAIYUNAPP_WEATHER_URL = "http://caiyunapp.com/fcgi-bin/v1/api.py?lonlat=%s,%s&format=json&product=minutes_prec&token=%s"
	CAIYUNAPP_TOKEN       = "TAkhjf8d1nlSlspN"
)

func NewWeather(latitude, longitude string) *Weather {
	w := new(Weather)

	err := util.HttpJson("GET", fmt.Sprintf(CAIYUNAPP_WEATHER_URL, longitude, latitude, CAIYUNAPP_TOKEN), "", nil, w)
	if err != nil {
		logger.Warn(err)
	}

	return w
}
