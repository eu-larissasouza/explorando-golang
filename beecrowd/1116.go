package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		var x, y int
		fmt.Scan(&x, &y)

		if y == 0 {
			fmt.Println("divisao impossivel")
		} else {
			fmt.Printf("%.1f\n", float64(x)/float64(y))
		}
	}
}
