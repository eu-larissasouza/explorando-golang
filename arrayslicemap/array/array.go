package main

// AULA - TRABALHANDO COM ARRAYS

import "fmt"

func main() {
	// Array - Estrutura homegênea (mesmo tipo) e estática (tamanho fixo)
	var notas [3]float64
	fmt.Println(notas)

	// Array é indexado
	notas[0], notas[1], notas[2] = 7.8, 4.7, 9.1
	fmt.Println(notas)

	total := 0.0
	for i := 0; i < len(notas); i++ {
		total += notas[i]
	}

	media := total / float64(len(notas))
	fmt.Printf("Média: %.2f\n", media)
}
