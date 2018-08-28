package models

import (
	"regexp"

	"github.com/mnhkahn/gogogo/logger"
	"github.com/mnhkahn/gogogo/util"
)

func GetDoodle() CyeamDoodle {
	v := Rss{}
	err := util.HttpXml("GET", "http://www.google.com/doodles/doodles.xml", "", nil, &v)
	if err != nil {
		logger.Warn(err)
	}

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
