package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)

	for x != 0 {
		for i := 1; i < x; i++ {
			fmt.Printf("%d ", i)
		}
		fmt.Printf("%d\n", x)

		fmt.Scan(&x)
	}
}
