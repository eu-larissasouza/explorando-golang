package main

import "fmt"

func main() {
	var n1, n2 float64
	fmt.Scan(&n1, &n2)

	media := (n1*3.5 + n2*7.5) / 11

	fmt.Printf("MEDIA = %.5f\n", media)
}
