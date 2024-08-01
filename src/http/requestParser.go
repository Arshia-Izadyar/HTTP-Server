package http

import (
	"strings"
)

func ParseHttpRequest(req string) *HttpRequest {
	request := &HttpRequest{}
	requestLines := strings.Split(req, "\r\n\r\n")
	head := strings.Split(requestLines[0], "\r\n")
	body := requestLines[1]
	var requestInfo []string
	requestInfo, head = strings.Split(head[0], " "), head[1:]
	request.Method = requestInfo[0]
	request.Path = requestInfo[1]
	request.HTTPVersion = strings.Split(requestInfo[2], "/")[1]
	headers := make(map[string]string)
	for _, line := range head {
		splitedLine := strings.Split(line, ": ")
		headers[splitedLine[0]] = splitedLine[1]
	}
	request.Headers = headers
	request.Body = body
	return request
}
