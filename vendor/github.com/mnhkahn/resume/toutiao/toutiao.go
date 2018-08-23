package toutiao

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/franela/goreq"
	"github.com/mnhkahn/resume/structs"
	"golang.org/x/net/html"
)

const (
	TOUTIAO_HOST = "https://toutiao.io"
)

func TouTiaoHotArticles(params *structs.Params) (*structs.TouTiao, []*structs.Article, error) {
	res, err := goreq.Request{Uri: TOUTIAO_HOST + "/subjects/" + params.TouTiao}.Do()
	if err != nil {
		return nil, nil, err
	}

	raw_document, err := html.Parse(res.Body)
	if err != nil {
		return nil, nil, err
	}

	articles := make([]*structs.Article, 0, params.TouTiaoLimit)
	toutiao := new(structs.TouTiao)
	document := goquery.NewDocumentFromNode(raw_document)

	toutiao.Name = strings.TrimSpace(document.Find("#main > div > div.text-center.m-b > h3").Text())
	toutiao.ID = params.TouTiao

	document.Find(`div.posts > div.post`).Each(func(i int, s *goquery.Selection) {
		article := new(structs.Article)
		article.Title = strings.TrimSpace(s.Find("h3.title").Text())
		u, _ := s.Find("div.content").Attr("data-url")
		article.Url = TOUTIAO_HOST + strings.TrimSpace(u)

		articles = append(articles, article)
	})

	if len(articles) > params.TouTiaoLimit {
		articles = articles[:params.TouTiaoLimit]
	}

	return toutiao, articles, nil
}
