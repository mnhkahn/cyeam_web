/*
 * @Author: lichao115
 * @Date: 2016-12-16 12:03:54
 * @Last Modified by: mnhkahn <lichao@cyeam.com>
 * @Last Modified time: 2016-12-17 17:00:27
 */
package controllers

import (
	gohttp "net/http"

	"cyeam/Godeps/_workspace/src/github.com/mnhkahn/asciiimg"
	"cyeam/Godeps/_workspace/src/github.com/mnhkahn/cygo/net/http"
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

func (this *ToolController) Resume() {
	this.ServeView("resume.html")
}

func (this *ToolController) Robots() {
	this.Ctx.Resp.Headers[http.HTTP_HEAD_CONTENTTYPE] = nil
	this.ServeFile("robots.txt")
}

func (this *ToolController) Sitemap() {
	this.Ctx.Resp.StatusCode = http.StatusFound
	this.Ctx.Resp.Headers.Add(http.HTTP_HEAD_LOCATION, "http://blog.cyeam.com/sitemap.xml")
}

func (this *ToolController) Feed() {
	this.Ctx.Resp.StatusCode = http.StatusFound
	this.Ctx.Resp.Headers.Add(http.HTTP_HEAD_LOCATION, "http://blog.cyeam.com/rss.xml")
}

func (this *ToolController) SSLVerify() {
	this.ServeView("fileauth.htm")
}
