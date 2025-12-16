package main

import (
	"fmt"
	"time"
)

// AULA - CUIDADO COM DEADLOCKS

func rotina(c chan int) {
	time.Sleep(time.Second)
	c <- 1 // operação bloqueante
	fmt.Println("Só depois que channel foi lido")
}

func main() {
	c := make(chan int) // canal sem buffet

	// sem buffer, a primeira operação com channel já bloqueia
	// com buffer, ele faz o envio de dados até alcançar o limite definido

	go rotina(c)

	// só recebe quando os dados estão prontos
	fmt.Println(<-c) // operação bloqueante

	fmt.Println("Foi lido")
	fmt.Println(<-c) // deadlock (não há nenhum dado no channel)

	// GERA UM ERRO POR DEADLOCK
	fmt.Println("FIM")
}
