package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

type QueryParameter map[int]map[string]string

func (q QueryParameter) Get(name string) (string, error) {
	for _, v := range q {
		if v, ok := v[name]; ok {
			return v, nil
		}
	}
	return "", errors.New("value not found")
}

func convertResponseToHTTP(res HttpResponse) (result string) {
	result = fmt.Sprintf("%s %d %s\r\n", res.Protocol, res.Code, res.Message)
	for k, v := range res.Headers {
		result += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	result += "\r\n"
	result += res.Body
	result += "\r\n\r\n"
	return
}

// create response
// TODO: add error
func Cr(code int, body map[string]string) HttpResponse {
	json, err := json.Marshal(body)
	if err != nil {
		return HttpResponse{}
	}
	headers := map[string]string{
		"Content-Type":   "application/json",
		"Content-Length": strconv.Itoa(len(string(json))),
	}
	return HttpResponse{
		Code:     code,
		Message:  http.StatusText(code),
		Protocol: "HTTP/1.1",
		Body:     string(json),
		Headers:  headers,
	}
}

func Response(code int, body, contentType string) HttpResponse {
	headers := map[string]string{
		"Content-Type":   contentType,
		"Content-Length": strconv.Itoa(len(body)),
	}
	return HttpResponse{
		Code:     code,
		Message:  http.StatusText(code),
		Protocol: "HTTP/1.1",
		Body:     body,
		Headers:  headers,
	}
}
