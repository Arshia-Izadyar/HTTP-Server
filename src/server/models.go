package server

type HttpResponse struct {
	HttpVersion     string
	StatusCode      int
	Status          string
	Body            string
	ContentType     string
	CloseConnection bool
}

type HttpRequest struct {
	Ip       string
	Url      string
	UrlParam map[string]string
	Method   string // make enums
	Protocol string
	Headers  map[string]string
	Body     string
}
