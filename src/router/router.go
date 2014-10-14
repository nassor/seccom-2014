package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homeHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Home")
}

func searchHandler(rw http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	query := mux.Vars(r)["query"]
	fmt.Fprintln(rw, query, page)
}

func main() {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/{query}", searchHandler)

	log.Print("Starting server at port 8080...")
	http.ListenAndServe(":8080", r)
}
