package http

import (
	"log"
	"reflect"
)

type Handler interface {
	ServeHTTP(ctx *Context)
}

type Handle struct {
	ctrl       ControllerIfac
	methodName string
	fn         reflect.Value
}

func (this *Handle) ServeHTTP(ctx *Context) {
	defer PanicRecover(ctx)
	this.ctrl.Init(ctx)
	this.fn.Call(nil)
}

func PanicRecover(ctx *Context) {
	if err := recover(); err != nil {
		log.Println(err)
		ctx.Resp.StatusCode = StatusNotFound
	}
}
