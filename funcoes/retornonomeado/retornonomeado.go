package main

import "fmt"

// AULA - RETORNO NOMEADO
// Funciona como variaveis existentes no escopo da minha função

func trocar(p1, p2 int) (segundo int, primeiro int) {
	segundo = p2
	primeiro = p1
	// retorno limpo = uma vez que definimos aquilo que será retornado, não precisamos especificar
	return
}

func main() {
	x, y := trocar(2, 3)
	fmt.Println("x:", x, "y:", y)

	segundo, primeiro := trocar(7, 5)
	fmt.Println("segundo:", segundo, "primeiro:", primeiro)
}
