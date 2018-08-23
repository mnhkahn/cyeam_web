package stackoverflow

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/franela/goreq"
	"github.com/mnhkahn/resume/structs"
)

const (
	STACKOVERFLOW_HOST = "https://api.stackexchange.com"
)

// https://api.stackexchange.com/2.2/users/1924657?order=desc&sort=reputation&site=stackoverflow
func GetStackOverflowProfile(params *structs.Params) (*structs.StackOverflowProfile, error) {
	res, err := goreq.Request{Uri: STACKOVERFLOW_HOST + fmt.Sprintf("/2.2/users/%s?order=desc&sort=reputation&site=stackoverflow", params.StackOverflow)}.Do()
	if err != nil {
		return nil, err
	}
	p := new(structs.StackOverflowProfile)
	err = res.Body.FromJsonTo(p)

	return p, err
}

// https://api.stackexchange.com/2.2/users/1924657/tags?order=desc&min=2&sort=popular&site=stackoverflow
func GetStackOverflowTags(params *structs.Params) (*structs.StackOverflowTags, error) {
	res, err := goreq.Request{Uri: STACKOVERFLOW_HOST + fmt.Sprintf("/2.2/users/%s/tags?order=desc&sort=popular&site=stackoverflow&min=2", params.StackOverflow)}.Do()
	if err != nil {
		return nil, err
	}
	p := new(structs.StackOverflowTags)
	err = res.Body.FromJsonTo(p)

	return p, err
}

func GetStackOverflowAnswersAll(params *structs.Params) ([]*structs.AnswerResult, error) {
	answers, err := GetStackOverflowAnswers(params)
	if err != nil {
		return nil, err
	}
	qids := make([]int, 0, len(answers.Items))
	qa := make(map[int]int, len(answers.Items))
	for _, item := range answers.Items {
		qids = append(qids, item.QuestionID)
		qa[item.QuestionID] = item.AnswerID
	}

	questions, err := GetStackOverflowQuestions(params, qids)
	if err != nil {
		return nil, err
	}

	res := make([]*structs.AnswerResult, 0, len(questions.Items))
	for _, item := range questions.Items {
		ar := new(structs.AnswerResult)
		ar.Title = item.Title
		ar.AID = qa[item.QuestionID]
		ar.QID = item.QuestionID
		res = append(res, ar)
	}

	return res, nil
}

// https://api.stackexchange.com/2.2/users/1924657/answers?order=desc&sort=activity&site=stackoverflow
func GetStackOverflowAnswers(params *structs.Params) (*structs.StackOverflowAnswers, error) {
	res, err := goreq.Request{Uri: STACKOVERFLOW_HOST + fmt.Sprintf("/2.2/users/%s/answers?order=desc&sort=activity&site=stackoverflow", params.StackOverflow)}.Do()
	if err != nil {
		return nil, err
	}
	p := new(structs.StackOverflowAnswers)
	err = res.Body.FromJsonTo(p)

	return p, err
}

// https://api.stackexchange.com/2.2/questions/10105935?order=desc&sort=votes&site=stackoverflow
func GetStackOverflowQuestions(params *structs.Params, qids []int) (*structs.StackOverflowQuestions, error) {
	qidsStr := make([]string, len(qids))
	for i, q := range qids {
		qidsStr[i] = strconv.Itoa(q)
	}
	res, err := goreq.Request{Uri: STACKOVERFLOW_HOST + fmt.Sprintf("/2.2/questions/%s?order=desc&sort=votes&site=stackoverflow", strings.Join(qidsStr, ";"))}.Do()
	if err != nil {
		return nil, err
	}
	p := new(structs.StackOverflowQuestions)
	err = res.Body.FromJsonTo(p)

	return p, err
}
