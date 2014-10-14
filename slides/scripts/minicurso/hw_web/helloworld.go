package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("Hello World"))
	})
	log.Print("Starting server at port 8080...")
	http.ListenAndServe(":8080", nil)
}
