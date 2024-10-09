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
	"go/format"
	"html"
	"net"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/ChimeraCoder/gojson"
	"github.com/miku/zek"
	ddlmaker "github.com/mnhkahn/ddl-maker"
	"github.com/mnhkahn/gogogo/app"
	"github.com/mnhkahn/gogogo/app/handler/func_to_handler"
	"github.com/mnhkahn/gogogo/logger"
	"github.com/mnhkahn/togo/dmltogo"
	"github.com/vmihailenco/msgpack"
	"github.com/yosssi/gohtml"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "5000"
	}

	str, _ := os.Getwd()
	logger.Info(str)
	err := filepath.Walk("./views", func(path string, info os.FileInfo, err error) error {
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

	app.Handle("/", &app.Got{H: controllers.ToolBox})
	app.Handle("/bing", &app.Got{H: controllers.Bing})
	app.Handle("/s", &app.Got{controllers.Search})
	app.Handle("/t", &app.Got{controllers.SearchView})
	app.Handle("/bincalc", &app.Got{controllers.BinCalc})
	app.Handle("/weixin", &app.Got{controllers.Weixin})
	app.Handle("/jos_guid.txt", &app.Got{controllers.JDVerify})
	app.Handle("/douban/movie", &app.Got{controllers.DoubanMovie})

	app.Handle("/favicon.ico", &app.Got{controllers.Favicon})
	app.Handle("/ads.txt", &app.Got{controllers.Ads})
	app.Handle("/ascii", &app.Got{controllers.Ascii})
	app.Handle("/sitemap.xml", app.Got{H: controllers.SiteMapXML})
	app.Handle("/sitemap.txt", app.Got{H: controllers.SiteMapRaw})
	app.Handle("/robots.txt", app.Got{controllers.Robots})
	app.Handle("/feed/", &app.Got{controllers.Feed})
	app.Handle("/resume", &app.Got{controllers.Resume})
	app.Handle("/geek", &app.Got{controllers.Toutiao})
	app.Handle("/neitui", &app.Got{controllers.Neitui})

	app.Handle("/.well-known/pki-validation/fileauth.htm", &app.Got{controllers.SSLVerify})
	app.Handle("/google97ec3a9b69e1f4db.html", &app.Got{controllers.GoogleVerify})
	app.Handle("/BingSiteAuth.xml", &app.Got{controllers.BingVerify})

	controllers.HandleViews("/tool", []string{"./views/formatjson.html", "./views/onlinetoolheader.html", "./views/onlinetooltail.html"}, controllers.DataDefaultGetter)
	controllers.HandleViews("/tool/json2gostruct", []string{"./views/jsontogostruct.html", "./views/onlinetoolheader.html", "./views/onlinetooltail.html"}, controllers.DataDefaultGetter)
	controllers.HandleViews("/tool/xml2gostruct", []string{"./views/xmltogostruct.html", "./views/onlinetoolheader.html", "./views/onlinetooltail.html"}, controllers.DataDefaultGetter)
	controllers.HandleViews("/tool/json2thriftstruct", []string{"./views/jsontothriftstruct.html", "./views/onlinetoolheader.html", "./views/onlinetooltail.html"}, controllers.DataDefaultGetter)
	controllers.HandleViews("/tool/dml2gostruct", []string{"./views/dmltogostruct.html", "./views/onlinetoolheader.html", "./views/onlinetooltail.html"}, controllers.DataDefaultGetter)
	controllers.HandleViews("/tool/formatjson", []string{"./views/formatjson.html", "./views/onlinetoolheader.html", "./views/onlinetooltail.html"}, controllers.DataDefaultGetter)
	controllers.HandleViews("/tool/formatxml", []string{"./views/formatxml.html", "./views/onlinetoolheader.html", "./views/onlinetooltail.html"}, controllers.DataDefaultGetter)
	app.Handle("/tool/urlescape", &app.Got{controllers.UrlEscape})
	controllers.HandleViews("/tool/urlunescape", []string{"./views/urlunescape.html", "./views/onlinetoolheader.html", "./views/onlinetooltail.html"}, controllers.DataDefaultGetter)
	app.Handle("/tool/base32", &app.Got{controllers.Base32})
	controllers.HandleViews("/tool/base32decode", []string{"./views/base32decode.html", "./views/onlinetoolheader.html", "./views/onlinetooltail.html"}, controllers.DataDefaultGetter)
	app.Handle("/tool/base64", &app.Got{controllers.Base64})
	controllers.HandleViews("/tool/base64decode", []string{"./views/base64decode.html", "./views/onlinetoolheader.html", "./views/onlinetooltail.html"}, controllers.DataDefaultGetter)
	app.Handle("/tool/hex", &app.Got{controllers.Hex})
	controllers.HandleViews("/tool/hexdecode", []string{"./views/hexdecode.html", "./views/onlinetoolheader.html", "./views/onlinetooltail.html"}, controllers.DataDefaultGetter)
	app.Handle("/tool/jsontomsgpack", &app.Got{controllers.JsonToMsgPack})
	controllers.HandleViews("/tool/msgpacktojson", []string{"./views/msgpacktojson.html", "./views/onlinetoolheader.html", "./views/onlinetooltail.html"}, controllers.DataDefaultGetter)
	controllers.HandleViews("/tool/jsonpack", []string{"./views/jsonpack.html", "./views/onlinetoolheader.html", "./views/onlinetooltail.html"}, controllers.DataDefaultGetter)
	controllers.HandleViews("/tool/timestamp", []string{"./views/timestamp.html", "./views/onlinetoolheader.html", "./views/onlinetooltail.html"}, func(c *app.Context) interface{} {
		cook, _ := c.Request.Cookie("zone")
		zone := "8"
		if cook != nil {
			zone = cook.Value
		}
		n := time.Now()
		return map[string]interface{}{
			"now":  n.Unix(),
			"zone": zone,
		}
	})
	controllers.HandleViews("/tool/diff", []string{"./views/diff.html", "./views/onlinetoolheader.html", "./views/onlinetooltail.html"}, controllers.DataDefaultGetter)
	controllers.HandleViews("/tool/json2ddl", []string{"./views/json2ddl.html", "./views/onlinetoolheader.html", "./views/onlinetooltail.html"}, controllers.DataDefaultGetter)
	controllers.HandleViews("/tool/curl2go", []string{"./views/curl2go.html", "./views/onlinetoolheader.html", "./views/onlinetooltail.html"}, controllers.DataDefaultGetter)
	controllers.HandleViews("/tool/arithmetic", []string{"./views/arithmetic.html", "./views/onlinetoolheader.html", "./views/onlinetooltail.html"}, controllers.DataDefaultGetter)

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
	app.Handle("/tool/jsontomsgpack/exec", func_to_handler.NewFuncToHandler(func(data string) string {
		m := make(map[string]interface{})
		err := json.Unmarshal([]byte(data), &m)
		if err != nil {
			return err.Error()
		}
		res, err := msgpack.Marshal(m)
		if err != nil {
			return err.Error()
		}
		h := hex.EncodeToString(res)
		return h
	}))
	app.Handle("/tool/msgpacktojson/exec", func_to_handler.NewFuncToHandler(func(data string) string {
		d, err := hex.DecodeString(data)
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
	app.Handle("/tool/timestamp/exec", &app.Got{H: controllers.TimestampExec})
	app.Handle("/tool/diff/exec", &app.Got{H: controllers.DiffExec})
	app.Handle("/tool/json2ddl/exec", func_to_handler.NewFuncToHandler(func(data string) string {
		conf := ddlmaker.Config{
			DB: ddlmaker.DBConfig{
				Driver:  "mysql",
				Engine:  "InnoDB",
				Charset: "utf8mb4",
			},
			OutFilePath: os.TempDir(),
		}

		dm, err := ddlmaker.New(conf)
		if err != nil {
			return err.Error()
		}
		res, err := dm.GenerateJSON(data)
		if err != nil {
			return err.Error()
		}
		return string(res)
	}))
	app.Handle("/tool/curl2go/exec", app.Got{H: controllers.Curl2GoExec})
	app.Handle("/tool/arithmetic/exec", func_to_handler.NewFuncToHandler(controllers.ArithmeticExec))
}
