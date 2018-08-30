// Package search
package search

import (
	"cyeam/structs"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"

	"github.com/mnhkahn/gogogo/logger"

	"github.com/mnhkahn/maodou"
	"github.com/mnhkahn/maodou/models"
)

var blogCrawler *CyeamBlogCrawler

func Peanut(q string, page, size int) *structs.SearchResult {
	if page <= 0 {
		page = 1
	}
	if size <= 0 {
		size = 10
	}

	start := time.Now()

	se := structs.NewSearchResult()

	if q == "*" {
		q = "golang"
	}

	cnt, _, res := blogCrawler.Dao.Search(q, size, (page-1)*size)

	se.Summary.Q = q
	se.Summary.NumDocs = cnt

	for i, _ := range res {
		doc := new(structs.Doc)
		doc.Title = res[i].Title
		doc.Link = res[i].Link
		doc.Des = res[i].Description
		doc.Figure = res[i].Figure
		doc.Date = res[i].ParseDate
		se.Docs = append(se.Docs, doc)
	}

	end := time.Now()
	se.Summary.Duration = int64(end.Sub(start) / time.Millisecond)

	return se
}

type CyeamBlogCrawler struct {
	maodou.MaoDou
}

func (this *CyeamBlogCrawler) Start() {
	resp, err := this.Cawl("http://blog.cyeam.com/all.html")
	if err != nil {
		logger.Warn(err)
	} else {
		this.Index(resp)
	}
}

func (this *CyeamBlogCrawler) Index(resp *maodou.Response) {
	resp.Doc(`#content > div > div > div > h2 > a`).Each(func(i int, s *goquery.Selection) {
		href, has := s.Attr("href")
		if has {
			resp, err := this.Cawl("http://blog.cyeam.com/" + href)
			if err != nil {
				logger.Warn(err)
			} else {
				this.Detail(resp)
			}
		}
	})
}

func (this *CyeamBlogCrawler) Detail(resp *maodou.Response) {
	res := new(models.Result)
	u, _ := url.Parse(resp.Url)
	res.Id = u.Path
	res.Title = resp.Doc("#content > div.page-header.c6 > h1").Text()
	res.Author = "Bryce"
	res.Figure, _ = resp.Doc("#content > div.row-fluid.post-full > div > p.figure.center > img").Attr("src")
	res.Link = resp.Url
	res.ParseDate = time.Now()
	res.Description = resp.Doc("#content > div.row-fluid.post-full > div > blockquote").Text()
	res.Detail = resp.Doc("#content > div.row-fluid.post-full > div > div.content").Text()
	_tag := resp.Doc("#content > div.row-fluid.post-full > div > ul:nth-child(7)").Text()
	tags := []string{}
	for _, _t := range strings.Split(_tag, " ") {
		_tt := strings.TrimSpace(_t)
		if len(_tt) > 1 {
			_, err := strconv.Atoi(_tt)
			if err != nil {
				tags = append(tags, strings.Split(strings.TrimSpace(_t), " ")[0])
			}
		}
	}
	res.Tags = strings.Join(tags, " ")
	_category := resp.Doc("#content > div.row-fluid.post-full > div > ul:nth-child(6) > li:nth-child(2) > a").Text()
	res.Category = strings.Split(strings.TrimSpace(_category), " ")[0]
	this.Result(res)
}

func (this *CyeamBlogCrawler) Result(result *models.Result) {
	err := this.Dao.AddResult(result)
	if err != nil {
		logger.Warn(err, result)
	}
}

func NewCyeamBlogCrawler() *CyeamBlogCrawler {
	c := new(CyeamBlogCrawler)
	c.MaoDou.Init()
	c.SetRate(24 * time.Hour)
	c.SetDao("peanut", "./peanut.db")
	return c
}

func InitMaodou() {
	logger.Info("start cyeam blog crawler.")
	blogCrawler = NewCyeamBlogCrawler()
	maodou.Register(blogCrawler)
}
