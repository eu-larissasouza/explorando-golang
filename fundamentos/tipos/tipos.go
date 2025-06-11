package main

// AULA - TIPOS DE DADOS EM GO
import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// TIPOS INTEIROS
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro: é ", reflect.TypeOf(1000))

	// TIPOS DE RANGE INTEIRO
	// -- Sem sinal (positivo)... uint8, uint16, uint32, uint64
	var b byte = 255 // byte é um alias para uint8
	fmt.Println("O byte é", reflect.TypeOf(b))

	// -- Com sinal (positivo e negativo)... int8, int16, int32, int64
	i1 := math.MaxInt64
	fmt.Println("Valor máximo do int é", i1)

	var i2 rune = 'a' // rune é um alias para int32, usado para representar caracteres Unicode
	fmt.Println("O rune é", reflect.TypeOf(i2))

	// TIPOS REAIS (float32, float64)
	var x float32 = 49.99
	fmt.Println("\nO tipo de x é", reflect.TypeOf(x))
	fmt.Println("Tipo literal é", reflect.TypeOf(49.99)) // float64

	// TIPOS BOOLEANOS
	var boo bool = true
	fmt.Println("\nO tipo booleano é", reflect.TypeOf(boo))

	// TIPOS ‘STRING’
	s1 := "Olá, mundo!"
	fmt.Println("Tamanho da string s1 é", len(s1))

	// STRING MULTILINHA
	s2 := `Olá, 
	mundo!`

	fmt.Println("\nMultilinha s2 é", s2)
	fmt.Println("Tamanho da string s2 é", len(s1))

	// char ???
	char := 'A' // Representa um caractere, que é um valor do tipo rune (int32)
	fmt.Println("\nO tipo do char é", reflect.TypeOf(char))

}
