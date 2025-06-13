package main

import "fmt"

// AULA - IF / ELSE

func imprimirResultado(nota float64) {
	// Estrutura condicional - Não aceita parenteses na expressão, exceto para precedência
	// Obrigatoriamente deve ter a chave para definir o bloco
	if nota >= 7 {
		fmt.Println("Aprovado com nota", nota)
	} else {
		fmt.Println("Reprovado com nota", nota)
	}
}

func main() {
	imprimirResultado(7.3)
	imprimirResultado(5.1)
}
