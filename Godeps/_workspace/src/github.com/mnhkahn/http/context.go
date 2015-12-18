package http

import (
	"fmt"
	"time"
)

type Context struct {
	Req    *Request
	Resp   *Response
	elapse time.Duration
}

func NewContext() *Context {
	ctx := new(Context)
	return ctx
}

const (
	LOG_CONTEXT = "%d %s %s %s %s %v"
)

func (this *Context) String() string {
	return fmt.Sprintf(LOG_CONTEXT, this.Resp.StatusCode, this.Req.Method, this.Req.Url, this.Req.UserAgent, this.Req.Host, this.elapse)
}
