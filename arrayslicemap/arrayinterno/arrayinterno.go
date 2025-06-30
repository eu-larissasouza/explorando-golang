package main

import "fmt"

// AULA - USANDO MESMO ARRAY INTERNO

func main() {
	s1 := make([]int, 10, 20)
	s2 := append(s1, 1, 2, 3)

	fmt.Println(s1, s2)

	// Um mesmo array interno pode ser referenciado em diferentes slices
	// E também por isso, o Slice é uma estrutura dinâmica e flexível
	s1[0] = 7
	fmt.Println(s1, s2)

}
