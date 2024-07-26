package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
)

func (r *HttpResponse) Build() (string, error) {
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
	response += fmt.Sprintf("\r\n\r\n%s\r\n\r\n", r.Body)
	return response, nil
}

func Response(res map[string]string, code int) *HttpResponse {
	body, err := json.Marshal(res)
	if err != nil {
		return nil
	}
	response := &HttpResponse{
		HttpVersion:     "HTTP/1.1",
		StatusCode:      code,
		Status:          http.StatusText(code),
		Body:            string(body),
		ContentType:     "application/json",
		CloseConnection: false,
	}
	return response
}

func (r *HttpResponse) Write(input []byte) (int, error) {
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

func (r *HttpResponse) WriteHeader(statusCode int) {
	r.StatusCode = statusCode
	r.Status = http.StatusText(statusCode)

}

func invalidRequestResponse(cnn net.Conn) {
	invalidResponse := HttpResponse{
		HttpVersion:     "HTTP/1.1",
		StatusCode:      400,
		Status:          "Bad Request",
		Body:            "",
		ContentType:     "text/html",
		CloseConnection: true,
	}
	res, _ := invalidResponse.Build()
	cnn.Write([]byte(res))
}

// TODO: Add debug
func internalServerErrorResponse(cnn net.Conn) {
	invalidResponse := HttpResponse{
		HttpVersion:     "HTTP/1.1",
		StatusCode:      500,
		Status:          "Internal Error",
		Body:            "",
		ContentType:     "text/html",
		CloseConnection: true,
	}
	res, _ := invalidResponse.Build()
	cnn.Write([]byte(res))
}

func notFoundResponse(cnn net.Conn) {
	invalidResponse := HttpResponse{
		HttpVersion:     "HTTP/1.1",
		StatusCode:      404,
		Status:          "Not Found",
		Body:            "",
		ContentType:     "text/html",
		CloseConnection: true,
	}
	res, _ := invalidResponse.Build()
	cnn.Write([]byte(res))
}
