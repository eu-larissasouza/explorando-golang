package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var contIn, contOut int

	for i := 0; i < n; i++ {
		var valor int
		fmt.Scan(&valor)

		if valor >= 10 && valor <= 20 {
			contIn++
		} else {
			contOut++
		}
	}
	fmt.Println(contIn, "in")
	fmt.Println(contOut, "out")
}
