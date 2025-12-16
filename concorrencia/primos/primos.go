package main

import (
	"fmt"
	"time"
)

// USANDO RANGE E CLOSE

func isPrimo(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func primos(n int, ch chan int) {
	inicio := 2

	// procura os primeiros n primos
	for i := 0; i < n; i++ {
		for primo := inicio; ; primo++ {
			if isPrimo(primo) {
				ch <- primo
				inicio = primo + 1
				time.Sleep(time.Millisecond * 100)
				break
			}
		}
	}
	// diz que o canal nunca mais vai ser usado,
	// e isso previne possiveis deadlocks
	close(ch)
}

func main() {
	c := make(chan int, 20)

	// cap(c) retorna a capacidade (buffer) do canal)
	go primos(cap(c), c)

	// range que usa o channel e quando fecha o canal, sai do laço
	for primo := range c {
		fmt.Printf("%d ", primo)
	}
	fmt.Println("FIM!")
}
