package main

import "fmt"

// AULA - PONTEIROS
// Um ponteiro é uma variável que armazena o endereço de memória de outro objeto ou variável.

func main() {
	i := 1

	// Notação asterisco: define que a variável é um ponteiro
	// O ponteiro nada mais é que uma referencia de memoria e, se eu alterar aquele dado
	// todas as referencias vão enxergar a alteração

	// Para permitir que uma função altere diretamente um dado, temos que compartilhar o ponteiro
	// Por padrão, a função faz uma cópia da variável e não altera o parâmetro

	var p *int = nil // Ponteiro inteiro sem referencia armazenada

	// Em Go, não temos aritmética de ponteiros -  Não permite operações diretas como p++ ou p--

	p = &i // Atribui o endereço da variável

	*p++ // Desreferencia o ponteiro, acessando o dado armazenado no endereço e incrementa 1
	i++

	fmt.Println("Ponteiro:", p, "Valor:", *p)
	fmt.Println("End. Variavel:", &i, "Valor:", i)

}
