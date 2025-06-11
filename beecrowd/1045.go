package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64
	fmt.Scan(&a, &b, &c)

	if a < b {
		a, b = b, a
	}
	if b < c {
		b, c = c, b
	}
	if a < c {
		a, c = c, a
	}
	if a < b {
		a, b = b, a
	}
	
	if a >= b+c {
		fmt.Println("NAO FORMA TRIANGULO")
	} else {
		if math.Pow(a, 2) == math.Pow(b, 2)+math.Pow(c, 2) {
			fmt.Println("TRIANGULO RETANGULO")
		}
		if math.Pow(a, 2) > math.Pow(b, 2)+math.Pow(c, 2) {
			fmt.Println("TRIANGULO OBTUSANGULO")
		}
		if math.Pow(a, 2) < math.Pow(b, 2)+math.Pow(c, 2) {
			fmt.Println("TRIANGULO ACUTANGULO")
		}
		if (a == b) && (b == c) && (a == c) {
			fmt.Println("TRIANGULO EQUILATERO")
		} else if (a == b) != (b == c) != (a == c) {
			fmt.Println("TRIANGULO ISOSCELES")
		}

	}

}
