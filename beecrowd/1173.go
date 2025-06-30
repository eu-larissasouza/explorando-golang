package main

import "fmt"

func main() {
	var N [10]int

	fmt.Scanf("%d", &N[0])

	for i := 1; i < 10; i++ {
		N[i] = N[i-1] * 2
	}

	for idx, num := range N {
		fmt.Printf("N[%d] = %d\n", idx, num)
	}
}
