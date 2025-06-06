package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2

	fmt.Println("Soma:", a+b)
	fmt.Println("Subtração:", a-b)
	fmt.Println("Multiplicação:", a*b)
	fmt.Println("Divisão:", a/b)
	fmt.Println("Módulo:", a%b)

	// Bit a bit
	// Operações bit a bit são realizadas em números inteiros
	fmt.Println("\nAND bit a bit:", a&b) // 11 & 10 = 10
	fmt.Println("OR bit a bit:", a|b)    // 11 | 10 = 11
	fmt.Println("XOR bit a bit:", a^b)   // 11 ^ 10 = 01

	c := 3.0
	d := 2.0

	// operações usando math
	fmt.Println("\nMaior:", math.Max(c, d))
	fmt.Println("Menor:", math.Min(c, d))
	fmt.Println("Exponenciação:", math.Pow(c, d)) // 3^2 = 9
}
