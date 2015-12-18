package http

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"reflect"
	"strings"
	"time"
)

var ErrLog *log.Logger

var HTTP_METHOD = map[string]string{
	"GET": "GET",
	//	"POST":    "POST",
	//	"HEAD":    "HEAD",
	//	"PUT":     "PUT",
	//	"TRACE":   "TRACE",
	"OPTIONS": "OPTIONS",
	//	"DELETE":  "DELETE",
}

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
	Addr             *Address
	Routes           *Route
	AllowHttpMethods []string
}

var DEFAULT_SERVER *Server

func Serve(addr string) {
	DEFAULT_SERVER.Addr = NewAddress(addr)
	ln, err := net.Listen("tcp", DEFAULT_SERVER.Addr.String())
	defer ln.Close()
	if err != nil {
		panic(err)
	}

	log.Printf("<<<Server Accepting on Port %s>>>\n", DEFAULT_SERVER.Addr.Port)
	for {
		conn, err := ln.Accept()
		if err != nil {
			ErrLog.Println(err)
		}
		go handleConnection(conn)
	}
}

func init() {
	DEFAULT_SERVER = new(Server)
	DEFAULT_SERVER.Routes = NewRoute()

	Router("/", "OPTIONS", &Controller{}, "Option")

	errlogFile, logErr := os.OpenFile("error.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if logErr != nil {
		fmt.Println("Fail to find", "error.log", " start Failed")
	}

	ErrLog = log.New(errlogFile, "", log.LstdFlags|log.Llongfile)
}

func handleConnection(conn net.Conn) {
	serve_time := time.Now()

	defer conn.Close()

	ctx := NewContext()
	ctx.Req = NewRequest()
	ctx.Resp = NewResponse()

	for {
		buf := make([]byte, 1024)
		reqLen, err := conn.Read(buf)
		if reqLen > 0 && err == nil {
			ctx.Req.Raw.Write(buf)
			if reqLen < len(buf) && err == nil {
				break
			}
		} else if err == io.EOF {
			break
		} else if neterr, ok := err.(net.Error); ok && neterr.Timeout() {
			break
		} else {
			ErrLog.Println("Error to read message because of ", err, reqLen)
			ctx.Resp.StatusCode = StatusInternalServerError
			break
		}
	}

	ctx.Req.Init()
	ctx.Resp.Proto = ctx.Req.Proto
	if DEFAULT_SERVER.Routes.routes[ctx.Req.Method][ctx.Req.Url] != nil {
		DEFAULT_SERVER.Routes.routes[ctx.Req.Method][ctx.Req.Url].ServeHTTP(ctx)
	} else {
		if _, exists := HTTP_METHOD[ctx.Req.Method]; !exists {
			ctx.Resp.StatusCode = StatusMethodNotAllowed
		} else {
			ctx.Resp.StatusCode = StatusNotFound
		}
	}

	if ctx.Resp.StatusCode == StatusNotFound {
		ctx.Resp.Body = DEFAULT_ERROR_PAGE
	}
	ctx.Resp.Headers.Add(HTTP_HEAD_DATE, serve_time.Format(time.RFC1123))
	ctx.Resp.Headers.Add(HTTP_HEAD_CONTENTLENGTH, fmt.Sprintf("%d", len(ctx.Resp.Body)))

	buffers := bytes.Buffer{}
	buffers.WriteString(fmt.Sprintf("%s %d %s\r\n", ctx.Resp.Proto, ctx.Resp.StatusCode, StatusText(ctx.Resp.StatusCode)))
	for k, v := range ctx.Resp.Headers {
		for _, vv := range v {
			buffers.WriteString(fmt.Sprintf("%s: %s\r\n", k, vv))
		}
	}
	buffers.WriteString("\r\n")
	buffers.WriteString(ctx.Resp.Body)
	_, err := conn.Write(buffers.Bytes())
	if err != nil {
		ErrLog.Println(err)
	}
	ctx.elapse = time.Now().Sub(serve_time)
	log.Println(ctx)
}

func Router(path string, method string, ctrl ControllerIfac, methodName string) {
	if _, exists := HTTP_METHOD[method]; !exists {
		ErrLog.Println("Method not allowed", method, path, methodName)
		return
	}
	handler := new(Handle)
	handler.ctrl = ctrl
	handler.methodName = methodName
	handler.fn = reflect.ValueOf(handler.ctrl).MethodByName(handler.methodName)
	DEFAULT_SERVER.Routes.routes[method][path] = handler
}

var DEFAULT_ERROR_PAGE = "<iframe scrolling='no' frameborder='0' src='http://yibo.iyiyun.com/js/yibo404/key/2354' width='640' height='464' style='display:block;'></iframe>"
