/*
 * @Author: lichao115
 * @Date: 2016-12-15 17:22:29
 * @Last Modified by: lichao115
 * @Last Modified time: 2016-12-15 17:26:22
 */
package controllers

import (
	"crypto/sha1"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"sort"
	"strings"
	"time"

	"cyeam/models"

	"cyeam/Godeps/_workspace/src/github.com/mnhkahn/cygo/net/http"
)

type WeixinController struct {
	http.Controller
}

const (
	TOKEN = "cyeam"

	Text     = "text"
	Location = "location"
	Image    = "image"
	Link     = "link"
	Event    = "event"
	Music    = "music"
	News     = "news"
)

func (this *WeixinController) Verify() {
	signature := this.GetString("signature")
	timestamp := this.GetString("timestamp")
	nonce := this.GetString("nonce")
	echostr := this.GetString("echostr")

	dict := []string{timestamp, nonce, echostr}
	sort.Strings(dict)

	h := sha1.New()
	io.WriteString(h, strings.Join(dict, ""))

	if Signature(timestamp, nonce) == signature {
		this.ServeRaw([]byte(echostr))
	} else {
		this.ServeRaw([]byte(""))
	}
}

func Signature(timestamp, nonce string) string {
	strs := sort.StringSlice{TOKEN, timestamp, nonce}
	sort.Strings(strs)
	str := ""
	for _, s := range strs {
		str += s
	}
	h := sha1.New()
	h.Write([]byte(str))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func (this *WeixinController) WeixinMsg() {
	if this.Ctx.Req.Body == "" {
		this.Ctx.Resp.StatusCode = http.StatusNotFound
		return
	}
	log.Println(this.Ctx.Req.Body, "FFFFFFFFFFFFF")

	var wreq *models.Request
	var err error
	if wreq, err = DecodeRequest([]byte(this.Ctx.Req.Body)); err != nil {
		this.Ctx.Resp.StatusCode = http.StatusNotFound
		return
	}

	wresp, err := dealwith(wreq, this.Ctx.Req)
	if err != nil {
		this.Ctx.Resp.StatusCode = http.StatusNotFound
		return
	}
	data, err := wresp.Encode()
	if err != nil {
		this.Ctx.Resp.StatusCode = http.StatusNotFound
		return
	}
	this.ServeRaw(data)
	return
}

func DecodeRequest(data []byte) (req *models.Request, err error) {
	req = &models.Request{}
	if err = xml.Unmarshal(data, req); err != nil {
		return
	}
	req.CreateTime *= time.Second
	return
}

func NewResponse() (resp *models.Response) {
	resp = &models.Response{}
	resp.CreateTime = time.Duration(time.Now().Unix())
	return
}

func dealwith(req *models.Request, r *http.Request) (resp *models.Response, err error) {
	resp = NewResponse()
	resp.ToUserName = req.FromUserName
	resp.FromUserName = req.ToUserName
	switch req.MsgType {
	case Text:
		resp = handleText(req, resp)
	case Image:
		resp = handleImage(req, resp)
	case Location:
		resp = handleLocation(req, resp)
	default:
		resp = handleText(req, resp)
	}

	return resp, nil
}

func handleLocation(req *models.Request, resp *models.Response) *models.Response {
	resp.MsgType = Text
	w := models.NewWeather(fmt.Sprintf("%v", req.Location_X), fmt.Sprintf("%v", req.Location_Y))
	resp.Content = fmt.Sprintf("%s，位置：%s，温度：%d，天气：%s", w.Summary, req.Label, w.Temp, w.Skycon)
	return resp
}

func handleImage(req *models.Request, resp *models.Response) *models.Response {
	resp.MsgType = News
	resp.Content = "Ascii"
	resp.ArticleCount = 1
	a := models.Item{}
	a.Title = "我的图片的ASCII码"
	a.PicUrl = req.PicUrl
	a.Description = "点击『查看原文』来查看详细说明"
	a.Url = "http://cyeam.com/ascii?url=" + strings.Trim(req.PicUrl, "http://")
	resp.FuncFlag = 1
	resp.Articles = append(resp.Articles, &a)

	return resp
}

func handleText(req *models.Request, resp *models.Response) *models.Response {
	resp.MsgType = Text

	resp.Content = req.Content
	if req.Content == "doodle" {
		doodle := models.GetDoodle()

		if doodle.Doodle != "" {
			resp.MsgType = News
			resp.Content = "doodle"
			resp.ArticleCount = 1

			a := models.Item{}
			a.Title = doodle.Title
			a.PicUrl = doodle.Doodle
			a.Description = "点击『查看原文』来查看详细说明"
			a.Url = "http://cyeam.com/"
			resp.FuncFlag = 1
			resp.Articles = append(resp.Articles, &a)
		} else {
			resp.Content = ""
		}
	} else if req.Content == "bing" {
		bing := models.GetBing()
		resp.MsgType = News
		resp.Content = "bing"
		resp.ArticleCount = 1

		a := models.Item{}
		a.Title = "bing"
		a.PicUrl = bing
		a.Description = "点击『查看原文』来查看详细说明"
		a.Url = "http://cyeam.com/"
		resp.FuncFlag = 1
		resp.Articles = append(resp.Articles, &a)
	}
	return resp
}
