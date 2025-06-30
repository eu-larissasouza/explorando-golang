package main

import "fmt"

// AULA - USANDO FUNÇÕES APPEND E COPY

func main() {
	array1 := [3]int{1, 2, 3}
	var slice1 []int

	// array1 = append(array1, 4,5,6) - só funciona para sliceß
	slice1 = append(slice1, 4, 5, 6)
	fmt.Println(array1, slice1)

	slice2 := make([]int, 2)

	// copy é um metodo que copia dados de um slice para outro
	// -> diferente do append, não expande o tamanho do slice de destino, apenas usa o espaço disponível
	copy(slice2, slice1)
	fmt.Println(slice1, slice2)

}
