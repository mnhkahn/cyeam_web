package controllers

import (
	"cyeam/models"
	"cyeam/search"
	"cyeam/service"
	"cyeam/structs"
	"net/http"

	"github.com/mnhkahn/gogogo/app"
	"github.com/mnhkahn/gogogo/logger"
	"github.com/mnhkahn/peanut/index"
)

func Index(c *app.Context) error {
	if c.Request.Host == "mail.cyeam.com" {
		c.HTML([]string{"mail.html"}, nil)
		return nil
	}
	peanut := search.Peanut("*", 1, 4, "PubDate", index.DESC)
	for _, p := range peanut.Docs {
		if p.Figure == "" {
			p.Figure = "http://cyeam.qiniudn.com/gopherswrench.jpg"
		}
	}
	c.HTML([]string{"./views/index.html", "./views/head.html", "./views/tail.html"}, peanut)
	return nil
}

func Search(c *app.Context) error {
	t := c.GetString("t")
	if len(t) == 0 {
		c.JSON(new(structs.SearchResult))
		return nil
	}
	page, _ := c.GetInt("page")
	if page <= 0 || page >= 100 {
		page = 1
	}
	size, _ := c.GetInt("ps")
	if size <= 0 || size >= 100 {
		size = 20
	}
	sortField := c.GetString("sort", "PV")
	asc, _ := c.GetBool("sort_asc")

	c.JSON(search.Peanut(t, page, size, sortField, asc))
	return nil
}

// https://startbootstrap.com/template-overviews/blog-home/
func SearchView(c *app.Context) error {
	t := c.GetString("t")
	if len(t) == 0 {
		// 302
		http.Redirect(c.ResponseWriter, c.Request, "/", http.StatusFound)
		return nil
	}

	page, _ := c.GetInt("page")
	if page <= 0 || page >= 100 {
		page = 1
	}
	size, _ := c.GetInt("ps")
	if size <= 0 || size >= 100 {
		size = 20
	}
	sortField := c.GetString("sort", "PV")
	asc, _ := c.GetBool("sort_asc")

	c.HTML([]string{"./views/search.html", "./views/head.html", "./views/tail.html"}, search.Peanut(t, page, size, sortField, asc))
	return nil
}

func Bing(c *app.Context) error {
	http.Redirect(c.ResponseWriter, c.Request, models.GetBing(), http.StatusFound)
	return nil
}

func BinCalc(c *app.Context) error {
	castruct := new(structs.CalcStruct)
	err := c.Request.ParseForm()
	if err != nil {
		return err
	}
	castruct.Dec = service.BinToDex(c.Request.FormValue("bin"))
	c.HTML([]string{"./views/calc.html", "./views/head.html", "./views/tail.html"}, castruct)

	logger.Info("bincalc", castruct.Bin, castruct.Dec)

	return nil
}

func JDVerify(c *app.Context) error {
	c.HTML([]string{"./static/jos_guid.txt"}, nil)
	return nil
}

func DoubanMovie(c *app.Context) error {
	name := c.GetString("s")
	c.JSON(models.Douban(name))
	return nil
}

func NotFound(c *app.Context) error {
	c.ResponseWriter.WriteHeader(http.StatusNotFound)
	c.HTML([]string{"./views/404.html", "./views/head.html", "./views/tail.html"}, nil)
	return nil
}
