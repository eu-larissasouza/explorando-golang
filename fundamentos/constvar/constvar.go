package main

/*
	AULA - CONSTANTES E VARIÁVEIS EM GO
*/

import (
	"fmt"
	m "math" // m - alias para o pacote math
	// _ "math" - Importaria o pacote math sem usá-lo diretamente
)

func main() {
	// Definimos uma constante ou variavel, seu tipo e o valor
	// Em Go, não é necessário declarar o tipo da variável, o compilador infere automaticamente
	const pi float64 = 3.1415926
	var raio = 3.2

	// Forma reduzida de declarar variáveis
	// Declaramos a variável area e atribuimos o valor
	// Uma variável declarada e não usada gera um erro de compilação
	area := pi * m.Pow(raio, 2)

	fmt.Println("A área da circunferência é:", area)

	// Declaração de constantes e variáveis em blocos
	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

	// Declarando multiplas variáveis
	var e, f bool = true, false
	fmt.Println(e, f)

	// Declarando multiplas variáveis com inferência de tipo
	g, h, i := 2, false, "opa"
	fmt.Println(g, h, i)
}
