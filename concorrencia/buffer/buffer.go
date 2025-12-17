package main

// AULA - BUFFER
// Como falado anteriormente, quando o channel tem um buffer,
// o Go faz o envio de dados para o canal até alcançar o limite definido

import (
	"fmt"
	"time"
)

func rotina(ch chan int) {
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	ch <- 5
	fmt.Println("Executou!")
	ch <- 6
}

func main() {
	ch := make(chan int, 3)
	go rotina(ch)

	time.Sleep(2 * time.Second)
	fmt.Println(<-ch)

}
