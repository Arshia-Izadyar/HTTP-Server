package http

import (
	"errors"
	"fmt"
	"net"
	"regexp"
	"strings"
)

type Handler interface {
	ServeHTTP(req *HttpRequest, c net.Conn)
}

type HandlerFunc func(req *HttpRequest) HttpResponse

type route struct {
	Method  string
	handler HandlerFunc
	Params  QueryParameter
	Path    *regexp.Regexp
}

type Mux struct {
	handlers []route
}

func NewMux() *Mux {
	return &Mux{
		handlers: []route{},
	}
}
func (m *Mux) HandlerFunc(path string, handler HandlerFunc) {

	err := m.setHandler(path, handler)
	if err != nil {
		panic(err)
	}
}
func (m *Mux) setHandler(path string, handler HandlerFunc) error {
	r := route{
		handler: handler,
		Params:  make(QueryParameter),
		Path:    nil,
	}
	if splitedPath := strings.Split(path, " "); len(splitedPath) == 2 {
		path = splitedPath[1]
		for _, method := range methods {
			if method == strings.ToUpper(splitedPath[0]) {
				r.Method = method
				break
			}
		}
		if r.Method == "" {
			return errors.New(fmt.Sprintf("the value %s is not a http method", splitedPath[0]))
		}
	} else if len(splitedPath) > 2 {
		return errors.New(fmt.Sprintf("the %s is not a valid path", path))
	} else {
		r.Method = "*"
	}
	// path = "^" + path + "$"
	path = "^" + regexp.QuoteMeta(path) + "$"
	re := regexp.MustCompile(`:([a-zA-Z0-9\-\.]+)`)
	found := re.FindAllString(path, -1)
	for idx, value := range found {
		r.Params[idx] = map[string]string{value[1:]: ""}
		path = strings.Replace(path, value, `([a-zA-Z0-9\-\.]+)`, 1)
	}
	pathRegex, err := regexp.Compile(path)
	if err != nil {
		return err
	}
	r.Path = pathRegex
	m.handlers = append(m.handlers, r)
	return nil
}

func (m *Mux) parseRequestParams(params *QueryParameter, reqPath string, handlerPath *regexp.Regexp) {
	subMatches := handlerPath.FindStringSubmatch(reqPath)[1:]
	for k, v := range *params {
		for paramName := range v {
			(*params)[k][paramName] = subMatches[k]
		}
	}
}

func (m *Mux) ServeHTTP(req *HttpRequest, c net.Conn) {
	defer func() {
		if r := recover(); r != nil {
			writeError(c, "")
			panic(r)
		}
	}()
	for _, route := range m.handlers {
		if route.Path.MatchString(req.Path) {
			if route.Method != "*" && route.Method != req.Method {
				writeNotFound(c)
				return
			}
			m.parseRequestParams(&route.Params, req.Path, route.Path)
			req.UrlParams = route.Params

			res := convertResponseToHTTP(route.handler(req))
			_, err := c.Write([]byte(res))
			if err != nil {
				writeError(c, err.Error())
			}
			return
		}
	}
	writeNotFound(c)
}

func writeError(c net.Conn, msg string) {
	var response HttpResponse
	if msg == "" {
		response = Cr(500, map[string]string{"status": "error"})
	} else {
		response = Cr(500, map[string]string{"status": msg})
	}
	_, err := c.Write([]byte(convertResponseToHTTP(response)))
	if err != nil {
		panic(err)
	}
}
func writeNotFound(c net.Conn) {
	var response HttpResponse = Cr(404, map[string]string{"status": "not found"})
	_, err := c.Write([]byte(convertResponseToHTTP(response)))
	if err != nil {
		panic(err)
	}
}
