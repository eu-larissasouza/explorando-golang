package main

import "fmt"

/*
	DESAFIO - RECURSIVIDADE SIMPLES
	Solução do desafio de fatorial com recursividade simples

	- Utilizar uint para evitar números negativos
*/

func fatorial(n uint) uint {
	if n == 0 {
		return 1
	}
	return n * fatorial(n-1)

}

func main() {
	resultado := fatorial(6)
	fmt.Println(resultado)

	//resultado2 := fatorial(-4)
	//fmt.Println(resultado2)
}
