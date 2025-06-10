package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64
	fmt.Scanf("%f %f %f", &a, &b, &c)

	delta := math.Pow(b, 2) - (4 * a * c)

	if delta < 0 || a == 0 {
		fmt.Println("Impossivel calcular")
	} else {
		r1 := ((-b) + math.Sqrt(delta)) / (2 * a)
		r2 := ((-b) - math.Sqrt(delta)) / (2 * a)

		fmt.Printf("R1 = %.5f\n", r1)
		fmt.Printf("R2 = %.5f\n", r2)
	}

}
