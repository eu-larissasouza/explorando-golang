package main

/*
	AULA 1 - Constantes e Variáveis em Go
*/

import (
	"fmt"
	"math"
)

func main() {
	// Definimos uma constante ou variavel, seu tipo e o valor
	// Em Go, não é necessário declarar o tipo da variável, o compilador infere automaticamente
	const pi float64 = 3.1415926
	var raio = 3.2

	// Forma reduzida de declarar variáveis
	// Declaramos a variável area e atribuimos o valor
	// Uma variavel declarada e não usada gera um erro de compilação
	area := pi * math.Pow(raio, 2)

	fmt.Println("A área da circunferência é:", area)
}
