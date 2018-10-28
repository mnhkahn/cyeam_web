package controllers

import (
	"net/http"
	"strings"

	"github.com/mnhkahn/resume"
	"github.com/mnhkahn/resume/structs"

	"github.com/mnhkahn/asciiimg"
	"github.com/mnhkahn/gogogo/app"
)

func ToolBox(c *app.Context) error {
	c.HTML([]string{"./views/toolbox.html", "./views/head.html", "./views/tail.html"}, nil)
	return nil
}

func Ascii(c *app.Context) error {
	url := c.GetString("url")

	res, err := http.Get("http://" + url)
	if err != nil {
		return err
	}

	ai, err := asciiimg.NewAsciiImg(res.Body)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	c.WriteBytes([]byte(ai.DoByCol(38)))
	return nil
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

func Resume(c *app.Context) error {
	var params *structs.Params
	if strings.Index(c.Request.URL.RawPath, "?") == -1 {
		params = DEFAULT_RESUME_PARAMS
	} else {
		params = new(structs.Params)
		params.TouTiao = c.GetString("toutiao")
		params.TouTiaoLimit, _ = c.GetInt("toutiaocnt")

		params.Output = c.GetString("o")
		params.GitHub = c.GetString("github")
		params.RepoLimit, _ = c.GetInt("githubcnt")

		params.Weixin = c.GetString("weixin")

		params.StackOverflow = c.GetString("stackoverflow")
	}

	body, err := resume.Resume(params)
	if err != nil {
		return err
	}

	c.WriteBytes(body)
	return nil
}

func Robots(c *app.Context) error {
	c.HTML([]string{"./views/robots.txt"}, nil)
	return nil
}

func Sitemap(c *app.Context) error {
	c.HTML([]string{"./static/sitemap.xml"}, nil)
	return nil
}

func Feed(c *app.Context) error {
	http.Redirect(c.ResponseWriter, c.Request, "http://blog.cyeam.com/rss.xml", http.StatusFound)
	return nil
}

func SSLVerify(c *app.Context) error {
	c.HTML([]string{"./static/fileauth.htm"}, nil)
	return nil
}

func GoogleVerify(c *app.Context) error {
	c.HTML([]string{"./static/google97ec3a9b69e1f4db.html"}, nil)
	return nil
}

func Toutiao(c *app.Context) error {
	c.HTML([]string{"./views/toutiao.html", "./views/head.html", "./views/tail.html"}, nil)
	return nil
}

func Neitui(c *app.Context) error {
	c.HTML([]string{"./views/neitui.html", "./views/head.html", "./views/tail.html"}, nil)
	return nil
}

func Favicon(c *app.Context) error {
	c.HTML([]string{"./static/c32.ico"}, nil)
	return nil
}

func CSRF(c *app.Context) error {
	c.WriteBytes([]byte(`
<html lang="en">
<head>
   <title>Statistics</title>
</head>
<body>
<img src="https://haohuo.snssdk.com/product/repin?id=6539897724514730248&action=repin&device_id=37450351129"></img>

</body>
`))

	return nil
}
