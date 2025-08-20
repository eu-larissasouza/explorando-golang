package main

import "fmt"

/*
    AULA - CLOSURE

	Closure é um conceito de programação que se refere a uma função que "fecha" sobre variáveis do seu escopo externo.
	Em Go, são funções que envolvem outras funções e capturam o ambiente em que foram criadas.

	-> A função interna tem ciencia do contexto lexical em que foi criada,
       podendo acessar variáveis do escopo externo mesmo após o escopo externo ter sido finalizado.

	-> Isso permite que a função interna "lembre" do estado das variáveis do escopo externo.
*/

func closure() func() {
	x := 10
	var funcao = func() {
		fmt.Println(x)
	}
	return funcao
}

func main() {
	x := 20
	fmt.Println(x)

	imprimeX := closure()
	imprimeX()
}
