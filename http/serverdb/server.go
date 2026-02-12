package main

import (
	"log"
	"net/http"
)

// executar nosso server db
func main() {
	http.HandleFunc("/usuarios/", UsuarioHandler)
	log.Println("Starting server at port 3000...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
