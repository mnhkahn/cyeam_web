package models

import (
	"log"

	"github.com/astaxie/beego/httplib"
	"github.com/mnhkahn/swiftype"
)

func GetCyeamBlog(sp *swiftype.SearchParam) *CyeamBlog {
	v := new(CyeamBlog)
	req := httplib.Get(cyeamBlogURL)
	err := req.ToXML(v)
	if err != nil {
		log.Println(err)
	}

	v.Count = len(v.LinkItemChannel)

	start := (sp.Page - 1) * sp.PerPage
	end := start + sp.PerPage

	if end > len(v.TitleItemChannel) || start > len(v.TitleItemChannel) {
		start = 0
		end = 20
	}
	v.Figure = v.Figure[start:end]
	v.LinkItemChannel = v.LinkItemChannel[start:end]
	v.PubDateItemChannel = v.PubDateItemChannel[start:end]
	v.TitleItemChannel = v.TitleItemChannel[start:end]
	v.DescriptionItemChannel = v.DescriptionItemChannel[start:end]
	v.Info = v.Info[start:end]

	return v
}

const cyeamBlogURL = `http://blog.cyeam.com/rss.xml`

// go get github.com/wicast/xj2s/cmd/...
// curl -s http://blog.cyeam.com/rss.xml|xmljson2struct
type CyeamBlog struct {
	Ttl                string   `xml:"channel>ttl"`
	Figure             []string `xml:"channel>item>figure"`
	LinkItemChannel    []string `xml:"channel>item>link"`
	PubDateItemChannel []string `xml:"channel>item>pubDate"`
	Title              string   `xml:"channel>title"`
	Description        string   `xml:"channel>description"`
	TitleItemChannel   []string `xml:"channel>item>title"`
	// Link                   []string `xml:"channel>link"`
	Info                   []string `xml:"channel>item>info"`
	DescriptionItemChannel []string `xml:"channel>item>description"`
	// Guid                   []string `xml:"channel>item>guid"`
	Version       string `xml:"version,attr"`
	LastBuildDate string `xml:"channel>lastBuildDate"`
	PubDate       string `xml:"channel>pubDate"`
	Count         int    `xml:"-"`
}
