package main

import "fmt"

/*
	AULA - FUNÇÕES COMO PRIMEIRA CLASSE
	Em Go uma função é um First-Class Citizen, ou seja, ela pode ser tratada como uma variável.

	Outras possibilidades:
	-> posso passar uma função como parâmetro
	-> retornar uma função de outra função.

*/

// Função anônima
var soma = func(a, b int) int {
	return a + b
}

func main() {
	fmt.Println(soma(2, 3))

	// Posso ter uma função local
	sub := func(a, b int) int {
		return a - b
	}
	fmt.Println(sub(2, 3))

}
