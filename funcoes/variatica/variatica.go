package main

import "fmt"

/*
	AULA - FUNÇÕES VARIÁTICAS
	Em Go, uma função variática é aquela que pode receber um número variável de argumentos.
	-> Similarmente, no Java temos o operador varargs "..." que permite passar um número variável de argumentos para uma função.
	-> E em JavaScript, podemos usar o operador spread "..." para coletar argumentos em um array.

*/

// Recebe um número de argumentos e retorna a média
func media(numeros ...float64) float64 {
	if len(numeros) == 0 {
		return 0
	}

	total := 0.0
	for _, numero := range numeros {
		total += numero
	}
	return total / float64(len(numeros))
}

func main() {
	resultado := media(1.5, 2.5, 3.0, 4.0, 5.0)
	fmt.Printf("Média: %.2f\n", resultado) // Média: 3.20

	resultado2 := media(10.0, 20.0, 30.0)
	fmt.Printf("Média 2: %.2f\n", resultado2) // Média 2: 20.00

	resultado3 := media()
	fmt.Printf("Média 3: %.2f\n", resultado3) // Média 3: 0.00
}
