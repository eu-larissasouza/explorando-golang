package main

import "fmt"

/* AULA - PADRÃO DE CONCORRÊNCIA: GENERATORS

O Generator Pattern em Go (Golang) é um padrão de concorrência que usa goroutines
e channels para criar funções que produzem uma sequência de valores sob demanda,
de forma preguiçosa (lazy), sem gerar todos de uma vez, sendo ótimo para iteradores,
grandes conjuntos de dados ou fontes de dados infinitas, economizando memória
e permitindo processamento paralelo.

*/

import (
	"fmt"
)

func main() {
	t1 := titulo("https://www.cod3r.com.br", "https://www.google.com")
	t2 := titulo("https://github.com/", "https://www.youtube.com")

	fmt.Println("Primeiros:", <-t1, "|", <-t2)
	fmt.Println("Segundos:", <-t1, "|", <-t2)
}
