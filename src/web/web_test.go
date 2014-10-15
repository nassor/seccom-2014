package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStringHandler(t *testing.T) {
	var stringHandler StringHandler
	stringHandler = "teste"

	response := httptest.NewRecorder() // cirando
	request, _ := http.NewRequest("GET", "", nil)

	stringHandler.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("Resultado esperado: %s / Resultado recebido:  %s", http.StatusOK, response.Code)
	}

	if response.Body.String() != "teste" {
		t.Fatalf("Resultado esperado: %s / Resultado recebido:  %s", "teste", response.Body.String())
	}
}

func TestOlaMundo(t *testing.T) {
	response := httptest.NewRecorder() // cirando
	request, _ := http.NewRequest("GET", "", nil)

	olaMundo(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("Resultado esperado: %s / Resultado recebido:  %s", http.StatusOK, response.Code)
	}

	if response.Body.String() != "Olá mundo!" {
		t.Fatalf("Resultado esperado: %s / Resultado recebido:  %s", "Olá mundo!", response.Body.String())
	}
}
