package controllers

import (
	"fmt"
	"strings"

	"github.com/ikeikeikeike/go-sitemap-generator/v2/stm"
	"github.com/mnhkahn/gogogo/app"
)

func SiteMapXML(c *app.Context) error {
	sm := buildSitemap()
	c.WriteBytes(sm.XMLContent())
	return nil
}
func buildSitemap() *stm.Sitemap {
	sm := stm.NewSitemap(1)
	sm.SetDefaultHost(fmt.Sprintf("https://%s", app.String("host")))
	sm.Create()
	for _, pattern := range app.GoEngine.Patterns() {
		if strings.HasPrefix(pattern, "/tool") || pattern == "/" {
			sm.Add(stm.URL{{"loc", pattern}, {"changefreq", "daily"}, {"priority", "1"}})
		}
	}

	return sm
}

func SiteMapRaw(c *app.Context) error {
	for _, pattern := range app.GoEngine.Patterns() {
		if strings.HasPrefix(pattern, "/tool") || pattern == "/" {
			c.WriteString(pattern + "\n")
		}
	}
	return nil
}

func Robots(c *app.Context) error {
	c.WriteString(fmt.Sprintf(`User-agent: *
Sitemap: https://%s/sitemap.xml`, app.String("host")))
	return nil
}
