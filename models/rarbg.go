package models

import (
	"cyeam/structs"
	"net/url"
	"time"

	"github.com/mnhkahn/gogogo/logger"
	"github.com/mnhkahn/gogogo/util"
	"github.com/patrickmn/go-cache"
)

var (
	DoubanCache = cache.New(365*24*time.Hour, 30*time.Second)
)

func Douban(name string) *structs.DoubanResult {
	return douban(name)
}

// http://api.douban.com/v2/movie/search?q=Vikings
func douban(name string) *structs.DoubanResult {
	name = url.QueryEscape(name)

	if v, exists := DoubanCache.Get(name); exists {
		return v.(*structs.DoubanResult)
	}

	res := new(structs.DoubanResult)

	u := "http://api.douban.com/v2/movie/search?q=" + name
	err := util.HttpJson("GET", u, "", nil, res)
	if err != nil {
		logger.Warn(err)
		return res
	}

	DoubanCache.Set(name, res, 7*24*time.Hour)

	return res
}
