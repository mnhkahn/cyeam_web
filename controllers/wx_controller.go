package controllers

import (
	"bytes"
	"crypto/sha1"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"sort"
	"strings"
	"time"

	"cyeam/models"
	"cyeam/search"

	"io/ioutil"

	"net/http"

	"github.com/mnhkahn/gogogo/app"
	"github.com/mnhkahn/gogogo/logger"
	"github.com/mnhkahn/swiftype"
)

const (
	TOKEN = "cyeam"

	Text     = "text"
	Location = "location"
	Image    = "image"
	Link     = "link"
	Event    = "event"
	Music    = "music"
	News     = "news"

	HELP = "文本框里面回复内容，可以搜索以往历史文章。\n发送图片，可以生成一张ASCII编码的图片。\n发送地址，可以查看当前地址天气。"
)

func Weixin(c *app.Context) error {
	if c.Request.Method == "GET" { // verify
		signature := c.GetString("signature")
		timestamp := c.GetString("timestamp")
		nonce := c.GetString("nonce")
		echostr := c.GetString("echostr")

		dict := []string{timestamp, nonce, echostr}
		sort.Strings(dict)

		h := sha1.New()
		io.WriteString(h, strings.Join(dict, ""))

		if Signature(timestamp, nonce) == signature {
			c.WriteBytes([]byte(echostr))
		} else {
			c.WriteBytes([]byte(""))
		}
		return nil
	} else { // msg
		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			return err
		}

		logger.Info(string(body))

		var wreq *models.Request
		if wreq, err = DecodeRequest(body); err != nil {
			return err
		}

		wresp, err := dealwith(wreq, c.Request)
		if err != nil {
			return err
		}
		data, err := wresp.Encode()
		if err != nil {
			return err
		}

		c.WriteBytes(data)
		return nil
	}

	return nil
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

	log.Println("REQ", req.FromUserName.Data, req.ToUserName.Data, req.MsgType.Data, req.CreateTime, req.MsgId)

	switch req.MsgType.Data {
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
	resp.MsgType.Data = Text
	w := models.NewWeather(fmt.Sprintf("%v", req.Location_X), fmt.Sprintf("%v", req.Location_Y))
	resp.Content.Data = fmt.Sprintf("%s，位置：%s，温度：%d，天气：%s", w.Summary, req.Label, w.Temp, w.Skycon)
	return resp
}

func handleImage(req *models.Request, resp *models.Response) *models.Response {
	resp.MsgType.Data = News
	resp.Content.Data = "Ascii"
	resp.ArticleCount = 1
	a := models.Item{}
	a.Title.Data = "我的图片的ASCII码"
	a.PicUrl = req.PicUrl
	a.Description.Data = "点击『查看原文』来查看详细说明"
	a.Url.Data = "http://cyeam.com/ascii?url=" + strings.Trim(req.PicUrl.Data, "http://")
	// resp.FuncFlag = 1
	resp.Articles = new(models.Articles)
	resp.Articles.Items = append(resp.Articles.Items, &a)

	return resp
}

func handleText(req *models.Request, resp *models.Response) *models.Response {
	resp.MsgType.Data = Text

	if req.Content.Data == "doodle" {
		doodle := models.GetDoodle()

		if doodle.Doodle != "" {
			resp.MsgType.Data = News
			resp.Content.Data = "doodle"
			resp.ArticleCount = 1

			a := models.Item{}
			a.Title.Data = doodle.Title
			a.PicUrl.Data = doodle.Doodle
			a.Description.Data = "点击『查看原文』来查看详细说明"
			a.Url.Data = "http://cyeam.com/"
			// resp.FuncFlag = 1
			resp.Articles = new(models.Articles)
			resp.Articles.Items = append(resp.Articles.Items, &a)
		} else {
			resp.Content.Data = ""
		}
	} else if req.Content.Data == "bing" {
		bing := models.GetBing()
		resp.MsgType.Data = News
		resp.Content.Data = "bing"
		resp.ArticleCount = 1

		a := models.Item{}
		a.Title.Data = "bing"
		a.PicUrl.Data = bing
		a.Description.Data = "点击『查看原文』来查看详细说明"
		a.Url.Data = "http://cyeam.com/"
		// resp.FuncFlag = 1
		resp.Articles = new(models.Articles)
		resp.Articles.Items = append(resp.Articles.Items, &a)
	} else {
		se := search.Search(swiftype.NewSearchParam(req.Content.Data))
		if len(se.Docs) > 0 {
			buf := bytes.NewBuffer(nil)

			buf.WriteString(fmt.Sprintf("搜索: %s 耗时：%dms\n", req.Content.Data, se.Summary.Duration))

			for i, doc := range se.Docs {
				buf.WriteString(fmt.Sprintf("%d. 《%s》 %s\n", i+1, doc.Title, doc.Link))
			}

			resp.Content.Data = buf.String()
		} else {
			resp.Content.Data = HELP
		}
	}
	return resp
}
