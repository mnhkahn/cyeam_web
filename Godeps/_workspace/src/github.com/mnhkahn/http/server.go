package http

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

type Address struct {
	Host string
	Port string
}

func NewAddress(addr_str string) *Address {
	addr := new(Address)
	addr_strs := strings.Split(addr_str, ":")
	if len(addr_strs) == 2 {
		addr.Host, addr.Port = addr_strs[0], addr_strs[1]
	}
	return addr
}

func (this *Address) String() string {
	return this.Host + ":" + this.Port
}

type Server struct {
	Addr *Address
}

var DEFAULT_SERVER *Server

func Serve(addr string) {
	DEFAULT_SERVER.Addr = NewAddress(addr)
	ln, err := net.Listen("tcp", DEFAULT_SERVER.Addr.String())
	defer ln.Close()
	if err != nil {
		panic(err)
	}

	log.Printf("<<<Server Accepting on Port %s>>>\n\n", DEFAULT_SERVER.Addr.Port)
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Panicln(err)
		}
		go handleConnection(conn)
	}
}

func init() {
	DEFAULT_SERVER = new(Server)
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	log.Printf("[%s]<<<Request From %s>>>\n", time.Now().String(), conn.RemoteAddr())

	ctx := NewContext()

	buf := make([]byte, 1024)
	reqLen, err := conn.Read(buf)
	if err != nil {
		log.Println("Error to read message because of ", err)
		return
	}
	ctx.Req = NewRequst(string(buf[:reqLen-1]))

	serve_time := time.Now()
	buffers := bytes.Buffer{}
	buffers.WriteString("HTTP/1.1 200 OK\r\n")
	buffers.WriteString("Server: Cyeam\r\n")
	buffers.WriteString("Date: " + serve_time.Format(time.RFC1123) + "\r\n")
	buffers.WriteString("Content-Type: text/html; charset=utf-8\r\n")
	buffers.WriteString("Content-length:" + fmt.Sprintf("%d", len(DEFAULT_HTML)) + "\r\n")
	buffers.WriteString("\r\n")
	buffers.WriteString(DEFAULT_HTML)
	_, err = conn.Write(buffers.Bytes())
	if err != nil {
		log.Println(err)
	}
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
        <link rel="shortcut icon" href="https://cyeam.com/static/c32.ico" />
        <link href="https://cyeam.com/static/css/bootstrap.css" rel="stylesheet" />
        <link href="https://cyeam.com/static/css/landing-page.css" rel="stylesheet" />
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
                    <a class="navbar-brand" href="https://www.cyeam.com" style="padding:0px">
                        <img src="http://cyeam.qiniudn.com/bryce.jpg" style="width:50px">
                    </a>
                </div>
                <div class="collapse navbar-collapse navbar-right navbar-ex1-collapse">
                    <ul class="nav navbar-nav">
                        <li>
                            <a href="https://blog.cyeam.com">
                                Blog
                            </a>
                        </li>
                        <li>
                            <a href="https://www.cyeam.com/haixiuzu">
                                骚年，来一发
                            </a>
                        </li>
                        <li>
                            <a href="https://www.digitalocean.com/?refcode=b3076e9613a4">
                                <img src="https://cyeam.com/static/img/do.png" width="32" border="0" alt="DigitalOcean">
                        </li>
                    </ul>
                </div>
            </div>
        </nav>
        <a name="home">
        </a>
        <div class="intro-header" id="intro-header" style="background: url('https://cn.bing.com/az/hprichbg/rb/PalmTreePantanal_EN-US12619823667_1366x768.jpg') no-repeat center center; padding-top: 0px; padding-bottom: 0px">
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
                                    <a href="https://github.com/mnhkahn" class="btn btn-default btn-lg">
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
        <script src="https://cyeam.com/static/js/jquery-1.10.2.js">
        </script>
        <script src="https://cyeam.com/static/js/bootstrap.js">
        </script>
    </body>

</html>	`
)
