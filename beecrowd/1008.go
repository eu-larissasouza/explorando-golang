package main

import "fmt"

func main() {
	var num, horas int
	var valorHora, salario float64

	fmt.Scan(&num, &horas, &valorHora)

	salario = float64(horas) * valorHora
	fmt.Printf("NUMBER = %d\n", num)
	fmt.Printf("SALARY = U$ %.2f\n", salario)
}
