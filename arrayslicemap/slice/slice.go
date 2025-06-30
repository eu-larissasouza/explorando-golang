package main

// AULA - CONHECENDO SLICE

import (
	"fmt"
	"reflect"
)

func main() {
	// Slice tem um tamanho variável e Array tem tamanho fixo
	// Slice não é um array! Define um pedaço de um array

	a1 := [3]int{1, 2, 3} // array
	s1 := []int{1, 2, 3}  // slice

	fmt.Println(a1, s1)
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	a2 := [5]int{1, 2, 3, 4, 5}

	// Uma forma de obter os valores de um slice é baseado em um intervalo que é composto
	// por dois índices, um de início e outro de final (separados por dois pontos).

	s2 := a2[1:3] // Intervalo semiaberto que inclui o primeiro elemento e exclui o ultimo
	fmt.Println(a2, s2)

	// É uma estrutura contínua que possui um ponteiro para o primeiro elemento que ele aponta no array

	s3 := a2[:2] // Novo slice aponta pro mesmo array
	fmt.Println(a2, s3)

	// Imagine o slice como uma estrutura com um tamanho e um ponteiro pra um elemento de um array

	// Slice de um slice - Aponta pro mesmo array
	s4 := s2[:1]
	fmt.Println(a2, s4)
}
