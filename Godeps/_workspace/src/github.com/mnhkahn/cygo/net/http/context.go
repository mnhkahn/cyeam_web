package http

import (
	cyurl "cyeam/Godeps/_workspace/src/github.com/mnhkahn/cygo/net/url"
)

type Context struct {
	Req     *Request
	Resp    *Response
	ReqAddr *cyurl.Host
	//	elapse time.Duration
}

func NewContext() *Context {
	ctx := new(Context)
	return ctx
}
