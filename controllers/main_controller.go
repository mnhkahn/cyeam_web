package controllers

import (
	"cyeam/models"
	"cyeam/search"
	"cyeam/services"
	"cyeam/structs"
	"fmt"
	"log"

	"cyeam/Godeps/_workspace/src/github.com/mnhkahn/cygo/net/url"

	"io/ioutil"

	"cyeam/Godeps/_workspace/src/github.com/astaxie/beego/httplib"
	"cyeam/Godeps/_workspace/src/github.com/mnhkahn/cygo/net/http"
)

type MainController struct {
	http.Controller
}

func (this *MainController) Get() {
	this.ServeView("index.html")
}

func (this *MainController) Search() {
	t := this.GetString("t")
	if len(t) == 0 {
		this.ServeJson(new(structs.SearchResult))
		return
	}
	this.ServeJson(search.Search(t))
}

func (this *MainController) SearchView() {
	t := this.GetString("t")
	if len(t) == 0 {
		// 302
		this.Ctx.Resp.StatusCode = http.StatusFound
		this.Ctx.Resp.Headers.Add(http.HTTP_HEAD_LOCATION, "/")
		return
	}
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

const (
	RARBG_HOST        = "https://rarbg.to/torrents.php"
	RARGB_TORRENT_URL = "https://rarbg.to/torrent/%s"
)

var (
	DEFAULT_RARBG_QUERY = url.ParseQuery("category=14;17;42;44;45;46;47;48&order=seeders&by=DESC")
)

func (this *MainController) Rarbg() {
	if len(this.Ctx.Req.Url.Query()) == 0 {
		this.Ctx.Resp.StatusCode = http.StatusFound
		this.Ctx.Resp.Headers.Add(http.HTTP_HEAD_LOCATION, "/rarbg?category=14;17;42;44;45;46;47;48&order=seeders&by=DESC")
		return
	}

	query := this.Ctx.Req.Url.Query()
	for k, v := range DEFAULT_RARBG_QUERY {
		for _, vv := range v {
			query.Add(k, vv)
		}
	}

	u := RARBG_HOST + "?" + query.String()
	log.Println(u)

	req := httplib.Get(u)
	for k, v := range this.Ctx.Req.Headers {
		for _, vv := range v {
			req = req.Header(k, vv)
		}
	}
	resp, err := req.DoRequest()
	a, _ := ioutil.ReadAll(resp.Body)
	log.Println(string(a))
	// body, err := os.Open("rarbgtest.html") // For read access.
	if err != nil {
		this.ServeRaw([]byte(err.Error()))
		return
	}
	// res, err := models.Rarbg(body)
	res, err := models.Rarbg(resp.Body)
	if err != nil {
		this.ServeRaw([]byte(err.Error()))
		return
	}

	this.ServeView("rarbg.html", res)
}

func (this *MainController) Torrents() {
	id := this.GetString("id")
	u := fmt.Sprintf(RARGB_TORRENT_URL, id)

	req := httplib.Get(u)
	resp, err := req.DoRequest()
	// body, err := os.Open("rarbgtor.html") // For read access.
	if err != nil {
		this.ServeRaw([]byte(err.Error()))
		return
	}
	// res, err := models.RarbgTorrent(body)
	res, err := models.RarbgTorrent(resp.Body)
	if err != nil {
		this.ServeRaw([]byte(err.Error()))
		return
	}

	this.Ctx.Resp.StatusCode = http.StatusFound
	this.Ctx.Resp.Headers.Add(http.HTTP_HEAD_LOCATION, res)
}
