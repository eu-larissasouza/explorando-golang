package main

import "fmt"

// AULA - OPERADORES UNÁRIOS

func main() {
	x := 1
	y := 2

	// Em Go, só temos incremento pós fixado
	x++ // x += 1 ou x = x + 1

	y-- // y -= 1 ou y = y - 1

	// Não permite comparação direta com incremento:
	// fmt.Println(x == y--)

	fmt.Println("x:", x, "y:", y)
}
