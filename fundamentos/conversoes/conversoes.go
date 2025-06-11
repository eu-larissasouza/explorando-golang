package main

// AULA - CONVERSÕES EM GO
// Em Go, as conversões de tipos são explícitas, ou seja, você precisa informar Que deseja converter um tipo para outro.
// Temos que tomar cuidado ao misturar tipos, porque o Go é bem rigoroso com manipulação de tipos diferentes.

import (
	"fmt"
	s "strconv"
)

func main() {

	x := 2.4
	y := 2

	fmt.Println(x / float64(y))

	nota := 6.9
	notaFinal := int(nota) // Conversão explícita de float64 para int
	fmt.Println(notaFinal)

	// cuidado, com a conversão explícita para string
	// Retorna um caractere ASCII correspondente ao valor 97, que é 'a'
	fmt.Println("Teste " + string(97))

	// int pra string
	numero := 123
	fmt.Println(s.Itoa(numero))

	// string pra int
	numeroString := "456"

	// Em Go, uma função pode retornar dois valores, nesse caso, o primeiro é o valor convertido e o segundo é um erro
	numeroConvertido, _ := s.Atoi(numeroString)
	fmt.Println(numeroConvertido - 111)

	b, _ := s.ParseBool("true") // Converte string para bool

	if b {
		fmt.Println("Verdadeira")
	}

}
