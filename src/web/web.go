package main

import (
	"fmt"
	"log"
	"net/http"
)

func olaMundo(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprint(rw, "Olá mundo!")
}

type StringHandler string

func (s StringHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprint(rw, s)
}

func main() {
	http.HandleFunc("/ola", olaMundo)

	var bonjourStringHandler StringHandler
	bonjourStringHandler = "Bonjour monde"
	http.Handle("/bonjour", bonjourStringHandler)

	http.HandleFunc("/hello", func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte("Hello World"))
	})

	log.Print("Starting server at port 8080...")
	http.ListenAndServe(":8080", nil)
}
