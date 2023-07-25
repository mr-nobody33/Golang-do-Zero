package main

import (
	"log"
	"net/http"
)

func main() {
	// HTTP É UM PROTOCOLO DE COMUNICAÇÃO - BASE DA COMUNICAÇÃO WEB

	// CLIENTE (FAZ REQUISIÇÃO) - SERVIDOR (PROCESSA REQUISIÇÃO E ENVIA RESPOSTA)

	// Request - Response

	// Rotas
	// URI - Identificador do Recurso
	// Método - GET, POST, PUT, DELETE

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Olá Mundo!"))
	})

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Carregar página de usuários"))
	})

	log.Fatal(http.ListenAndServe(":5000", nil))
}
