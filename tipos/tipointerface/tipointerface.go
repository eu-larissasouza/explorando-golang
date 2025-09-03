package main

/*
	AULA - TIPO INTERFACE
	O tipo interface é um tipo especial que pode receber qualquer outro tipo.
	Ele é como um tipo genérico e dinâmico.

	Esse conceito é útil para abstrair comportamentos comuns entre tipos diferentes
*/

import "fmt"

type curso struct {
	nome string
}

func main() {
	var coisa interface{}
	fmt.Println(coisa)

	coisa = 3.14
	fmt.Println(coisa)

	type dinamico interface {
	}
	var coisa2 dinamico = "Opa"
	fmt.Println(coisa2)

	coisa2 = true
	fmt.Println(coisa2)

	coisa2 = curso{"Golang"}
	fmt.Println(coisa2)

}
