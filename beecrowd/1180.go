package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	X := make([]int, N)

	var men, pos int
	for i := 0; i < N; i++ {
		fmt.Scan(&X[i])
		if i == 0 {
			men = X[i]
			pos = i
		} else if men > X[i] {
			men = X[i]
			pos = i
		}
	}

	fmt.Println("Menor valor:", men)
	fmt.Println("Posição:", pos)

}
