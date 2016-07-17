package main

import (
	"log"
	"strings"
	"time"

	"cyeam/Godeps/_workspace/src/github.com/PuerkitoBio/goquery"

	cygostr "cyeam/Godeps/_workspace/src/github.com/mnhkahn/cygo/strings"
	"cyeam/Godeps/_workspace/src/github.com/mnhkahn/maodou"
	"cyeam/Godeps/_workspace/src/github.com/mnhkahn/maodou/models"
)

type Haixiu struct {
	maodou.MaoDou
}

func (this *Haixiu) Start() {
	resp, err := this.Cawl("http://blog.cyeam.com")
	if err == nil {
		this.Index(resp)
	}
}

func (this *Haixiu) Index(resp *maodou.Response) {
	resp.Doc(`#content > div.row-fluid > div > div`).Each(func(i int, s *goquery.Selection) {
		href, has := s.Find("h2 > a").Attr("href")
		if has {
			log.Printf("===============%d==============\n", i)
			resp, err := this.Cawl("http://blog.cyeam.com" + href)
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
	// res.Id = strings.Split(resp.Url, "/")[5]
	res.Title = resp.Doc("#content > div > h1").Text()
	// res.Author = resp.Doc("#content > div > div.article > div.topic-content.clearfix > div.topic-doc > h3 > span.from > a").Text()
	res.Figure, _ = resp.Doc("#content > div.row-fluid.post-full > div > p.figure.center > img").Attr("src")
	res.Link = resp.Url
	// res.Source = "www.douban.com/group/haixiuzu/discussion"
	// res.ParseDate = cygo.Now()
	res.Category = strings.Split(strings.TrimSpace(resp.Doc("#content > div.row-fluid.post-full > div > ul:nth-child(5) > li:nth-child(2) > a").Text()), " ")[0]
	var tags []string
	resp.Doc(`#content > div.row-fluid.post-full > div > ul:nth-child(6) > li > a`).Each(func(i int, s *goquery.Selection) {
		tags = append(tags, strings.Split(s.Text(), " ")[0])
	})
	res.Tags = strings.Join(tags, ",")
	res.Description = cygostr.TrimAllSpace(resp.Doc("#content > div.row-fluid.post-full > div > blockquote").Text())
	res.Detail = cygostr.TrimAllSpace(resp.Doc("#content > div.row-fluid.post-full > div > div.content").Text())
	this.Result(res)
}

func (this *Haixiu) Result(result *models.Result) {
	this.Dao.AddResult(result)
}

func main() {
	haixiu := new(Haixiu)
	haixiu.Init()
	haixiu.SetRate(time.Duration(30) * time.Minute)
	// haixiu.SetProxy("xici", `{"max_cawl_cnt":12,"cnt":10,"min_cnt":1,"root":"http://www.douban.com/group/haixiuzu/discussion"}`)
	haixiu.SetDao("sqlite", "./data.db")
	haixiu.Dao.Debug(true)
	maodou.Register(haixiu)
}
