package main

// AULA - OPERADORES RELACIONAIS

import (
	"fmt"
	"time"
)

func main() {
	// == Faz comparações sobre os valores
	// Quando trabalharmos com ponteiro, ele verificará o endereço de referência dos objetos

	fmt.Println("Strings:", "Banana" == "Banana")
	fmt.Println("!=", 3 != 2)
	fmt.Println(">", 3 > 2)
	fmt.Println("<", 3 < 2)
	fmt.Println(">=", 3 >= 2)
	fmt.Println("<=", 3 <= 2)

	d1 := time.Unix(0, 0)
	d2 := time.Unix(0, 0)

	fmt.Println("Datas:", d1 == d2)
	fmt.Println("Datas:", d1.Equal(d2)) // Verifica apenas o valor

	type Pessoa struct {
		Nome string
	}
	p1 := Pessoa{"João"}
	p2 := Pessoa{"Jo"}

	fmt.Println("Pessoas:", p1 == p2) // Verifica o valor dos campos

}
