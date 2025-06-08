package main

import "fmt"

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)

	diferenca := (a * b) - (c * d)
	fmt.Println("DIFERENCA =", diferenca)
}
