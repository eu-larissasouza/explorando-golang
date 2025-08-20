package main

import "fmt"

/*
	AULA - FUNÇÃO VARIÁTICA #2 - COM SLICE

	Quando uma função recebe um número variável de argumentos,
	esses argumentos são tratados como um slice/array dentro da função.

*/

func imprimirAprovados(aprovados ...string) {
	fmt.Println("Lista de aprovados:")
	for i, aprovado := range aprovados {
		fmt.Printf("%d. %s\n", i+1, aprovado)
	}
}

func main() {
	aprovados := []string{"Alice", "Bob", "Charlie", "Diana"}
	imprimirAprovados(aprovados...) // Usando o operador spread para passar o slice como argumentos
	// Explode cada um dos elementos do slice como argumentos individuais

	// Também podemos passar diretamente uma lista de strings
	imprimirAprovados("Eve", "Frank", "Grace")

}
