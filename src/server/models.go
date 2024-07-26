package server

type Response struct {
	HttpVersion     string
	StatusCode      int
	Status          string
	Body            string
	ContentType     string
	CloseConnection bool
}

type Request struct {
	Ip       string
	Url      string
	UrlParam map[string]string
	Method   string // make enums
	Protocol string
	Headers  map[string]string
	Body     string
}
