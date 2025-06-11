package main

import "fmt"

func main() {
	var vendedor string
	var salarioFixo, vendas float64

	fmt.Scan(&vendedor, &salarioFixo, &vendas)

	total := salarioFixo + (0.15 * vendas)

	fmt.Printf("TOTAL = R$ %.2f\n", total)
}
