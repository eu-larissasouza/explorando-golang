package main

import "fmt"

func ordenarPar(m, n *int) {
	if *m > *n {
		*m, *n = *n, *m
	}
}

func main() {
	var m, n int
	fmt.Scan(&m, &n)

	for m > 0 && n > 0 {
		ordenarPar(&m, &n)

		var sum int

		for i := m; i <= n; i++ {
			fmt.Printf("%d ", i)
			sum += i
		}
		fmt.Printf("Sum=%d\n", sum)

		fmt.Scan(&m, &n)
	}
}
