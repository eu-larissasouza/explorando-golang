package main

import (
	"fmt"
	"math"
)

func main() {
	const pi = 3.14159
	var a, b, c float64
	fmt.Scanf("%f %f %f", &a, &b, &c)

	tri := (a * c) / 2
	cir := pi * math.Pow(c, 2)
	tra := (a + b) * c / 2
	qua := math.Pow(b, 2)
	ret := a * b

	fmt.Printf("TRIANGULO: %.3f\n", tri)
	fmt.Printf("CIRCULO: %.3f\n", cir)
	fmt.Printf("TRAPEZIO: %.3f\n", tra)
	fmt.Printf("QUADRADO: %.3f\n", qua)
	fmt.Printf("RETANGULO: %.3f\n", ret)

}
