package main

import (
	"log"

	"github.com/Arshia-Izadyar/HTTP-Server/src/http"
	"github.com/Arshia-Izadyar/HTTP-Server/src/impl"
)

func main() {
	mx := http.NewMux()
	mx.HandlerFunc("GET /", impl.ResponseHtml)
	mx.HandlerFunc("GET /file/", impl.ServeImage)
	mx.HandlerFunc("GET /file/:filename", impl.ServeFile)
	mx.HandlerFunc("POST /echo/body", impl.EchoBody)
	mx.HandlerFunc("GET /echo/:echo", impl.EchoParameter)

	server := &http.Server{
		Port:    6969,
		Addr:    "127.0.0.1",
		Handler: mx,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}

/*
C vs Rust

	C =>
		1. cool
		2. resources
		3. new prespective
		5. Portabillity with zig

	Rust =>
		1. cool
		2. new prespective
		3. i Like it
		4. maybe less difficalt than C

*/
