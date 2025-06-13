package main

// AULA - SWITCH #2

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	// procura o primeiro case com expressão verdadeira
	switch { // mesma coisa de usar "switch true"
	case t.Hour() < 12:
		fmt.Println("Bom dia")
	case t.Hour() < 18:
		fmt.Println("Boa tarde")
	default:
		fmt.Println("Boa noite")
	}
}
