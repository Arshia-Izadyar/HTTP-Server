package main

import (
	"log"

	"github.com/Arshia-Izadyar/HTTP-Server/src/server"
)

func main() {
	m := server.NewServe()
	m.HandlerFunc("/arshia", func(a server.ResponseWriter, b server.Request) {

	})
	m.HandlerFunc("/arshia2", func(a server.ResponseWriter, b server.Request) {

	})
	m.HandlerFunc("/arshia3", func(a server.ResponseWriter, b server.Request) {

	})
	m.HandlerFunc("/arshia4", func(a server.ResponseWriter, b server.Request) {

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
