package http

import (
	"fmt"
	"log"
	"net"
)

// TODO: add timeout
type Server struct {
	Port    int
	Addr    string
	Handler Handler
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) ListenAndServe() error {

	li, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.Addr, s.Port))
	if err != nil {
		return err
	}
	for {
		c, err := li.Accept()
		if err != nil {
			return err
		}
		// go handler()

		go handleRequest(c, s.Handler)
	}
}

func handleRequest(c net.Conn, handler Handler) {

	defer func() {
		err := c.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	var request = make([]byte, 1024)
	bytesRed, err := c.Read(request)
	if err != nil {
		// TODO: fix error handleing
		panic(err)
	}
	a := string(request[:bytesRed])
	// parse the handleRequest
	req := ParseHttpRequest(a) // a request with headers / body / etc
	// handlers should handle this
	handler.ServeHTTP(req, c)
	return
}
