package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func (r *Response) Build() (string, error) {
	if r.ContentType == "application/json" && r.Body != "" {
		_, err := json.Marshal(r.Body)
		if err != nil {
			return "", errors.New("body is invalid json")
		}
	}
	response := fmt.Sprintf("%s %d %s\r\nContent-Type: %s; charset=UTF-8\r\nContent-Length: %d\r\nServer: go-server/0.0.1\r\n", r.HttpVersion, r.StatusCode, r.Status, r.ContentType, len(r.Body))

	if r.CloseConnection {
		response += "Connection: close\r\n"
	}
	response += fmt.Sprintf("\r\n\r\n %s", r.Body)
	return response, nil
}

func (r *Response) Write(input []byte) (int, error) {
	if len(input) == 0 {
		return 0, errors.New("invlid body length")
	}
	err := json.Unmarshal(input, &struct{}{})
	if err != nil {
		r.ContentType = "text/html"
	} else {
		r.ContentType = "application/json"
	}

	r.Body = string(input)
	return len(r.Body), nil
}

func (r *Response) WriteHeader(statusCode int) {
	r.StatusCode = statusCode
	r.Status = http.StatusText(statusCode)

}
