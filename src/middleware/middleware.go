package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

// handlers totalmente fake!
func homeHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Home")
}

func adminHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Admin Works!")
}

func searchHandler(rw http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	query := mux.Vars(r)["query"]
	fmt.Fprintln(rw, query, page)
}

// Middleware que possui alguma lógica
func adminMiddleware(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	if r.URL.Query().Get("password") == "12345" {
		next(rw, r)
	} else {
		http.Error(rw, "Not Authorized", 401)
		log.Println("Falha na tentativa de acesso:", r.URL.String())
	}
}

func main() {
	// Criação do middleware
	n := negroni.New(
		negroni.NewRecovery(),
		negroni.NewLogger(),
	)

	// Router principal da aplicação
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/q/{query}", searchHandler)

	// Adicionando um router específico para área de admin
	rAdmin := mux.NewRouter()
	rAdmin.HandleFunc("/admin", adminHandler)

	// Adicionando-o na rota principal, sendo que ele só vai ser executado após a execução do middleware
	r.Handle("/admin", negroni.New(
		negroni.HandlerFunc(adminMiddleware),
		negroni.Wrap(rAdmin),
	))

	// Adicionando o router no middleware.
	n.UseHandler(r)

	// Iniciando a aplicação
	n.Run(":8080")
}
