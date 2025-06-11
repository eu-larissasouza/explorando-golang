package main

import "fmt"

func main() {

	var cod1, cod2, qtd1, qtd2 int
	var val1, val2 float64

	fmt.Scanf("%d %d %f", &cod1, &qtd1, &val1)
	fmt.Scanf("%d %d %f", &cod2, &qtd2, &val2)

	total := float64(qtd1)*val1 + float64(qtd2)*val2

	fmt.Printf("VALOR A PAGAR: R$ %.2f\n", total)
}
