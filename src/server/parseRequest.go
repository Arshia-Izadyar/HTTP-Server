package server

import (
	"strings"
)

func ParseHttpRequest(req string, remoteAddr string) *HttpRequest {
	parsedRequest := &HttpRequest{
		Headers:  make(map[string]string), // evade nil map
		UrlParam: make(map[string]string),
	}
	request := strings.Split(req, "\r\n\r\n")
	reqInfo := strings.Split(request[0], "\r\n")

	parsedRequest.Ip = remoteAddr
	splitedFirstLine := strings.Split(reqInfo[0], " ") // METHOD PATH PROTOCOL
	if len(splitedFirstLine) != 3 {
		return nil
	}
	parsedRequest.Method = splitedFirstLine[0]
	parsedRequest.Url = splitedFirstLine[1]

	for _, i := range strings.Split(splitedFirstLine[1], "?")[1:] {
		splitedparam := strings.Split(i, "=")
		parsedRequest.UrlParam[splitedparam[0]] = splitedparam[1]
	}

	parsedRequest.Protocol = splitedFirstLine[2]
	headers := reqInfo[2:]
	for _, i := range headers {
		splitedHeader := strings.Split(i, ": ")
		parsedRequest.Headers[splitedHeader[0]] = splitedHeader[1]
	}
	if parsedRequest.Method != "GET" && len(request) > 1 {
		// try to get body
		parsedRequest.Body = request[1] // Parse body later
	}

	return parsedRequest
}
