package models

import (
	"cyeam/structs"
	"fmt"
	"io"
	"log"
	"net/url"
	"strconv"
	"strings"
	"time"

	"cyeam/Godeps/_workspace/src/github.com/PuerkitoBio/goquery"
	"cyeam/Godeps/_workspace/src/github.com/astaxie/beego/httplib"
	cache "cyeam/Godeps/_workspace/src/github.com/patrickmn/go-cache"
	"cyeam/Godeps/_workspace/src/golang.org/x/net/html"
)

var (
	Resolution = map[string]bool{"1080p": true, "DVDScr": true, "720p": true}
)

func Rarbg(body io.ReadCloser) (*structs.RarbgView, error) {
	raw_document, err := html.Parse(body)
	if err != nil {
		return nil, err
	}
	// log.Println("AAAAAA")
	res := new(structs.RarbgView)
	res.Items = make([]*structs.RarbgResult, 0, 25)

	document := goquery.NewDocumentFromNode(raw_document)

	document.Find(`body > table:nth-child(6) > tbody > tr > td:nth-child(2) > div > table > tbody > tr:nth-child(2) > td > table.lista2t > tbody > tr.lista2`).Each(func(i int, s *goquery.Selection) {
		r := new(structs.RarbgResult)

		s.Find("td").Each(func(i int, td *goquery.Selection) {
			switch i {
			case 1:
				r.Name, r.OriginalTitle, r.Figure, r.Format, r.Rate = DoubanName(td.Text())
				// r.Name += "===" + td.Text()
				href, _ := td.Find("a").Attr("href")
				r.ID = strings.TrimLeft(href, "/torrent/")
			case 2:
				r.Date = td.Text()
			case 3:
				r.Size = td.Text()
			}

		})

		// log.Println(i, s.Text())
		res.Items = append(res.Items, r)
	})

	minPage, maxPage, curPage := 0, 0, 0
	document.Find("#pager_links").Each(func(i int, s *goquery.Selection) {
		if i == 0 {
			curPage, _ = strconv.Atoi(strings.TrimSpace(s.Find("b").Text()))

			s.Find("a").Each(func(i int, ss *goquery.Selection) {
				t := strings.TrimSpace(ss.Text())
				page, _ := strconv.Atoi(t)
				if page == 0 {
					return
				}

				if minPage > page {
					minPage = page
				} else if maxPage < page {
					maxPage = page
				}

			})
		}
	})

	for i := minPage; i <= maxPage; i++ {
		nav := new(structs.Nav)

		nav.Url = navUrl(i)
		// log.Println(nav.Url, i)

		nav.Page = i
		if i == curPage {
			nav.Selected = true
		}

		res.Pages = append(res.Pages, nav)
	}
	// page := s.Text()
	prePage := curPage - 1
	if prePage <= 0 {
		prePage = 1
	}
	res.Pre = new(structs.Nav)
	res.Pre.Url = navUrl(prePage)
	nextPage := curPage + 1
	res.Next = new(structs.Nav)
	res.Next.Url = navUrl(nextPage)

	// log.Println(curPage, minPage, maxPage)

	return res, nil
}

func navUrl(page int) string {
	return "/rarbg?category=14;17;42;44;45;46;47;48&order=seeders&by=DESC&page=" + strconv.Itoa(page)
}

var (
	DoubanCache = cache.New(365*24*time.Hour, 30*time.Second)
)

func DoubanName(engName string) (string, string, string, string, float64) {
	var figure, name, originName, format string
	var rate float64
	for r, _ := range Resolution {
		if i := strings.Index(engName, r); i >= 0 {
			name = engName[:i]
			name = strings.Replace(name, ".", " ", -1)
			originName = engName[:i]
			dd := douban(name)
			if dd.Count > 0 {
				name = dd.Subjects[0].Title
				figure = dd.Subjects[0].Images.Large
				rate = dd.Subjects[0].Rating.Average
			}
			format = r
		}
	}

	if len(name) == 0 {
		name = engName
	}

	return name, originName, figure, format, rate
}

func Douban(name string) *structs.DoubanResult {
	return douban(name)
}

// http://api.douban.com/v2/movie/search?q=Vikings
func douban(name string) *structs.DoubanResult {
	// name = "Doctor Strange"
	name = url.QueryEscape(name)

	if v, exists := DoubanCache.Get(name); exists {
		return v.(*structs.DoubanResult)
	}

	res := new(structs.DoubanResult)
	// log.Println(name, DoubanCache.ItemCount())

	u := "http://api.douban.com/v2/movie/search?q=" + name
	req := httplib.Get(u).Debug(true)
	err := req.ToJSON(res)
	if err != nil {
		raw, _ := req.String()
		log.Println(err, u, raw)
		return res
	}

	DoubanCache.SetDefault(name, res)

	return res
}

func RarbgTorrent(body io.ReadCloser) (string, error) {
	raw_document, err := html.Parse(body)
	if err != nil {
		return "", err
	}

	torrent := ""
	document := goquery.NewDocumentFromNode(raw_document)

	herf, exits := document.Find(`body > table:nth-child(6) > tbody > tr > td:nth-child(2) > div > div > table > tbody > tr:nth-child(2) > td > div > table > tbody > tr:nth-child(1) > td.lista > a:nth-child(2)`).Attr("href")
	if !exits {
		return "", fmt.Errorf("Can't find download url.")
	}
	torrent = "https://rarbg.to" + herf

	return torrent, nil
}
