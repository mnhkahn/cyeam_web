package main

import (
	"log"
	"strings"
	"time"

	"cyeam/Godeps/_workspace/src/github.com/PuerkitoBio/goquery"

	"cyeam/Godeps/_workspace/src/github.com/mnhkahn/maodou"
	"cyeam/Godeps/_workspace/src/github.com/mnhkahn/maodou/cygo"
	"cyeam/Godeps/_workspace/src/github.com/mnhkahn/maodou/models"
)

type Haixiu struct {
	maodou.MaoDou
}

func (this *Haixiu) Start() {
	resp, err := this.Cawl("http://www.douban.com/group/haixiuzu/discussion")
	if err == nil {
		this.Index(resp)
	}
}

func (this *Haixiu) Index(resp *maodou.Response) {
	resp.Doc(`#content > div > div.article > div:nth-child(2) > table > tbody > tr > td.title > a`).Each(func(i int, s *goquery.Selection) {
		href, has := s.Attr("href")
		if has {
			log.Printf("===============%d==============\n", i)
			resp, err := this.Cawl(href)
			if err == nil {
				this.Detail(resp)
			}
		}
	})
}

func (this *Haixiu) Detail(resp *maodou.Response) {
	if len(strings.Split(resp.Url, "/")) < 6 {
		return
	}
	res := new(models.Result)
	res.Id = strings.Split(resp.Url, "/")[5]
	res.Title = resp.Doc("#content > h1").Text()
	res.Author = resp.Doc("#content > div > div.article > div.topic-content.clearfix > div.topic-doc > h3 > span.from > a").Text()
	figures := []string{}
	resp.Doc("#link-report > div.topic-content > div.topic-figure.cc").Each(func(i int, s *goquery.Selection) {
		f, exists := s.Find("img").Attr("src")
		if exists {
			figures = append(figures, f)
		}
	})
	res.Figure = strings.Join(figures, ",")
	res.Link = resp.Url
	res.Source = "www.douban.com/group/haixiuzu/discussion"
	res.ParseDate = cygo.Now()
	res.Description = "haixiuzu"
	this.Result(res)
}

func (this *Haixiu) Result(result *models.Result) {
	if result.Figure != "" {
		this.Dao.AddResult(result)
	} else {
		log.Println("No pic for save.")
	}
}

func HaixiuzuStart() {
	haixiu := new(Haixiu)
	haixiu.Init()
	haixiu.SetRate(time.Duration(30) * time.Minute)
	haixiu.SetProxy("xici", `{"max_cawl_cnt":12,"cnt":10,"min_cnt":1,"root":"http://www.douban.com/group/haixiuzu/discussion"}`)
	haixiu.SetDao("duoshuo", `{"short_name":"cyeam","secret":"df66f048bd56cba5bf219b51766dec0d","thread_key":"haixiuzucyeam"}`)
	haixiu.Dao.Debug(true)
	maodou.Register(haixiu)
}