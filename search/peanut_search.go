// Package search
package search

import (
	"cyeam/structs"
	"net/url"
	"strconv"
	"strings"
	"time"

	"golang.org/x/oauth2/jwt"

	"github.com/mnhkahn/maodou/analytics"

	"github.com/PuerkitoBio/goquery"

	"github.com/mnhkahn/gogogo/logger"
	"github.com/mnhkahn/maodou"
	"github.com/mnhkahn/maodou/models"
)

var blogCrawler *CyeamBlogCrawler

func Peanut(q string, page, size int, sort string, asc bool) *structs.SearchResult {
	if page <= 0 {
		page = 1
	}
	if size <= 0 {
		size = 10
	}

	start := time.Now()

	se := structs.NewSearchResult()

	cnt, _, res := blogCrawler.Dao.Search(q, size, (page-1)*size, sort, asc)

	se.Summary.Q = q
	se.Summary.NumDocs = cnt

	for i, _ := range res {
		doc := new(structs.Doc)
		doc.Title = res[i].Title
		doc.Link = res[i].Link
		doc.Des = res[i].Description
		doc.Figure = res[i].Figure
		doc.Date = res[i].ParseDate
		doc.PV = res[i].PV
		se.Docs = append(se.Docs, doc)
	}

	end := time.Now()
	se.Summary.Duration = int64(end.Sub(start) / time.Microsecond)

	return se
}

type CyeamBlogCrawler struct {
	maodou.MaoDou
	analytics map[string]int
}

func (this *CyeamBlogCrawler) Index(resp *maodou.Response, jobs chan string) error {
	resp.Doc(`#content > div > div > div > h2 > a`).Each(func(i int, s *goquery.Selection) {
		href, has := s.Attr("href")
		if has {
			this.AddJob("http://blog.cyeam.com/" + href)
		}
	})
	return nil
}

func (this *CyeamBlogCrawler) Detail(resp *maodou.Response) (*models.Result, error) {
	var err error
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
	date := resp.Doc("#content > div.row-fluid.post-full > div > div.date > span:nth-child(2)").Text()
	res.CreateTime, err = time.Parse("02 January 2006", date)
	if err != nil {
		return res, err
	}

	if this.analytics != nil {
		// skip one slash //golang/2018/08/22/json-number
		res.PV = this.analytics[u.Path[1:]]
	}

	return res, nil
}

const analyticsLink = "https://www.googleapis.com/analytics/v3/data/ga?ids=ga%3A84991381&start-date=30daysAgo&end-date=yesterday&metrics=ga%3Apageviews&dimensions=ga%3ApagePath&key=https://www.googleapis.com/analytics/v3/data/ga?ids=ga%3A84991381&start-date=30daysAgo&end-date=yesterday&metrics=ga%3Apageviews&dimensions=ga%3ApagePath"
const crawlRoot = "http://blog.cyeam.com/all.html"

