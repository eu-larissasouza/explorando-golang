package main

import (
	"fmt"
	"time"
)

// AULA - ESTRUTURA DE REPETIÇÃO
// É a unica estrutura de repetição no Go, mas temos diferentes formas de usá-lo

func main() {

	i := 1

	for i <= 10 { // não tem while em Go
		fmt.Println(i)
		i++
	}

	for i := 0; i <= 20; i += 2 {
		fmt.Printf("%d ", i)
	}

	fmt.Println("\nMisturando...")

	for i := 1; i <= 10; i++ {
		if (i % 2) == 0 {
			fmt.Printf("Par ")
		} else {
			fmt.Printf("Impar ")
		}
		fmt.Println()
	}

	for {
		// laço infinito
		fmt.Println("Pra sempre...")
		time.Sleep(time.Second) // aguarda 1 segundo e continua
	}
}
