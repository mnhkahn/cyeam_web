package main

import (
	"bytes"
	"cyeam/controllers"
	"cyeam/util"
	"encoding/base32"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/ChimeraCoder/gojson"
	"github.com/miku/zek"
	"github.com/mnhkahn/gogogo/app"
	"github.com/mnhkahn/gogogo/app/handler/func_to_handler"
	"github.com/mnhkahn/gogogo/logger"
	"github.com/mnhkahn/pkg/xhex"
	"github.com/mnhkahn/togo/dmltogo"
	"github.com/vmihailenco/msgpack"
	"github.com/yosssi/gohtml"
	"go/format"
	"html"
	"net"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "5000"
	}

	str, _ := os.Getwd()
	logger.Info(str)
	err := filepath.Walk("/go/src/app/", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}
		fmt.Printf("dir: %v: name: %s\n", info.IsDir(), path)
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}

	l, err := net.Listen("tcp", ":"+port)
	if err != nil {
		logger.Errorf("Listen: %v", err)
		return
	}
	app.Serve(l)
}

func init() {
	app.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	app.Handle("/", &app.Got{controllers.Index})
	app.Handle("/bing", &app.Got{controllers.Bing})
	app.Handle("/s", &app.Got{controllers.Search})
	app.Handle("/t", &app.Got{controllers.SearchView})
	app.Handle("/bincalc", &app.Got{controllers.BinCalc})
	app.Handle("/weixin", &app.Got{controllers.Weixin})
	app.Handle("/jos_guid.txt", &app.Got{controllers.JDVerify})
	app.Handle("/douban/movie", &app.Got{controllers.DoubanMovie})

	app.Handle("/favicon.ico", &app.Got{controllers.Favicon})
	app.Handle("/ads.txt", &app.Got{controllers.Ads})
	app.Handle("/toolbox", &app.Got{controllers.ToolBox})
	app.Handle("/ascii", &app.Got{controllers.Ascii})
	app.Handle("/robots.txt", &app.Got{controllers.Robots})
	app.Handle("/sitemap.xml", &app.Got{controllers.Sitemap})
	app.Handle("/feed/", &app.Got{controllers.Feed})
	app.Handle("/resume", &app.Got{controllers.Resume})
	app.Handle("/geek", &app.Got{controllers.Toutiao})
	app.Handle("/neitui", &app.Got{controllers.Neitui})

	app.Handle("/.well-known/pki-validation/fileauth.htm", &app.Got{controllers.SSLVerify})
	app.Handle("/google97ec3a9b69e1f4db.html", &app.Got{controllers.GoogleVerify})

	app.Handle("/tool", &app.Got{controllers.FormatJson})
	app.Handle("/tool/json2gostruct", &app.Got{controllers.JsonToGoStruct})
	app.Handle("/tool/xml2gostruct", &app.Got{controllers.XMLToGoStruct})
	app.Handle("/tool/json2thriftstruct", &app.Got{controllers.JsonToThriftStruct})
	app.Handle("/tool/dml2gostruct", &app.Got{controllers.DMLToGoStruct})
	app.Handle("/tool/formatjson", &app.Got{controllers.FormatJson})
	app.Handle("/tool/formatxml", &app.Got{controllers.FormatXML})
	app.Handle("/tool/urlescape", &app.Got{controllers.UrlEscape})
	app.Handle("/tool/urlunescape", &app.Got{controllers.UrlUnEscape})
	app.Handle("/tool/base32", &app.Got{controllers.Base32})
	app.Handle("/tool/base32decode", &app.Got{controllers.Base32Decode})
	app.Handle("/tool/base64", &app.Got{controllers.Base64})
	app.Handle("/tool/base64decode", &app.Got{controllers.Base64Decode})
	app.Handle("/tool/hex", &app.Got{controllers.Hex})
	app.Handle("/tool/hexdecode", &app.Got{controllers.HexDecode})
	app.Handle("/tool/ascii", &app.Got{controllers.Hex})
	app.Handle("/tool/msgpacktojson", &app.Got{controllers.MsgPackToJson})
	app.Handle("/tool/jsonpack", &app.Got{controllers.JsonPack})
	app.Handle("/tool/json2gostruct/exec", func_to_handler.NewFuncToHandler(func(data string) (string, error) {
		var parser gojson.Parser = gojson.ParseJson
		if output, err := gojson.Generate(bytes.NewBufferString(data), parser, "Foo", "dto", []string{"json"}, true, true); err != nil {
			return "", err
		} else {
			return string(output), nil
		}
	}))
	app.Handle("/tool/xml2gostruct/exec", func_to_handler.NewFuncToHandler(func(data string) (string, error) {
		root := new(zek.Node)
		if _, err := root.ReadFrom(strings.NewReader(data)); err != nil {
			return err.Error(), err
		}
		var buf bytes.Buffer
		sw := zek.NewStructWriter(&buf)
		if err := sw.WriteNode(root); err != nil {
			return err.Error(), err
		}

		var out []byte
		var err error
		if out, err = format.Source(buf.Bytes()); err != nil {
			return err.Error(), err

		}
		return string(out), nil
	}))
	app.Handle("/tool/dml2gostruct/exec", func_to_handler.NewFuncToHandler(dmltogo.DmlToGo))
	app.Handle("/tool/jsontothriftstruct/exec", func_to_handler.NewFuncToHandler(util.JsonToThrift))
	app.Handle("/tool/formatjson/exec", func_to_handler.NewFuncToHandler(func(data string) ([]byte, error) {
		var out bytes.Buffer
		err := json.Indent(&out, []byte(data), "", "    ")
		if err != nil {
			return nil, err
		}
		return out.Bytes(), nil
	}))
	app.Handle("/tool/formatxml/exec", func_to_handler.NewFuncToHandler(func(data string) ([]byte, error) {
		x := html.EscapeString(gohtml.FormatWithLineNo(data))
		return []byte(x), nil
	}))
	app.Handle("/tool/urlescape/exec", func_to_handler.NewFuncToHandler(url.QueryEscape))
	app.Handle("/tool/urlunescape/exec", func_to_handler.NewFuncToHandler(url.QueryUnescape))
	app.Handle("/tool/base32/exec", func_to_handler.NewFuncToHandler(func(data string) string {
		return base32.StdEncoding.EncodeToString([]byte(data))
	}))
	app.Handle("/tool/base32decode/exec", func_to_handler.NewFuncToHandler(base32.StdEncoding.DecodeString))
	app.Handle("/tool/base64/exec", func_to_handler.NewFuncToHandler(func(data string) string {
		return base64.StdEncoding.EncodeToString([]byte(data))
	}))
	app.Handle("/tool/base64decode/exec", func_to_handler.NewFuncToHandler(base64.StdEncoding.DecodeString))
	app.Handle("/tool/hex/exec", func_to_handler.NewFuncToHandler(func(data string) string {
		return hex.EncodeToString([]byte(data))
	}))
	app.Handle("/tool/hexdecode/exec", func_to_handler.NewFuncToHandler(hex.DecodeString))
	app.Handle("/tool/msgpacktojson/exec", func_to_handler.NewFuncToHandler(func(data string) string {
		d, err := xhex.DecodeString(data)
		if err != nil {
			return err.Error()
		}

		res := make(map[string]interface{})
		err = msgpack.Unmarshal(d, &res)
		if err != nil {
			return err.Error()
		}
		buf, err := json.Marshal(res)
		if err != nil {
			return err.Error()
		}

		var out bytes.Buffer
		err = json.Indent(&out, buf, "", "    ")
		if err != nil {
			return ""
		}
		return out.String()
	}))
	app.Handle("/tool/jsonpack/exec", func_to_handler.NewFuncToHandler(func(data string) string {
		out := strings.Replace(data, " ", "", -1)
		out = strings.Replace(out, "\n", "", -1)
		return out
	}))

}
