package http

import (
	"fmt"
	"regexp"
	"strings"
)

func ParseHttpRequest(req, ip string) *HttpRequest {
	request := &HttpRequest{}
	request.Ip = ip
	requestLines := strings.Split(req, "\r\n\r\n")
	head := strings.Split(requestLines[0], "\r\n")
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
	if v, ok := headers["Content-Type"]; ok {
		if v == "application/json" {
			body := requestLines[1]
			request.Body = body
		} else if strings.HasPrefix(v, "multipart/form-data") {
			body := requestLines[1:]
			request.Body = parseFormDataRequest(strings.Join(body, ""))
		}
	} else {
		// no header
		request.Body = strings.Join(requestLines[1:], "")
	}
	request.Headers = headers
	return request
}

func parseFormDataRequest(body string) string {

	a := strings.Split(body, "\r\n")
	boundry := a[0]
	re := regexp.MustCompile(`name="([^"]+)"([^"]+)`)
	mappedBody := ""
	for _, i := range a {
		if i == boundry || i == boundry+"--" || i == "" || strings.Trim(i, " ") == "" {
			continue
		}
		res := re.FindStringSubmatch(i)[1:]
		if len(res) == 2 {
			mappedBody += fmt.Sprintf("%s=%s\n", res[0], res[1])
		} else {
			// TODO: error
		}
	}
	return mappedBody
}
