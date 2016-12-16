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
	PicUrl                 string
	MsgId                  int
}

type Response struct {
	XMLName xml.Name `xml:"xml"`
	msgBase
	ArticleCount int     `xml:",omitempty"`
	Articles     []*Item `xml:"Articles>item,omitempty"`
	FuncFlag     int     `xml:",omitempty"`
}

type Item struct {
	XMLName     xml.Name `xml:"item"`
	Title       string
	Description string
	PicUrl      string
	Url         string
}

func (resp Response) Encode() (data []byte, err error) {
	data, err = xml.MarshalIndent(resp, "", "")
	return
}
