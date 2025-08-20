package main

import "fmt"

/*
	AULA - FUNÇÃO INIT
	- É uma convenção do Go que permite a execução de um bloco de código antes da função main
	- Essa função init é executada automaticamente pelo Go, não precisamos chamá-la
 	- Podemos ter vários blocos init em um mesmo pacote, todos serão executados antes da função main
*/

func init() {
	fmt.Println("Inicializando...")
}

func main() {
	fmt.Println("Main...")
}
