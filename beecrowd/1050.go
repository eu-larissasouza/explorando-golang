package main

import "fmt"

func main() {
	var ddd int
	fmt.Scan(&ddd)

	switch ddd {
	case 11:
		fmt.Println("Sao Paulo")
	case 19:
		fmt.Println("Campinas")
	case 21:
		fmt.Println("Rio de Janeiro")
	case 27:
		fmt.Println("Vitoria")
	case 31:
		fmt.Println("Belo Horizonte")
	case 32:
		fmt.Println("Juiz de Fora")
	case 61:
		fmt.Println("Brasilia")
	case 71:
		fmt.Println("Salvador")
	default:
		fmt.Println("DDD nao cadastrado")
	}

}
