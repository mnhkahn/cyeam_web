package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"cyeam/Godeps/_workspace/src/github.com/mnhkahn/resume"
	"cyeam/Godeps/_workspace/src/github.com/mnhkahn/resume/structs"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "resume"
	app.Usage = "A resume generator"
	app.Version = resume.VERSION
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "toutiao",
			Value: "148931",
			Usage: "今日头条独家号Subject ID",
		},
		cli.StringFlag{
			Name:  "toutiaocnt",
			Value: "5",
			Usage: "获取今日头条文章数量，默认是5",
		},
		cli.StringFlag{
			Name:  "o",
			Value: "resume.html",
			Usage: "输出结果",
		},
		cli.StringFlag{
			Name:  "github",
			Value: "mnhkahn",
			Usage: "GitHub账号",
		},
		cli.StringFlag{
			Name:  "githubcnt",
			Value: "5",
			Usage: "展示代码仓库数量",
		},
		cli.StringFlag{
			Name:  "weixin",
			Value: "360924857",
			Usage: "微信号",
		},
		cli.StringFlag{
			Name:  "stackoverflow",
			Value: "1924657",
			Usage: "StackOverflow的用户ID",
		},
	}
	app.Action = func(c *cli.Context) error {
		params := new(structs.Params)
		params.TouTiao = c.String("toutiao")
		params.TouTiaoLimit = c.Int("toutiaocnt")

		params.Output = c.String("o")
		params.GitHub = c.String("github")
		params.RepoLimit = c.Int("githubcnt")

		params.Weixin = c.String("weixin")

		params.StackOverflow = c.String("stackoverflow")

		body, err := resume.Resume(params)
		if err != nil {
			return err
		}
		err = ioutil.WriteFile(params.Output, body, 0644)
		if err != nil {
			return fmt.Errorf("Error %v %s", err, params.Output)
		}
		return nil
	}

	app.Run(os.Args)
}
