package main

import "fmt"

// AULA - FUNÇÕES BÁSICAS EM GO

// Funções são blocos de código que podem ser reutilizados em diferentes partes do programa.
// Elas podem receber parâmetros e retornar multiplos valores.
func somar(a int, b int) int {
	return a + b

}

func imprimir(valor int) {
	fmt.Println(valor)
}
