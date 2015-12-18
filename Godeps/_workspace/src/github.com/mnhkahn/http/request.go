package http

import (
	"bytes"
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

	Raw bytes.Buffer
}

func NewRequest() *Request {
	r := new(Request)
	return r
}

func (this *Request) Init() {
	b := string(this.Raw.Bytes())
	this.Headers = make(map[string][]string, 0)

	startLine := strings.Split(b[:strings.Index(b, CRLF)], " ")
	if len(startLine) == 3 {
		this.Method, this.Url, this.Proto = startLine[0], startLine[1], startLine[2]
	}
	b = b[strings.Index(b, CRLF)+len(CRLF):]

	if strings.LastIndex(b, CRLF+CRLF) != -1 {
		this.Body = b[strings.LastIndex(b, CRLF+CRLF):]
		b = b[:strings.LastIndex(b, CRLF+CRLF)-2]
	}
	b = strings.TrimSpace(b)

	for _, line := range strings.Split(b, CRLF) {
		k, v := line[:strings.Index(line, ":")], line[strings.Index(line, ":")+1:]
		k, v = strings.TrimSpace(k), strings.TrimSpace(v)
		if k == HTTP_HEAD_USERAGENT {
			this.UserAgent = v
		} else if k == HTTP_HEAD_HOST {
			this.Host = v
		} else {
			this.Headers[k] = append(this.Headers[k], v)
		}
	}
}
