package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)

	for i := 2; i <= n; i += 2 {
		fmt.Printf("%d^%d = %.0f\n", i, 2, math.Pow(float64(i), 2))
	}
}
