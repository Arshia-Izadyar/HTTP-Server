package server

import (
	"errors"
	"fmt"
	"net"
)

type Server struct {
	Addr        string
	Port        uint
	Type        string
	HttpVersion int
	Handler     Serve
}

func (s *Server) Listen() (err error) {

	if s.Port <= 100 {
		return errors.New("server port is invalid")
	}
	li, err := net.Listen(s.Type, fmt.Sprintf("%s:%d", s.Addr, s.Port))
	if err != nil {
		return err
	}

	for {
		cnn, err := li.Accept()
		if err != nil {
			return err
		}
		go s.handleIncomingRequest(cnn)
	}

}

func (s *Server) handleIncomingRequest(cnn net.Conn) {
	defer cnn.Close()

	var byteRequest = make([]byte, 1024)

	bytesRed, err := cnn.Read(byteRequest)

	if err != nil {
		// TODO: FIX ERROR
		panic(err)
	}

	request := ParseHttpRequest(string(byteRequest[:bytesRed]), cnn.RemoteAddr().String())
	if request == nil {
		invalidRequestResponse(cnn)
		return
	}
	// TODO remove

}

func invalidRequestResponse(cnn net.Conn) {
	invalidResponse := Response{
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
