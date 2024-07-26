package main

import (
	"log"

	"github.com/Arshia-Izadyar/HTTP-Server/src/server"
)

func main() {
	m := server.NewServe()
	m.HandlerFunc("/", func(r *server.HttpRequest) *server.HttpResponse {
		response := server.Response(map[string]string{"status": "ok"}, 200)

		return response
	})

	var httpServer = &server.Server{
		Addr:    "127.0.0.1",
		Port:    6969,
		Type:    "tcp",
		Handler: *m,
	}
	err := httpServer.Listen()
	if err != nil {
		log.Fatal(err)
	}

}
