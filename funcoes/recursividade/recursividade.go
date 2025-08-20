package main

import "fmt"

/*
	AULA - RECURSIVIDADE

	- Função que chama a si mesma
	- Importante: deve ter uma condição de parada muito clara para evitar um stack overflow
	Obs: Stack overflow é um estouro de pilha, quando a pilha de execução não tem mais espaço alocado para executar a função
*/

// Muito comum ter funções que retornam um valor e um erro
func fatorial(n int) (int, error) {
	switch {
	// não pode haver fatorial de número negativo
	case n < 0:
		return -1, fmt.Errorf("número inválido: %d", n)
	// caso base
	case n == 0:
		return 1, nil
	// caso recursivo
	default:
		fatorialAnterior, _ := fatorial(n - 1)
		return n * fatorialAnterior, nil
	}
}

func main() {
	resultado, _ := fatorial(6)
	fmt.Println(resultado)

	if _, err := fatorial(-4); err != nil {
		fmt.Println(err)
	}

}
