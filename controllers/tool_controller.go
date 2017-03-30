/*
 * @Author: lichao115
 * @Date: 2016-12-16 12:03:54
 * @Last Modified by: mnhkahn <lichao@cyeam.com>
 * @Last Modified time: 2016-12-17 17:00:27
 */
package controllers

import (
	gohttp "net/http"
	"strings"

	"cyeam/Godeps/_workspace/src/github.com/mnhkahn/resume"
	"cyeam/Godeps/_workspace/src/github.com/mnhkahn/resume/structs"

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

var (
	DEFAULT_RESUME_PARAMS = &structs.Params{
		TouTiao:       "148931",
		TouTiaoLimit:  10,
		GitHub:        "mnhkahn",
		RepoLimit:     10,
		Weixin:        "360924857",
		StackOverflow: "1924657",
	}
)

func (this *ToolController) Resume() {
	var params *structs.Params
	if strings.Index(this.Ctx.Req.Url.RawPath, "?") == -1 {
		params = DEFAULT_RESUME_PARAMS
	} else {
		params = new(structs.Params)
		params.TouTiao = this.GetString("toutiao")
		params.TouTiaoLimit = this.GetInt("toutiaocnt")

		params.Output = this.GetString("o")
		params.GitHub = this.GetString("github")
		params.RepoLimit = this.GetInt("githubcnt")

		params.Weixin = this.GetString("weixin")

		params.StackOverflow = this.GetString("stackoverflow")
	}

	body, err := resume.Resume(params)
	if err != nil {
		this.ServeRaw([]byte(err.Error()))
		return
	}

	this.ServeRaw(body)
}

func (this *ToolController) Robots() {
	this.Ctx.Resp.Headers[http.HTTP_HEAD_CONTENTTYPE] = nil
	this.ServeFile("robots.txt")
}

func (this *ToolController) Sitemap() {
	this.ServeView("sitemap.xml")
}

func (this *ToolController) Feed() {
	this.Ctx.Resp.StatusCode = http.StatusFound
	this.Ctx.Resp.Headers.Add(http.HTTP_HEAD_LOCATION, "http://blog.cyeam.com/rss.xml")
}

func (this *ToolController) SSLVerify() {
	this.ServeView("fileauth.htm")
}

func (this *ToolController) GoogleVerify() {
	this.ServeView("google97ec3a9b69e1f4db.html")
}

func (this *ToolController) Mail() {
	this.ServeView("mail.html")
}

func (this *ToolController) Toutiao() {
	this.ServeView("toutiao.html")
}
