package main

import "fmt"

func main() {
	var valor float64
	fmt.Scan(&valor)

	switch {
	case valor >= 0 && valor <= 25:
		fmt.Println("Intervalo [0,25]")
	case valor > 25 && valor <= 50:
		fmt.Println("Intervalo (25,50]")
	case valor > 50 && valor <= 75:
		fmt.Println("Intervalo (50,75]")
	case valor > 75 && valor <= 100:
		fmt.Println("Intervalo (75,100]")
	default:
		fmt.Println("Fora de intervalo")
	}
}
