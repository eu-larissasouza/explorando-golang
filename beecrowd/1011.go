package main

import (
	"fmt"
	"math"
)

func main() {
	const pi = 3.14159
	var raio, volume float64

	fmt.Scan(&raio)

	volume = (4 / 3.0) * pi * math.Pow(raio, 3)
	fmt.Printf("VOLUME = %.3f\n", volume)
}
