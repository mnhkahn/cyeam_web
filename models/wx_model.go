/*
 * @Author: lichao115
 * @Date: 2016-12-15 17:27:18
 * @Last Modified by:   lichao115
 * @Last Modified time: 2016-12-15 17:27:18
 */
package models

import (
	"encoding/xml"
	"time"
)

type msgBase struct {
	ToUserName   CDATA
	FromUserName CDATA
	CreateTime   time.Duration
	MsgType      CDATA
	Content      CDATA
}

type CDATA struct {
	Data string `xml:",cdata"`
}

type Request struct {
	XMLName                xml.Name `xml:"xml"`
	msgBase                         // base struct
	Location_X, Location_Y float32
	Scale                  int
	Label                  string
	PicUrl                 CDATA
	MsgId                  int64
}

type Response struct {
	XMLName xml.Name `xml:"xml"`
	msgBase
	ArticleCount int       `xml:",omitempty"`
	Articles     *Articles `xml:",omitempty"`
	// FuncFlag     int     `xml:",omitempty"`
}

type Articles struct {
	Items []*Item
}

type Item struct {
	XMLName     xml.Name `xml:"item"`
	Title       CDATA
	Description CDATA
	PicUrl      CDATA
	Url         CDATA
}

func (resp Response) Encode() (data []byte, err error) {
	data, err = xml.MarshalIndent(resp, "", "")
	// data = []byte("<xml><ToUserName><![CDATA[ovQatjr9KDDjbLIE5Cwp5ZNnQDts]]></ToUserName><FromUserName><![CDATA[gh_aa23c6563b3d]]></FromUserName><CreateTime>1481859430</CreateTime><MsgType><![CDATA[text]]></MsgType><Content><![CDATA[123]]></Content></xml>")
	return
}
