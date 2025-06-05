package main

// AULA - PRINTS E FORMATACAO EM GO

import (
	"fmt"
)

func main() {
	fmt.Print("Mesma")
	fmt.Print(" linha.")

	fmt.Println(" Nova")
	fmt.Println("linha.")

	x := 3.1415926

	// Não permite concatenar tipos diferentes diretamente
	// fmt.Println("O valor de x é: " + x)

	// Para concatenar, precisamos converter o tipo float64 para string
	xs := fmt.Sprint(x)
	fmt.Println("O valor de x é: " + xs)

	// Imprime o valor de x sem formatação
	fmt.Println("O valor de x:", x)

	// printf - Permite formatar a saída
	fmt.Printf("O valor de x é: %.2f.", x)

	a := 1
	b := 1.999
	c := false
	d := "opa"

	/*
		GANCHOS DE FORMATACAO
		>  %v - Valor padrão do tipo
		>  %d - Inteiro
		>  %.xf - Float com x casas decimais ou %f sem especificar casas decimais
		>  %t - Booleano
		>  %s - String
	*/
	fmt.Printf("\nValores: %d, %f, %.1f, %t, %s", a, b, b, c, d)
}
