package main

import "fmt"

/*
	AULA - TIPO PERSONALIZADO

	Em Go, podemos criar tipos baseados em tipos primitivos, isso
	nos permite adicionar métodos a esses tipos e melhorar a legibilidade do código.
*/

type nota float64

func (n nota) entre(inicio, fim float64) bool {
	return float64(n) >= inicio && float64(n) <= fim
}

func notaParaConceito(n nota) string {
	switch {
	case n.entre(9, 10):
		return "A"
	case n.entre(8, 9):
		return "B"
	case n.entre(5, 8):
		return "C"
	case n.entre(3, 5):
		return "D"
	default:
		return "E"
	}
}

func main() {
	fmt.Println(notaParaConceito(9.8))
	fmt.Println(notaParaConceito(6.9))
	fmt.Println(notaParaConceito(2.8))
}
