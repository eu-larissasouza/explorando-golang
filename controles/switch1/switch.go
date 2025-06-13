package main

// AULA - SWITCH #1

import "fmt"

func notaParaConceito(n float64) string {
	nota := int(n)
	switch nota {
	case 10:
		// é como um continue de outras linguagens
		// obriga a executar o proximo case, ao invés de sair do bloco switch
		fallthrough
	case 9:
		return "A"
	case 8, 7:
		// case com multiplos valores
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1, 0:
		return "E"

	default:
		return "Nota inválida"
	}
}

func main() {
	fmt.Println(notaParaConceito(9.8))
	fmt.Println(notaParaConceito(6.9))
	fmt.Println(notaParaConceito(2.1))
}
