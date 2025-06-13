package main

import (
	"fmt"
	"strings"
)

func main() {
	var nCasos int
	fmt.Scan(&nCasos)

	var C, R, S int
	for i := 0; i < nCasos; i++ {
		var quantia int
		var tipo string
		fmt.Scanf("%d %s", &quantia, &tipo)

		tipo = strings.ToUpper(tipo)
		switch tipo {
		case "C":
			C += quantia
		case "R":
			R += quantia
		case "S":
			S += quantia
		}
	}
	total := C + R + S

	fmt.Println("Total:", total, "cobaias")
	fmt.Println("Total de coelhos:", C)
	fmt.Println("Total de ratos:", R)
	fmt.Println("Total de sapos:", S)
	fmt.Printf("Percentual de coelhos: %.2f %%\n", float64(C)/float64(total)*100)
	fmt.Printf("Percentual de ratos: %.2f %%\n", float64(R)/float64(total)*100)
	fmt.Printf("Percentual de sapos: %.2f %%\n", float64(S)/float64(total)*100)
}
