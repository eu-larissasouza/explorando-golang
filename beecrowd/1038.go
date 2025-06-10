package main

import "fmt"

func main() {
	var item, qtd int
	var total float64

	fmt.Scan(&item, &qtd)

	switch item {
	case 1:
		total = float64(qtd) * 4
	case 2:
		total = float64(qtd) * 4.50
	case 3:
		total = float64(qtd) * 5
	case 4:
		total = float64(qtd) * 2
	case 5:
		total = float64(qtd) * 1.50
	}
	fmt.Printf("Total: R$ %.2f\n", total)

}
