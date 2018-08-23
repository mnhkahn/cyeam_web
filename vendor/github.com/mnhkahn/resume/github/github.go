package github

import (
	"fmt"

	"github.com/franela/goreq"
	"github.com/mnhkahn/resume/structs"
)

const (
	GITHUB_HOST = "https://api.github.com"
)

// https://api.github.com/users/mnhkahn
func GitHubProfile(params *structs.Params) (*structs.Profile, error) {
	res, err := goreq.Request{Uri: GITHUB_HOST + "/users/" + params.GitHub}.Do()
	if err != nil {
		return nil, err
	}
	p := new(structs.Profile)
	err = res.Body.FromJsonTo(p)

	return p, err
}

// https://api.github.com/search/repositories?q=user:mnhkahn%20language:go&sort=stars
func GitHubRepos(params *structs.Params) (*structs.Repos, error) {
	res, err := goreq.Request{Uri: fmt.Sprintf("%s/search/repositories?q=user:%s&sort=stars", GITHUB_HOST, params.GitHub)}.Do()
	if err != nil {
		return nil, err
	}
	repos := new(structs.Repos)
	err = res.Body.FromJsonTo(repos)
	if err != nil {
		return repos, err
	}

	items := make([]*structs.Item, 0, len(repos.Items))
	for _, item := range repos.Items {
		if item.Description != nil {
			items = append(items, item)
		}
	}
	repos.Items = items

	if len(repos.Items) > params.RepoLimit {
		repos.Items = repos.Items[:params.RepoLimit]
	}
	return repos, nil
}
