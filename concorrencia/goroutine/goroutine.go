package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtd int) {
	for i := 0; i < qtd; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteração %d)\n", pessoa, texto, i+1)
	}
}

// vamos fazer experimentos com a função fale
func main() {
	// chamada comum (sem concorrência)
	// processo serial, um processo só termina depois do outro
	//fale("Maria", "Por que você não fala comigo?", 3)
	//fale("João", "Só posso falar depois de você!", 1)

	// chamada com goroutine
	// incluir esse go cria uma goroutine que pode ser excutada paralelamente
	// temos processos independentes

	// uma goroutine só continua se a "thread principal continuar
	// como são duas goroutines, a thread principal termina antes das goroutines

	// go fale("Maria", "Ei...", 10)
	// go fale("João", "Opa...", 5)

	go fale("Maria", "Entendi...", 10)
	fale("João", "Parabens...", 5)

	fmt.Println("Fim!") // essa linha é executada antes das goroutines terminarem

	//
}
