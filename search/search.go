/*
 * @Author: lichao115
 * @Date: 2016-12-16 12:53:05
 * @Last Modified by: mnhkahn <lichao@cyeam.com>
 * @Last Modified time: 2016-12-17 18:51:54
 */
package search

import (
	"cyeam/structs"
	"encoding/json"
	"log"
	"time"

	"cyeam/Godeps/_workspace/src/github.com/mnhkahn/swiftype"
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

func Search(q string) *structs.SearchResult {
	start := time.Now()

	se := structs.NewSearchResult()

	res := new(structs.SwiftypeResult)
	data := SWIFTYPE.Search(SWIFTYPE_ENGINE, q)
	err := json.Unmarshal(data, res)
	if err != nil {
		log.Println(err)
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
	}

	se.Summary.Q = q
	se.Summary.NumDocs = res.Info.Page.TotalResultCount

	end := time.Now()
	se.Summary.Duration = int64(end.Sub(start) / time.Millisecond)

	return se
}

// import (
// 	"cyeam/structs"
// 	"strings"
// 	"time"

// 	"cyeam/Godeps/_workspace/src/github.com/huichen/wukong/engine"
// 	"cyeam/Godeps/_workspace/src/github.com/huichen/wukong/types"
// 	"cyeam/Godeps/_workspace/src/github.com/mnhkahn/maodou/dao"
// 	"cyeam/Godeps/_workspace/src/github.com/mnhkahn/maodou/models"
// )

// var (
// 	// searcher是协程安全的
// 	searcher  = engine.Engine{}
// 	cacheDocs []*models.Result

// 	d dao.DaoContainer
// )

// func Index() {
// 	// 加载Result到内存
// 	var err error
// 	d, err = dao.NewDao("sqlite", "./data.db")
// 	if err != nil {
// 		panic(err)
// 	}
// 	cacheDocs, err = d.GetResults()
// 	if err != nil {
// 		panic(err)
// 	}

// 	searcher.Init(types.EngineInitOptions{
// 		SegmenterDictionaries: "./dictionary.txt"})
// 	defer searcher.Close()

// 	// 将文档加入索引
// 	for _, r := range cacheDocs {
// 		searcher.IndexDocument(r.Id, types.DocumentIndexData{Content: strings.Join([]string{r.Title, r.Description, r.Detail}, "@")})
// 	}
// 	// searcher.IndexDocument(0, types.DocumentIndexData{Content: "此次百度收购将成中国互联网最大并购"})
// 	// searcher.IndexDocument(1, types.DocumentIndexData{Content: "百度宣布拟全资收购91无线业务"})
// 	// searcher.IndexDocument(2, types.DocumentIndexData{Content: "百度是中国最大的搜索引擎"})

// 	// 等待索引刷新完毕
// 	searcher.FlushIndex()
// }

// func Search(t string) *structs.SearchResult {
// 	start := time.Now()

// 	s := searcher.Search(types.SearchRequest{Text: t})

// 	se := structs.NewSearchResult()
// 	se.Summary.Duration = time.Now().Sub(start).Nanoseconds()
// 	se.Summary.NumDocs = s.NumDocs
// 	se.Docs = make([]*structs.Doc, 0, s.NumDocs)
// 	for _, doc := range s.Docs {
// 		d := new(structs.Doc)
// 		cache := cacheDocs[doc.DocId]
// 		d.Title = cache.Title
// 		d.Link = cache.Link
// 		d.Des = cache.Description
// 		d.Figure = cache.Figure
// 		se.Docs = append(se.Docs, d)
// 	}

// 	return se
// }