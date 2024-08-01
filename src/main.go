package main

import (
	"log"

	"github.com/Arshia-Izadyar/HTTP-Server/src/http"
)

func main() {
	mx := http.NewMux()
	mx.HandlerFunc("/", func(req *http.HttpRequest) http.HttpResponse { return http.HttpResponse{} })

	mx.HandlerFunc("/prod/:test", func(req *http.HttpRequest) http.HttpResponse {
		test, err := req.UrlParams.Get("test")
		if err != nil {

		}
		return http.Cr(200, map[string]string{"test": test})
	})

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
