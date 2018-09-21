// Package controllers
package controllers

import "github.com/mnhkahn/gogogo/app"

func OnlineToolHome(c *app.Context) error {
	c.HTML([]string{"./views/onlinetool.html", "./views/onlinetoolheader.html", "./views/onlinetooltail.html"}, nil)
	return nil
}

func JsonToGoStruct(c *app.Context) error {
	c.HTML([]string{"./views/jsontogostruct.html", "./views/onlinetoolheader.html", "./views/onlinetooltail.html"}, nil)
	return nil
}

func FormatJson(c *app.Context) error {
	c.HTML([]string{"./views/formatjson.html", "./views/onlinetoolheader.html", "./views/onlinetooltail.html"}, nil)
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
