package models

import (
	"encoding/xml"

	"github.com/mnhkahn/gogogo/util"
)

func GetBing() string {
	v := Bing{}
	err := util.HttpXml("GET", "http://www.bing.com/HPImageArchive.aspx?format=json&idx=0&n=1", "", nil, &v)
	if len(v.Images) > 0 && err == nil {
		return bingURL + v.Images[0].Url
	}

	return "http://cyeam.qiniudn.com/zhonghuan.jpg"
}

const bingURL = `http://cn.bing.com`

type Bing struct {
	XMLName xml.Name `xml:"images"`
	Images  []Image  `xml:"image"`
}

type Image struct {
	XMLName       xml.Name `xml:"image"`
	Startdate     string   `xml:"startdate"`
	Fullstartdate string   `xml:"fullstartdate"`
	Url           string   `xml:"url"`
}
