package main

import (
	"fmt"
	"math"
)

func ordenarValores(a, b, c *float64) {
	if *a < *b {
		*a, *b = *b, *a
	}
	if *b < *c {
		*b, *c = *c, *b
	}
	if *a < *c {
		*a, *b = *c, *a
	}
	if *a < *b {
		*a, *b = *b, *a
	}
}

func main() {
	var a, b, c float64
	fmt.Scan(&a, &b, &c)

	ordenarValores(&a, &b, &c)

	switch {
	case a >= b+c:
		fmt.Println("NAO FORMA TRIANGULO")
	case math.Pow(a, 2) == math.Pow(b, 2)+math.Pow(c, 2):
		fmt.Println("TRIANGULO RETANGULO")
	case math.Pow(a, 2) > math.Pow(b, 2)+math.Pow(c, 2):
		fmt.Println("TRIANGULO OBTUSANGULO")
	case math.Pow(a, 2) < math.Pow(b, 2)+math.Pow(c, 2):
		fmt.Println("TRIANGULO ACUTANGULO")
	}

	switch {
	case (a == b) && (a == c):
		fmt.Println("TRIANGULO EQUILATERO")
	case (a == b) != (a == c) != (b == c):
		fmt.Println("TRIANGULO ISOSCELES")
	}
}
