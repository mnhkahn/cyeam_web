package main

import (
	"cyeam/controllers"
	"cyeam/models"
	"cyeam/search"
	"cyeam/services"
	"cyeam/structs"
	"fmt"
	"os"

	"cyeam/Godeps/_workspace/src/github.com/mnhkahn/cygo/net/http"
)

type MainController struct {
	http.Controller
}

func (this *MainController) Get() {
	//	this.Ctx.Resp.Body = DEFAULT_HTML
	this.ServeView("index.html")
}

func (this *MainController) Search() {
	t := this.GetString("t")
	this.ServeJson(search.Search(t))
}

func (this *MainController) SearchView() {
	t := this.GetString("t")
	this.ServeView("search.html", search.Search(t))
}

func (this *MainController) Bing() {
	this.Ctx.Resp.StatusCode = http.StatusFound
	this.Ctx.Resp.Headers.Add(http.HTTP_HEAD_LOCATION, models.GetBing())
}

func (this *MainController) BinCalc() {
	castruct := new(structs.CalcStruct)
	err := this.ParseForms(castruct)
	if err != nil {
		this.Ctx.Resp.Body = err.Error()
		return
	}
	castruct.Dec = services.BinToDex(castruct.Bin)
	fmt.Println(castruct.Bin, castruct.Dec, "--------------------")
	this.ServeView("calc.html", castruct)
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "5000"
	}
	http.Serve(":" + port)
}

func init() {
	if err := search.InitSwiftype(); err != nil {
		panic(err)
	}

	http.Router("/", "GET", &MainController{}, "Get")
	http.Router("/s", "GET", &MainController{}, "Search")
	http.Router("/t", "GET", &MainController{}, "SearchView")
	http.Router("/bing", "GET", &MainController{}, "Bing")
	http.Router("/bincalc", "GET", &MainController{}, "BinCalc")
	http.Router("/bincalc", "POST", &MainController{}, "BinCalc")

	http.Router("/weixin", "GET", &controllers.WeixinController{}, "Verify")
	http.Router("/weixin", "POST", &controllers.WeixinController{}, "WeixinMsg")

	http.Router("/ascii", "GET", &controllers.ToolController{}, "Ascii")

	http.Router("/robots.txt", "GET", &controllers.ToolController{}, "Robots")
	http.Router("/sitemap.xml", "GET", &controllers.ToolController{}, "Sitemap")
	http.Router("/feed/", "GET", &controllers.ToolController{}, "Feed")
}
