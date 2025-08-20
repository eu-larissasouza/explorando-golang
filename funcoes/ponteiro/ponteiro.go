package main

import "fmt"

/*
	AULA - PONTEIROS EM FUNÇÕES
	Utilizamos ponteiros para alterar diretamente o valor de uma variável em uma função
	- Ponteiro é uma referência de memória
	- O ponteiro é representado pelo tipo *Tipo, e o valor apontado é representado por *ponteiro
*/

func inc1(n int) {
	n++ // n = n + 1
}

func inc2(n *int) {
	// revisão: * é usado para desreferenciar, ou seja, obter o valor apontado pelo ponteiro
	// ter acesso ao valor da variável que o ponteiro aponta
	*n++
}

func main() {
	n := 1

	// priorize esse modo:
	inc1(n) // passagem por valor
	fmt.Println(n)

	// revisão: & usado para obter o endereço da variável
	inc2(&n) // passagem por referência
	fmt.Println(n)

}
