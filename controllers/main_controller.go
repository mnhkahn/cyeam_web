package controllers

import (
	"cyeam/models"
	"cyeam/search"
	"cyeam/services"
	"cyeam/structs"
	"fmt"

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
