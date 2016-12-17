package main

import (
	"cyeam/controllers"
	"cyeam/search"
	"os"

	"cyeam/Godeps/_workspace/src/github.com/mnhkahn/cygo/net/http"
)

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

	http.Router("/", "GET", &controllers.MainController{}, "Get")
	http.Router("/s", "GET", &controllers.MainController{}, "Search")
	http.Router("/t", "GET", &controllers.MainController{}, "SearchView")
	http.Router("/bing", "GET", &controllers.MainController{}, "Bing")
	http.Router("/bincalc", "GET", &controllers.MainController{}, "BinCalc")
	http.Router("/bincalc", "POST", &controllers.MainController{}, "BinCalc")

	http.Router("/weixin", "GET", &controllers.WeixinController{}, "Verify")
	http.Router("/weixin", "POST", &controllers.WeixinController{}, "WeixinMsg")

	http.Router("/ascii", "GET", &controllers.ToolController{}, "Ascii")

	http.Router("/robots.txt", "GET", &controllers.ToolController{}, "Robots")
	http.Router("/sitemap.xml", "GET", &controllers.ToolController{}, "Sitemap")
	http.Router("/feed/", "GET", &controllers.ToolController{}, "Feed")
	http.Router("/resume", "GET", &controllers.ToolController{}, "Resume")
}
