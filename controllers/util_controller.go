package controllers

import (
	"github.com/mnhkahn/gogogo/app"
)

var viewsMap = map[string][]string{}

func HandleViews(pattern string, views []string, dataGetter func(c *app.Context) interface{}) {
	viewsMap[pattern] = views
	app.Handle(pattern, &app.Got{H: func(c *app.Context) error {
		c.HTML(views, dataGetter(c))
		return nil
	}})
}

func DataDefaultGetter(c *app.Context) interface{} {
	return nil
}
