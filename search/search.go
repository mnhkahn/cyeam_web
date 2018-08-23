package search

import (
	"cyeam/models"
	"cyeam/structs"
	"log"
	"time"

	"github.com/mnhkahn/swiftype"
)

var (
	SWIFTYPE        *swiftype.Client
	SWIFTYPE_APIKEY = "zbnPrzRmw-fgYPNk4t76"
	SWIFTYPE_HOST   = "api.swiftype.com"
	SWIFTYPE_ENGINE = "zhan-nei-sou-suo"
)

func InitSwiftype() error {
	SWIFTYPE = swiftype.NewClientWithApiKey(SWIFTYPE_APIKEY, SWIFTYPE_HOST)
	return nil
}

func Search(sp *swiftype.SearchParam) *structs.SearchResult {
	start := time.Now()

	se := structs.NewSearchResult()

	if sp.Q == "*" {
		blogs := models.GetCyeamBlog(sp)
		if blogs != nil && len(blogs.TitleItemChannel) > 0 {
			for i, _ := range blogs.TitleItemChannel {
				doc := new(structs.Doc)
				doc.Title = blogs.TitleItemChannel[i]
				doc.Link = blogs.LinkItemChannel[i]
				doc.Des = blogs.Info[i]
				doc.Figure = blogs.Figure[i]
				if doc.Figure == "" {
					doc.Figure = "http://cyeam.qiniudn.com/gopherswrench.jpg"
				}
				se.Docs = append(se.Docs, doc)
			}
			se.Summary.NumDocs = blogs.Count
		}
	} else {
		res, err := SWIFTYPE.Search(SWIFTYPE_ENGINE, sp)
		if err != nil {
			log.Println(err)
			return se
		}

		if res != nil {
			for _, page := range res.Records.Page {
				doc := new(structs.Doc)
				doc.Title = page.Title
				doc.Link = page.URL
				if len(page.Highlight.Body) != 0 {
					doc.Des = page.Highlight.Body
				} else {
					doc.Des = string([]rune(page.Body)[0:300])
				}
				doc.Figure = page.Image
				se.Docs = append(se.Docs, doc)
			}
			se.Summary.NumDocs = res.Info.Page.TotalResultCount
		}
	}

	se.Summary.Q = sp.Q

	end := time.Now()
	se.Summary.Duration = int64(end.Sub(start) / time.Millisecond)

	return se
}
