package server

import (
	"errors"
	"fmt"
	"net"
	"regexp"
	"strings"
)

type HandlerFunc func(*HttpRequest) *HttpResponse

type Serve struct {
	patterns []Pattern
}

type ResponseWriter interface {
	Write([]byte) (int, error)
	WriteHeader(statusCode int)
}

type Pattern struct {
	str     *regexp.Regexp
	method  string
	handler HandlerFunc
}

func NewServe() *Serve {
	return &Serve{}
}

func (s *Serve) HandlerFunc(pattern string, f HandlerFunc) error {
	if pattern == "" {
		return errors.New("invalid pattern")

	}
	s.sethandler(pattern, f)
	return nil
}

func (s *Serve) sethandler(pattern string, f HandlerFunc) {
	p := Pattern{}
	var pateernStr string = ""
	var err error
	splitedPattern := strings.Split(pattern, " ")
	if len(splitedPattern) == 2 {
		p.method = splitedPattern[0]
		pateernStr = splitedPattern[1]
	} else if len(splitedPattern) == 1 {
		p.method = "*"
		pateernStr = pattern
	}

	if !strings.HasPrefix(pateernStr, "/") {
		pateernStr = "^\\/" + pateernStr + "$"
	} else {
		pateernStr = "^\\" + pateernStr + "$"
	}

	p.str, err = regexp.Compile(pateernStr)
	if err != nil {
		panic(errors.New("pattern is not standard format"))
	}
	p.handler = f
	s.patterns = append(s.patterns, p)
}

func (s *Serve) ServHTTP(req *HttpRequest, c net.Conn) {
	var matchedPattern *Pattern
	for _, p := range s.patterns {
		if p.str.MatchString(req.Url) {
			matchedPattern = &p
			break
		}
	}
	if matchedPattern != nil {

		res := matchedPattern.handler(req)
		strRes, err := res.Build()

		fmt.Println(strRes)

		if err != nil {
			internalServerErrorResponse(c)
		}
		c.Write([]byte(strRes))
	} else {
		notFoundResponse(c)
	}

}
