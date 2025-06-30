package main

import "fmt"

// AULA - TRABALHANDO COM MAPS #1

func main() {
	// É uma estrutura de dados chave-valor que garante unicidade (não aceita repetição)
	// Possui um tipo homogeneo na chave e tipo homogeneo no valor

	// var aprovados map[int]string
	// maps devem ser inicializados sempre
	aprovados := make(map[int]string)

	aprovados[12345678978] = "Maria"
	aprovados[98765432100] = "Pedro"
	aprovados[95135745682] = "Ana"
	fmt.Println(aprovados)

	// percorre a chave e o valor
	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	// lemos o dado usando a chave
	fmt.Println(aprovados[95135745682])

	// delete permite excluir um valor
	delete(aprovados, 95135745682)
	fmt.Println(aprovados[95135745682], "- Elemento excluido")
}
