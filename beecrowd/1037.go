package main

import "fmt"

func main() {
	var a float64
	fmt.Scan(&a)

	if a >= 0 && a <= 25 {
		fmt.Println("Intervalo [0,25]")
	} else if a > 25 && a <= 50 {
		fmt.Println("Intervalo (25,50]")
	} else if a > 50 && a <= 75 {
		fmt.Println("Intervalo (50,75]")
	} else if a > 75 && a <= 100 {
		fmt.Println("Intervalo (75,100]")
	} else {
		fmt.Println("Fora de intervalo")
	}

}
