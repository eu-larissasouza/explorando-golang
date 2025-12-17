package main

import "fmt"

// AULA - CHANNELS

// Channels são a forma de comunicação entre goroutines
// Criamos ele de uma forma similar a um slice com make
// Channels são tipos, então podemos criar um channel de int, string, etc

func main() {
	// channel trafega dados de um tipo específico
	ch := make(chan int, 1) // criando um channel de int com buffer de 1
	// buffer é a quantidade de dados que o channel pode armazenar

	ch <- 1 // enviando dados para o channel (escrita)
	<-ch    // recebendo dados do channel (leitura)

	ch <- 2
	fmt.Println(<-ch)

	// o canal é um ponto de sincronização com a goroutine
	
}
