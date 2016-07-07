package main

import (
	"os"

	"cyeam/Godeps/_workspace/src/github.com/mnhkahn/Cyeam/models"
	"cyeam/Godeps/_workspace/src/github.com/mnhkahn/cygo/net/http"
)

type MainController struct {
	http.Controller
}

func (this *MainController) Get() {
	this.Ctx.Resp.Body = DEFAULT_HTML
}

func (this *MainController) Bing() {
	this.Ctx.Resp.StatusCode = http.StatusFound
	this.Ctx.Resp.Headers.Add(http.HTTP_HEAD_LOCATION, models.GetBing())
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "5000"
	}

	http.Serve(":" + port)
}

func init() {
	http.Router("/", "GET", &MainController{}, "Get")
	http.Router("/bing", "GET", &MainController{}, "Bing")
}

const (
	DEFAULT_HTML = `<!DOCTYPE html>
<html lang="en" xmlns:wb="https://open.weibo.com/wb">
    
    <head>
        <meta charset="utf-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <meta name="description" content="">
        <meta name="author" content="Bryce">
        <title>
            Cyeam
        </title>
        <link rel="shortcut icon" href="/static/c32.ico" />
        <link href="/static/bootstrap.css" rel="stylesheet" />
        <link href="/static/landing-page.css" rel="stylesheet" />
    </head>
    
    <body>
        <nav class="navbar navbar-default navbar-fixed-top" role="navigation">
            <div class="container">
                <div class="navbar-header">
                    <button type="button" class="navbar-toggle" data-toggle="collapse" data-target=".navbar-ex1-collapse">
                        <span class="sr-only">
                            Toggle navigation
                        </span>
                        <span class="icon-bar">
                        </span>
                        <span class="icon-bar">
                        </span>
                        <span class="icon-bar">
                        </span>
                    </button>
                    <a class="navbar-brand" href="/" style="padding:0px">
                        <img src="/bryce.jpg" style="width:50px">
                    </a>
                </div>
                <div class="collapse navbar-collapse navbar-right navbar-ex1-collapse">
                    <ul class="nav navbar-nav">
                        <li>
                            <a href="http://blog.cyeam.com">
                                Blog
                            </a>
                        </li>
                        <li>
                            <a href="/haixiuzu">
                                骚年，来一发
                            </a>
                        </li>
                        <li>
                            <a href="//www.digitalocean.com/?refcode=b3076e9613a4">
                                <img src="/static/do.png" width="32" border="0" alt="DigitalOcean">
                        </li>
                    </ul>
                </div>
            </div>
        </nav>
        <a name="home">
        </a>
        <div class="intro-header" id="intro-header" style="background: url('/bing') no-repeat center center; padding-top: 0px; padding-bottom: 0px">
            <div class="container" id="container">
                <div class="row">
                    <div class="col-lg-12">
                        <div class="intro-message">
                            <h1>
                                Cyeam
                            </h1>
                            <hr class="intro-divider">
                            <ul class="list-inline intro-social-buttons">
                                <li>
                                    <a href="/resume" class="btn btn-default btn-lg">
                                        <i class="fa fa-twitter fa-fw">
                                        </i>
                                        <span class="network-name">
                                            Resume
                                        </span>
                                    </a>
                                </li>
                                <li>
                                    <a href="//github.com/mnhkahn" class="btn btn-default btn-lg">
                                        <i class="fa fa-github fa-fw">
                                        </i>
                                        <span class="network-name">
                                            Github
                                        </span>
                                    </a>
                                </li>
                                <li>
                                    <a class="btn btn-default btn-lg" data-email="%6c%69%63%68%61%6f%30%34%30%37%40%67%6d%61%69%6c%2e%63%6f%6d"
                                    href="/cdn-cgi/l/email-protection#1975707a717876597a607c7874377a7674">
                                        <i class="fa fa-github fa-fw">
                                        </i>
                                        <span class="network-name">
                                            Email
                                        </span>
                                    </a>
                                </li>
                            </ul>
                        </div>
                    </div>
                </div>
                <div class="container">
                    <div class="row">
                        <div class="col-lg-12">
                            <p class="copyright text-muted small">
                                Copyright &copy; Cyeam 2015. All Rights Reserved
                            </p>
                        </div>
                    </div>
                </div>
            </div>
        </div>
        <script src="/static/jquery-1.10.2.js">
        </script>
        <script src="/static/bootstrap.js">
        </script>
    </body>

</html>	`
)
