package main

// AULA - PERCORRENDO ARRAY COM FOR RANGE (EQUIVALE AO FOR EACH)

import "fmt"

func main() {
	// Notação ... - Define que é um array com tamanho fixo, mas o compilador que define
	// o número de elementos com base na inicialização
	numeros := [...]int{1, 2, 3, 4, 5}

	// Range retorna dois elementos: o indice e o elemento do array que está sendo iterado no momento
	for i, numero := range numeros {
		fmt.Printf("%d) %d\n", i, numero)
	}

	// Podemos ignorar o indice, pois uma vez declarada, a variável deve ser usada
	for _, num := range numeros {
		fmt.Println(num)
	}

}
