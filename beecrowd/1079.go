package main

import "fmt"

func main() {
	var nCasos int
	fmt.Scan(&nCasos)

	var n1, n2, n3 float64
	for i := 0; i < nCasos; i++ {
		fmt.Scanf("%f %f %f", &n1, &n2, &n3)

		media := (n1*2 + n2*3 + n3*5) / 10
		fmt.Printf("%.1f\n", media)
	}
}
