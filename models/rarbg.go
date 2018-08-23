package models

import (
	"cyeam/structs"
	"log"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego/httplib"
	cache "github.com/patrickmn/go-cache"
)

var (
	Resolution = map[string]bool{"1080p": true, "DVDScr": true, "720p": true}
)

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

	DoubanCache.Set(name, res, 7*24*time.Hour)
	// DoubanCache.SetDefault(name, res)

	return res
}
