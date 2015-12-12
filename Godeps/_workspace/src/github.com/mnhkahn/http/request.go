package http

import (
	"strings"
)

type Request struct {
	Method string

	Url string

	Proto string

	UserAgent string

	Host string

	Headers Header

	Body string
}

type Header map[string][]string

func NewRequst(b string) *Request {
	r := new(Request)
	r.Headers = make(map[string][]string, 0)
	for i, line := range strings.Split(b, CRLF) {
		if i == 0 {
			startLine := strings.Split(line, " ")
			if len(startLine) == 3 {
				r.Method, r.Url, r.Proto = startLine[0], startLine[1], startLine[2]
			}
		} else {
			// kv := strings.Split(line, ":")
			// r.Headers[kv[0]] = append(r.Headers[kv[0]], strings.TrimSpace(kv[1]))
			if line != "" {
				if i == len(strings.Split(b, CRLF))+1 {
					r.Body = line
				}
			}
		}
	}

	return r
}
