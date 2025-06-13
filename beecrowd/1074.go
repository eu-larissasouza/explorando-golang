package main

import "fmt"

func main() {
	var n, valor int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&valor)
		switch {
		case valor == 0:
			fmt.Println("NULL")
		case valor%2 == 0:
			fmt.Printf("EVEN ")
		case valor%2 != 0:
			fmt.Printf("ODD ")
		}

		switch {
		case valor > 0:
			fmt.Printf("POSITIVE\n")
		case valor < 0:
			fmt.Printf("NEGATIVE\n")
		}
	}
}
