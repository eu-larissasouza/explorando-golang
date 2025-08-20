package main

import "fmt"

// AULA - FUNÇÕES BÁSICAS
// Sintaxe: func nomeDaFuncao (parametro tipo...) tipoRetorno {}
// função pura: não gera nenhum efeito colateral a nenhuma informação de fora
// -> ex: uma função que recebe parametros e simplesmente os utiliza, sem alterar variáveis externas
// Go permite retornar multiplos valores em uma função

func f1() {
	fmt.Println("F1")
}

func f2(p1 string, p2 string) {
	fmt.Printf("F2 %s %s\n", p1, p2)
}

func f3() string {
	return "F3"
}

// fmt.Sprintf - não exibe no console, mas retorna a string formatada
func f4(p1, p2 string) string {
	return fmt.Sprintf("F4: %s %s", p1, p2)
}

func f5() (string, string) {
	return "Retorno 1", "Retorno 2"
}

func main() {
	f1()
	f2("Param 1", "Param 2")

	r3, r4 := f3(), f4("Param 1", "Param 2")
	fmt.Println(r3, r4)

	re1, re2 := f5()
	fmt.Println(re1, re2)
}
