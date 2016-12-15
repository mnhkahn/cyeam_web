/*
 * @Author: lichao115
 * @Date: 2016-12-15 17:28:46
 * @Last Modified by: lichao115
 * @Last Modified time: 2016-12-15 17:51:17
 */
package models

import (
	"encoding/xml"
	"fmt"
	"regexp"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
)

func GetDoodle() CyeamDoodle {
	req := httplib.Get("http://www.google.com/doodles/doodles.xml")
	req.SetTimeout(time.Duration(5)*time.Second, time.Duration(5)*time.Second)
	req.Debug(beego.AppConfig.String("runmode") == "dev")
	contents, err := req.Bytes()

	if err != nil {
		fmt.Println(err)
	}
	return ParseDoodle(contents)
}

func ParseDoodle(contents []byte) CyeamDoodle {
	v := Rss{}
	xml.Unmarshal(contents, &v)

	re_img := regexp.MustCompile("<img.*src=(.*?)[^>]*?>")
	img := re_img.FindAllString(v.Channel.Items[0].Description, -1)

	re_src := regexp.MustCompile("src=\"?(.*?)(\"|>|\\s+)")
	src := re_src.FindString(img[0])

	url := "http:" + src[5:len(src)-1]

	cy := CyeamDoodle{}
	cy.Title = v.Channel.Items[0].Title
	cy.Doodle = url

	return cy
}

type CyeamDoodle struct {
	Title  string `json:"title"`
	Doodle string `json:"doodle"`
}

type Rss struct {
	Channel Channel `xml:"channel"`
}

type Channel struct {
	Title string        `xml:"title"`
	Items []ChannelItem `xml:"item"`
}

type ChannelItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
}
