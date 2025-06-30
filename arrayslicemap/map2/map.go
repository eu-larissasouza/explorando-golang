package main

import "fmt"

// AULA - TRABALHANDO COM MAPS #2

func main() {
	// forma de inicializar o map na declaração da variavel
	funcsESalarios := map[string]float64{
		"José João":      11325.45,
		"Gabriela Silva": 15456.78,
		"Pedro Junior":   1200.0, // ultimo elemento tem que ter virgula
	}

	funcsESalarios["Rafael Filho"] = 1350.0

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}
}
