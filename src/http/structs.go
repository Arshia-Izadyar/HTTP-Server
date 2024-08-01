package http

type HttpRequest struct {
	Method      string
	Path        string
	Ip          string
	HTTPVersion string
	Headers     map[string]string
	Body        string
	UrlParams   QueryParameter
}

type HttpResponse struct {
	Code     int
	Message  string
	Protocol string
	Body     string
	Headers  map[string]string
}
