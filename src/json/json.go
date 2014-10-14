package main

import (
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"gopkg.in/unrolled/render.v1"

	"json/models"
)

var rnd *render.Render

// handlers totalmente fake!
func homeHandler(rw http.ResponseWriter, r *http.Request) {
	user := models.User{"Nassor Paulino da Silva", 30}
	rnd.JSON(rw, http.StatusOK, user)
}

func main() {
	// Criação de um objeto para renderização
	rnd = render.New(render.Options{})

	// Criação do middleware
	n := negroni.New(
		negroni.NewRecovery(),
		negroni.NewLogger(),
	)

	// Router principal da aplicação
	router := mux.NewRouter()
	router.HandleFunc("/", homeHandler)

	// Adicionando o router no middleware.
	n.UseHandler(router)

	// Iniciando a aplicação
	n.Run(":8080")
}
