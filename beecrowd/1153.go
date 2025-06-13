package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	fatorial := 1
	for i := N; i > 0; i-- {
		fatorial *= i
	}
	fmt.Println(fatorial)
}
