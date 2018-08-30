package structs

import "time"

type SearchResult struct {
	Summary *Summary `json:"summary"`
	Docs    []*Doc   `json:"docs"`
}

func NewSearchResult() *SearchResult {
	se := new(SearchResult)
	se.Summary = new(Summary)
	return se
}

type Summary struct {
	Q        string `json:"q"`
	NumDocs  int    `json:"num"`
	Duration int64  `json:"d"`
}

type Doc struct {
	Title  string    `json:"title"`
	Link   string    `json:"link"`
	Des    string    `json:"des"`
	Figure string    `json:"figure"`
	Date   time.Time `json:"date"`
}
