// Package controllers
package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/hexops/gotextdiff"
	"github.com/hexops/gotextdiff/myers"
	"github.com/hexops/gotextdiff/span"
	curl_to_go "github.com/mnhkahn/curl-to-go"
	"github.com/mnhkahn/gogogo/app"
)

func OnlineToolHome(c *app.Context) error {
	c.HTML([]string{"./views/onlinetool.html", "./views/onlinetoolheader.html", "./views/onlinetooltail.html"}, nil)
	return nil
}

func JsonToGoStruct(c *app.Context) error {
	c.HTML([]string{"./views/jsontogostruct.html", "./views/onlinetoolheader.html", "./views/onlinetooltail.html"}, nil)
	return nil
}

func XMLToGoStruct(c *app.Context) error {
	c.HTML([]string{"./views/xmltogostruct.html", "./views/onlinetoolheader.html", "./views/onlinetooltail.html"}, nil)
	return nil
}

func JsonToThriftStruct(c *app.Context) error {
	c.HTML([]string{"./views/jsontothriftstruct.html", "./views/onlinetoolheader.html", "./views/onlinetooltail.html"}, nil)
	return nil
}

func DMLToGoStruct(c *app.Context) error {
	c.HTML([]string{"./views/dmltogostruct.html", "./views/onlinetoolheader.html", "./views/onlinetooltail.html"}, nil)
	return nil
}

func FormatJson(c *app.Context) error {
	c.HTML([]string{"./views/formatjson.html", "./views/onlinetoolheader.html", "./views/onlinetooltail.html"}, nil)
	return nil
}

func FormatXML(c *app.Context) error {
	c.HTML([]string{"./views/formatxml.html", "./views/onlinetoolheader.html", "./views/onlinetooltail.html"}, nil)
	return nil
}

func UrlEscape(c *app.Context) error {
	c.HTML([]string{"./views/urlescape.html", "./views/onlinetoolheader.html", "./views/onlinetooltail.html"}, nil)
	return nil
}

func UrlUnEscape(c *app.Context) error {
	c.HTML([]string{"./views/urlunescape.html", "./views/onlinetoolheader.html", "./views/onlinetooltail.html"}, nil)
	return nil
}

func Base32(c *app.Context) error {
	c.HTML([]string{"./views/base32.html", "./views/onlinetoolheader.html", "./views/onlinetooltail.html"}, nil)
	return nil
}

func Base32Decode(c *app.Context) error {
	c.HTML([]string{"./views/base32decode.html", "./views/onlinetoolheader.html", "./views/onlinetooltail.html"}, nil)
	return nil
}

func Base64(c *app.Context) error {
	c.HTML([]string{"./views/base64.html", "./views/onlinetoolheader.html", "./views/onlinetooltail.html"}, nil)
	return nil
}

func Base64Decode(c *app.Context) error {
	c.HTML([]string{"./views/base64decode.html", "./views/onlinetoolheader.html", "./views/onlinetooltail.html"}, nil)
	return nil
}

func Hex(c *app.Context) error {
	c.HTML([]string{"./views/hex.html", "./views/onlinetoolheader.html", "./views/onlinetooltail.html"}, nil)
	return nil
}

func HexDecode(c *app.Context) error {
	c.HTML([]string{"./views/hexdecode.html", "./views/onlinetoolheader.html", "./views/onlinetooltail.html"}, nil)
	return nil
}

func MsgPackToJson(c *app.Context) error {
	c.HTML([]string{"./views/msgpacktojson.html", "./views/onlinetoolheader.html", "./views/onlinetooltail.html"}, nil)
	return nil
}

func JsonToMsgPack(c *app.Context) error {
	c.HTML([]string{"./views/jsontomsgpack.html", "./views/onlinetoolheader.html", "./views/onlinetooltail.html"}, nil)
	return nil
}

func JsonPack(c *app.Context) error {
	c.HTML([]string{"./views/jsonpack.html", "./views/onlinetoolheader.html", "./views/onlinetooltail.html"}, nil)
	return nil
}

func Timestamp(c *app.Context) error {
	cook, _ := c.Request.Cookie("zone")
	zone := "8"
	if cook != nil {
		zone = cook.Value
	}
	n := time.Now()
	c.HTML([]string{"./views/timestamp.html", "./views/onlinetoolheader.html", "./views/onlinetooltail.html"},
		map[string]interface{}{
			"now":  n.Unix(),
			"zone": zone,
			// "res":  n.Format(time.DateTime),
		})
	return nil
}

func TimestampExec(c *app.Context) error {
	m := make(map[string]interface{})
	if err := json.NewDecoder(c.Request.Body).Decode(&m); err != nil {
		c.Request.Body.Close()
		return err
	}
	var loc *time.Location
	var err error

	z := m["zone"].(string)
	timestamp := m["timestamp"].(string)
	t, _ := strconv.ParseInt(timestamp, 10, 64)
	zone, _ := strconv.ParseInt(z, 10, 64)

	res := make(map[string]interface{}, 1)
	if zone == 0 {
		loc = time.UTC
	} else if zone == 8 {
		loc, err = time.LoadLocation("Asia/Shanghai")
		if err != nil {
			return err
		}
	}

	if len(timestamp) == 13 {
		res["data"] = time.Unix(int64(t)/1000, int64(t)-(int64(t)/1000)*1000).In(loc).Format(time.DateTime)
	} else {
		res["data"] = time.Unix(int64(t), 0).In(loc).Format(time.DateTime)
	}

	cookie := http.Cookie{
		Name:     "zone",
		Value:    z,
		HttpOnly: true,
		Path:     "/",
	}
	http.SetCookie(c.ResponseWriter, &cookie)
	c.JSON(res)
	return nil
}

func Diff(c *app.Context) error {
	c.HTML([]string{"./views/diff.html", "./views/onlinetoolheader.html", "./views/onlinetooltail.html"},
		map[string]interface{}{})
	return nil
}

func DiffExec(c *app.Context) error {
	m := make(map[string]interface{})
	if err := json.NewDecoder(c.Request.Body).Decode(&m); err != nil {
		c.Request.Body.Close()
		return err
	}
	a := m["a"].(string)
	b := m["b"].(string)

	edits := myers.ComputeEdits(span.URIFromPath("a.txt"), a, b)
	diff := fmt.Sprint(gotextdiff.ToUnified("a.txt", "a.txt", a, edits))
	log.Println(diff)
	d, err := c.ResponseWriter.Write([]byte(diff))
	log.Println(d)
	return err
}

func Json2DDL(c *app.Context) error {
	c.HTML([]string{"./views/json2ddl.html", "./views/onlinetoolheader.html", "./views/onlinetooltail.html"},
		map[string]interface{}{})
	return nil
}

func Curl2Go(c *app.Context) error {
	c.HTML([]string{"./views/curl2go.html", "./views/onlinetoolheader.html", "./views/onlinetooltail.html"},
		map[string]interface{}{})
	return nil
}

func Curl2GoExec(c *app.Context) error {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		return err
	}
	result := curl_to_go.Parse(string(body))
	c.WriteString(result)
	return nil
}

// func_to_handler.NewFuncToHandler(func(data string) string {
// 	result := curl_to_go.Parse(data)
// 	return result
// }
