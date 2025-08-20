package main

import "fmt"

/*
	AULA - DEFER
	- É uma função que atrasa a execução de uma sentença de código até o momento em que a função atual terminará
	- Muito comum para fechar conexões de recursos abertos, arquivos, etc
	- O defer é executado mesmo que ocorra um erro na função
	- Garante que a sentença de código será executada no momento mais tardio possível, antes do return da função
*/

func obterValorAprovado(num int) int {
	defer fmt.Println("fim da execução!")
	if num > 500 {
		fmt.Println("Valor alto...")
		return 500
	} else {
		fmt.Println("Valor baixo...")
		return num
	}
}

func main() {
	fmt.Println(obterValorAprovado(6))
}
