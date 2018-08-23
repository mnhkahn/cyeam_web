package resume

import (
	"bytes"
	"log"
	"text/template"

	"github.com/mnhkahn/resume/github"
	"github.com/mnhkahn/resume/stackoverflow"
	"github.com/mnhkahn/resume/structs"
	"github.com/mnhkahn/resume/toutiao"
)

const (
	VERSION = "1.0"

	DEFAULT_MAX_LIMIT = 100
	DEFAULT_LIMIT     = 5
)

func Resume(params *structs.Params) ([]byte, error) {
	var err error

	params = HandleParams(params)

	res := new(structs.ResumeResult)
	res.Weixin = params.Weixin
	res.TouTiao, res.HotArticles, err = toutiao.TouTiaoHotArticles(params)
	if err != nil {
		return nil, err
	}

	res.Profile, err = github.GitHubProfile(params)
	if err != nil {
		return nil, err
	}

	res.Repos, err = github.GitHubRepos(params)
	if err != nil {
		return nil, err
	}

	res.StackOverflowProfile, err = stackoverflow.GetStackOverflowProfile(params)
	if err != nil {
		return nil, err
	}

	res.StackOverflowTags, err = stackoverflow.GetStackOverflowTags(params)
	if err != nil {
		return nil, err
	}

	res.Answers, err = stackoverflow.GetStackOverflowAnswersAll(params)
	if err != nil {
		return nil, err
	}

	body, err := GenerateHTML(params, res, "./views/resume.html")
	if err != nil {
		return nil, err
	}

	return body, nil
}

func HandleParams(params *structs.Params) *structs.Params {
	if params.TouTiaoLimit > DEFAULT_MAX_LIMIT {
		params.TouTiaoLimit = DEFAULT_LIMIT
	}

	if params.RepoLimit > DEFAULT_MAX_LIMIT {
		params.RepoLimit = DEFAULT_LIMIT
	}

	return params
}

// GetStackOverflowAnswers
func GenerateHTML(params *structs.Params, res *structs.ResumeResult, templ string) ([]byte, error) {
	t, err := template.ParseFiles(templ)
	if err != nil {
		return nil, err
	}

	body := new(bytes.Buffer)
	err = t.Execute(body, res)
	if err != nil {
		return nil, err
	} else {
		return body.Bytes(), nil
	}
	return nil, nil
}

func init() {
	log.SetFlags(log.Lshortfile)
}
