package http

type Context struct {
	Req  *Request
	Resp *Response
}

func NewContext() *Context {
	ctx := new(Context)
	return ctx
}
