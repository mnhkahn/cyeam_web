package controllers

import (
	"fmt"
	"os"
	"sort"
	"time"

	"github.com/ikeikeikeike/go-sitemap-generator/v2/stm"
	"github.com/mnhkahn/gogogo/app"
	"github.com/mnhkahn/gogogo/logger"
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

	patterns := make([]string, 0, len(viewsMap))
	for pattern := range viewsMap {
		patterns = append(patterns, pattern)
	}
	sort.Slice(patterns, func(i, j int) bool {
		return patterns[i] < patterns[j]
	})

	for _, pattern := range patterns {
		viewName := viewsMap[pattern][0]
		fileInfo, err := os.Stat(viewName)
		lastmod := time.Now().Format(time.RFC3339)
		if err != nil {
			logger.Error(err.Error())
		} else {
			lastmod = fileInfo.ModTime().Format(time.RFC3339)
		}
		sm.Add(stm.URL{{"loc", pattern}, {"changefreq", "daily"}, {"priority", "1"}, {"lastmod", lastmod}})
	}

	return sm
}

func SiteMapRaw(c *app.Context) error {
	patterns := make([]string, 0, len(viewsMap))
	for pattern := range viewsMap {
		patterns = append(patterns, pattern)
	}
	sort.Slice(patterns, func(i, j int) bool {
		return patterns[i] < patterns[j]
	})
	for _, pattern := range patterns {
		c.WriteString(pattern + "\n")
	}
	return nil
}

func Robots(c *app.Context) error {
	c.WriteString(fmt.Sprintf(`User-agent: *
Disallow:
Sitemap: https://%s/sitemap.xml
Disallow: /*/exec$`, app.String("host")))
	return nil
}
