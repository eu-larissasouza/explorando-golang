package main

import "fmt"

func main() {
	var X [10]int

	for i := 0; i < len(X); i++ {
		fmt.Scanf("%d", &X[i])
		if X[i] <= 0 {
			X[i] = 1
		}
	}

	for i, x := range X {
		fmt.Printf("X[%d] = %d\n", i, x)
	}

}
