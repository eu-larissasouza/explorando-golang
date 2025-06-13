package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	if x > y {
		x, y = y, x
	}

	var soma int
	for num := x + 1; num < y; num++ {
		if num%2 != 0 {
			soma += num
		}
	}
	fmt.Println(soma)
}
