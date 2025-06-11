package main

import "fmt"

// EM GO, NÃO TEMOS OPERADOR TERNÁRIO
// A ALTERNATIVA SERIA USAR O PRÓPRIO IF

func obterResultado(nota float64) string {
	if nota >= 6 {
		return "Aprovado"
	}
	return "Reprovado"
}

func main() {
	fmt.Println(obterResultado(6.2))
}
