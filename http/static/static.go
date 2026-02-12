package main

import (
	"log"
	"net/http"
)

// CRIANDO UM SERVIDOR ESTÁTICO
// Usando esse pacote http, já temos inumeras possibilidades, pois é como se o Golang
// internalisasse um web server, ou seja, não precisamos instalar nenhum outro software para rodar um servidor web

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs) // quando chegar requisição para a raiz do servidor, ele vai servir os arquivos do diretório "public"

	// para rodar, temos que usar a linha de comando e assim, já conseguimos
	// acessar os arquivos do diretório "public" através do navegador
	log.Println("Servidor rodando na porta 3000...")
	log.Fatal(http.ListenAndServe(":3000", nil))

}
