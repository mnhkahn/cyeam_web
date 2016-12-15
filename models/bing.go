package models

import (
	"encoding/xml"

	"github.com/astaxie/beego/httplib"
)

func GetBing() string {
	v := Bing{}
	req := httplib.Get("http://www.bing.com/HPImageArchive.aspx?format=json&idx=0&n=1")
	err := req.ToXml(&v)
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
