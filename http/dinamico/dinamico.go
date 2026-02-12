package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// sempre que chegar uma requisição para a rota "/hora", a função horaCerta vai ser chamada
func horaCerta(w http.ResponseWriter, r *http.Request) {
	// aqui, podemos usar a função time.Now() para pegar a hora atual e escrever na resposta
	// o 03:04:05 é o formato de hora que queremos, ou seja, horas:minutos:segundos
	s := time.Now().Format("02/01/2006 03:04:05")
	fmt.Fprintf(w, "<h1> Hora certa: %s</h1>", s) // escreve a string no writer
}

func main() {
	// a função que o handle utiliza deve ter um padrao, com parametros response e request
	http.HandleFunc("/hora", horaCerta)
	log.Println("Executando...")
	log.Fatal(http.ListenAndServe(":3000", nil)) // escuta a porta e disponibiliza nela

}
