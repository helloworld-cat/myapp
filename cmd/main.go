package main

import (
	"log"
	"net/http"

	"github.com/pagedegeek/myapp/handlers"
)

func main() {
	mux := http.NewServeMux()

	helloHandler := handlers.NewHelloHandler()

	mux.Handle("/", helloHandler)

	addr := "0.0.0.0:8080"
	log.Printf("Listen: %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}
