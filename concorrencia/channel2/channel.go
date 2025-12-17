package main

import (
	"fmt"
	"time"
)

// AULA - GOROUTINES E CHANNEL

// Função que multiplica a partir de uma goroutine, recebe informações do resultado por meio do canal
// Retorna informações por meio do canal

// Channel (canal) - é o ponto de sincronismo entre as goroutines, uma forma de comunicação
// Channel é um tipo

func doisTresQuatroVezes(base int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * base // enviando dados para o canal

	time.Sleep(time.Second)
	c <- 3 * base

	time.Sleep(3 * time.Second)
	c <- 4 * base
}

func main() {
	c := make(chan int) // sem buffer
	go doisTresQuatroVezes(10, c)

	// para o código passar pra próxima parte, a informação do canal tem que ser lida
	a, b := <-c, <-c // recebendo dados do canal

	// assim sendo, a gente pode usar os channels pra fazer um sincronismo entre dados
	// de várias goroutines que são executadas nesses processos independentes

	fmt.Println(a, b)

	fmt.Println(<-c)
}
