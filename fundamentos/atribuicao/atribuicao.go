package main

// AULA - OPERADORES DE ATRIBUIÇÃO

import "fmt"

func main() {
	var b byte = 3
	fmt.Println(b)

	i := 5 // inferencia de tipo
	i += 2 // i = i +2
	i -= 3 // i = i - 3
	i /= 2 // i = i / 2
	i *= 2 // i = i * 2
	i %= 2 // i = i % 2

	fmt.Println(i)

	x, y := 1, 2 // declaracao e atribuicao de variaveis
	fmt.Println(x, y)

	x, y = y, x // troca de valores
	fmt.Println(x, y)
}