func NewCyeamBlogCrawler() *CyeamBlogCrawler {
	c := new(CyeamBlogCrawler)
	c.MaoDou.Init(crawlRoot)
	c.SetRate(24 * time.Hour)
	c.SetD(false)

	var err error
	c.analytics, err = analytics.Analytics(analyticsLink, &jwt.Config{
		Email: "peanut-pv@tensile-market-90013.iam.gserviceaccount.com",
		// The contents of your RSA private key or your PEM file
		// that contains a private key.
		// If you have a p12 file instead, you
		// can use openssl to export the private key into a pem file.
		//
		// $ openssl pkcs12 -in key.p12 -passin pass:notasecret -out key.pem -nodes
		//
		// The field only supports PEM containers with no passphrase.
		// The openssl command will convert p12 keys to passphrase-less PEM containers.
		//PrivateKeyID: "bbbf30e4201cf59ce0bbcb2f320830385a3978d6",
		PrivateKey: []byte("-----BEGIN PRIVATE KEY-----\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDK7EZnzQw0g9dX\nzY/1+Gnsswo/YFcb0DoUMIORWKf2HvMpqxAd5R9a+s6vWn1Its+C3HNkErMVTtTG\nLE//RB3XsT42rDR1zKZ+UhZjgUm+/nyTnVIyVcYodzh1skD2SHGXlJqd0l3VP/0V\nXumn/ShpZfel7M7D+YyA38v5+Dfqo5DbqO3aYlzYRAj+k+Hzo75QMxb36z+9IZ6X\nCv5Mt9XSbeMtWMAw1YxeZ8qSAkJc8suaJc4ThSzOIyM5U79U6JIkMufJ3QPCEZ0l\nLdixD8jAQooVC6COqb1Z7BtwQdOWKxAOVI5IGg51vp5b4k12rSyzeHPtjGVVykmR\n30fW4D8JAgMBAAECggEABcCfRFJd6LUzaxPy40xTUGLzmj//EtTweIxhhm1xdtAs\ny38ZGgfZ6jMsSHCebEEbw9VqWKlmYJHCBWDR3xDNbxphWEfBxZ8eDycA5yQ4n3dJ\nDR7gwCLfVwao1+lMNHnG7T7zXSBOtFINOMheg6Ubuqp/KTnRKLa6eBggCP+I3O1m\nK6UpUgu0WrwgjYEnUzRGRFOFZExcNzjwUHNLJ5FCwh9YLW6kwOsgT3heyUY8N3pj\nT6mwVAgbYxze9TIpuyVIaqf8LQ5dP3SzEjjW2jSPm/2cjEk9NKB47VhRcZH+KJQY\n9iI1HVyPpYiYtxAPf/6vcrK71Tj6j0rHlc4Nca/vKQKBgQDpW5AmCO5ZsngXuj2S\n1AauTElJJ8oIAaJ3szKUsuYp1doFpK4nRJbTuzROT43hbAfHHz3+Pd+pW57uqHuB\nsjPvkiiH6LLJPRAbU5HWhXUOcOm6xLnswHajCsR76puGBvrcjrxrVOnOP64v2fJL\niY734dhSdO4PyRGGxFXHBj90ZQKBgQDenLxBiA0MHNevLRcbj9Vyl/rt69CVMSEa\nEUHD5g3jph6+ia2304cCUlBFrEFTrvIeuRjSd9zw8OhfXWhWoDNOUTi4PxEP5ZWc\nJf+INzbOWhxEPzFp375zscJZm0TrWP/mBpwq0VXp5ICEO/f30EsW4zE5efn5jqZn\nu1RHsFrb1QKBgCb4zLUdbrj6LkZAK0JXOJppUR/vjjUSGNEG0160Fe5MsbGZlCAo\nu0u3CwA9FwPbp9zgYdkQ+kZtb7iJ2L6LRVMwRKaV/S3Qjd0SctuxxB/aSZ6QdkCM\n0ANgq/nJ75lNlx24lM0UDEwOpIeHTYjB+2d4h0kWECAAw3WPWof3iidlAoGBALza\nHuNBNkBmX5vfFtFtDllvEZOyEHvg+AITTcWRb4sHLOHcDyH6M3kGt87DuY/yxLjH\nsoUq5qcI2Tm+FnwW4C+6u/GinyjrTibwHX5DyRz6WSyUp6j4BaxEy2oVTTyTflR4\nmxfAC7CnB1gnP9BeRrWd++6UyjqqiAVMaM2AkTQZAoGBANW95OXeHdQ03vyWq7tG\n6FU5TCpW8qOhE6St+UoWLQ7M36GD4xMtjDxkAKac0nRdyh8601qqj/y5cAy0RSRO\n2y3WqrwigcTgGndZXuiQBu94ic6qKlliJpS1NgFOxC4X2geCTq7e2yZAaIxqAQ/1\nYypZcVIFC+bhzPPt8ajp5bVc\n-----END PRIVATE KEY-----\n"),
		Scopes: []string{
			"https://www.googleapis.com/auth/analytics.readonly",
		},
		//TokenURL: google.JWTTokenURL,
		// If you would like to impersonate a user, you can
		// create a transport with a subject. The following GET
		// request will be made on the behalf of user@example.com.
		// Optional.
		//Subject: "user@example.com",
		TokenURL: "https://oauth2.googleapis.com/token",
	})

	if err != nil {
		logger.Warn(err)
	}

	return c
}

func InitMaodou() {
	logger.Info("start cyeam blog crawler.")
	blogCrawler = NewCyeamBlogCrawler()
	maodou.Register(blogCrawler)
}
