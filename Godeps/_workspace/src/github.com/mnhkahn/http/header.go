package http

type Header map[string][]string

func NewHeader() Header {
	h := make(Header)
	h.Add(HTTP_HEAD_SERVER, "Cyeam")
	h.Add(HTTP_HEAD_CONTENTTYPE, "text/html; charset=utf-8")
	return h
}

func (this Header) Add(k, v string) {
	this[k] = append(this[k], v)
}

const (
	HTTP_HEAD_USERAGENT     = "User-Agent"
	HTTP_HEAD_HOST          = "Host"
	HTTP_HEAD_LOCATION      = "Location"
	HTTP_HEAD_SERVER        = "Server"
	HTTP_HEAD_CONTENTTYPE   = "Content-Type"
	HTTP_HEAD_CONTENTLENGTH = "Content-length"
	HTTP_HEAD_DATE          = "Date"
	HTTP_HEAD_ALLOW         = "Allow"
)
