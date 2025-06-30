package main

// AULA - CONSTRUINDO SLICES COM MAKE

import "fmt"

func main() {
	// make - gera um array interno com os valores em 0 e devolve uma slice que faz referência a esse array:
	s := make([]int, 10)
	s[9] = 12
	fmt.Println(s)

	// comprimento len() - numero de elementos do slice
	// capacidade cap(s) - numero maximo de elementos do array subjacente

	// cria um slice com 10 elementos, porém o array interno terá 20 elementos
	s = make([]int, 10, 20)
	fmt.Println(s, len(s), cap(s))

	// anexamos mais elementos ao slice usando append()
	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	fmt.Println(s, len(s), cap(s))

	// o slice vai crescendo dinamicamente
	// já estavámos com a capacidade max, ele cria novos arrays internos e aumenta a capacidade
	s = append(s, 20)
	fmt.Println(s, len(s), cap(s))

}
