package http

type HttpRequest struct {
	Method      string
	Path        string
	HTTPVersion string
	Headers     map[string]string
	Body        any
	UrlParams   QueryParameter
}

type HttpResponse struct {
	Code     int
	Message  string
	Protocol string
	Body     string
	Headers  map[string]string
}
