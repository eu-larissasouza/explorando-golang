package main

import "fmt"

func main() {
	var salario, imposto float64
	fmt.Scan(&salario)

	switch {
	case salario >= 0 && salario <= 2000:
		fmt.Println("Isento")
	case salario > 2000 && salario <= 3000:
		imposto = (salario - 2000) * 0.08
		fmt.Printf("R$ %.2f\n", imposto)
	case salario > 3000 && salario <= 4500:
		imposto = (salario-3000)*0.18 + 1000*0.08
		fmt.Printf("R$ %.2f\n", imposto)
	case salario > 4500:
		imposto = (salario-4500)*0.28 + 1500*0.18 + 1000*0.08
		fmt.Printf("R$ %.2f\n", imposto)
	}
}
