package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./")))
	log.Print("Starting server...")
	http.ListenAndServe(":8123", nil)
}
