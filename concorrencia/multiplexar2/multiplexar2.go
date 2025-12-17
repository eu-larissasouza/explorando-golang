package main

import (
	"fmt"
	"time"
)

// AULA - MULTIPLEXAR COM SELECT

// internamente vai funcionar como um generator
func falar(pessoa string) <-chan string {
	// como resposta do metodo retorno um canal
	ch := make(chan string)
	// criamos a função anonima
	go func() {
		for i := 0; i < 3; i++ {
			time.Sleep(time.Second)
			ch <- fmt.Sprintf("%s falando %d", pessoa, i)
		}
	}() // () invoca a função anonima imediatamente
	return ch
}

func juntar(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case msg1 := <-entrada1:
				c <- msg1
			case msg2 := <-entrada2:
				c <- msg2
			}
		}
	}()
	return c
}

func main() {
	c := juntar(falar("Maria"), falar("João"))
	fmt.Println(<-c, " | ", <-c)
	fmt.Println(<-c, " | ", <-c)
	fmt.Println(<-c, " | ", <-c)
}
