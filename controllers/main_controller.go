package controllers

import (
	"cyeam/models"
	"cyeam/search"
	"cyeam/services"
	"cyeam/structs"
	"net/http"

	"github.com/mnhkahn/gogogo/app"
	"github.com/mnhkahn/gogogo/logger"
)

func Index(c *app.Context) error {
	if c.Request.Host == "mail.cyeam.com" {
		c.HTML([]string{"mail.html"}, nil)
		return nil
	}
	peanut := search.Peanut("*", 1, 3)
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
	c.JSON(search.Peanut(t, page, size))
	return nil
}

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

	c.HTML([]string{"./views/search.html", "./views/head.html", "./views/tail.html"}, search.Peanut(t, page, size))
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
	castruct.Dec = services.BinToDex(c.Request.FormValue("bin"))
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
