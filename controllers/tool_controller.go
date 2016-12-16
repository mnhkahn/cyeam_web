/*
 * @Author: lichao115
 * @Date: 2016-12-16 12:03:54
 * @Last Modified by: lichao115
 * @Last Modified time: 2016-12-16 12:05:54
 */
package controllers

import (
	gohttp "net/http"

	"github.com/mnhkahn/asciiimg"
	"github.com/mnhkahn/cygo/net/http"
)

type ToolController struct {
	http.Controller
}

func (this *ToolController) Ascii() {
	url := this.GetString("url")

	res, err := gohttp.Get("http://" + url)
	if err != nil {
		this.ServeRaw([]byte(err.Error()))
		return
	}

	ai, err := asciiimg.NewAsciiImg(res.Body)
	if err != nil {
		this.ServeRaw([]byte(err.Error()))
		return
	}
	defer res.Body.Close()

	this.ServeRaw([]byte(ai.DoByCol(38)))
}
