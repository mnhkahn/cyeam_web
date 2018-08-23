package main

import (
	"cyeam/controllers"
	"cyeam/search"
	"net/http"
	"os"

	"net"

	"github.com/mnhkahn/gogogo/app"
	"github.com/mnhkahn/gogogo/logger"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "5000"
	}

	l, err := net.Listen("tcp", ":"+port)
	if err != nil {
		logger.Errorf("Listen: %v", err)
		return
	}
	app.Serve(l)
}

func init() {
	if err := search.InitSwiftype(); err != nil {
		panic(err)
	}

	app.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	app.Handle("/", &app.Got{controllers.Index})
	app.Handle("/bing", &app.Got{controllers.Bing})
	app.Handle("/s", &app.Got{controllers.Search})
	app.Handle("/t", &app.Got{controllers.SearchView})
	app.Handle("/bincalc", &app.Got{controllers.BinCalc})
	app.Handle("/weixin", &app.Got{controllers.Weixin})
	app.Handle("/jos_guid.txt", &app.Got{controllers.JDVerify})
	app.Handle("/douban/movie", &app.Got{controllers.DoubanMovie})

	app.Handle("/toolbox", &app.Got{controllers.ToolBox})
	app.Handle("/ascii", &app.Got{controllers.Ascii})
	app.Handle("/robots.txt", &app.Got{controllers.Robots})
	app.Handle("/sitemap.xml", &app.Got{controllers.Sitemap})
	app.Handle("/feed/", &app.Got{controllers.Feed})
	app.Handle("/resume", &app.Got{controllers.Resume})
	app.Handle("/geek", &app.Got{controllers.Toutiao})

	app.Handle("/.well-known/pki-validation/fileauth.htm", &app.Got{controllers.SSLVerify})
	app.Handle("/google97ec3a9b69e1f4db.html", &app.Got{controllers.GoogleVerify})
}
