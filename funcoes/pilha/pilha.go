package main

import "runtime/debug"

// AULA - PILHA DE FUNÇÕES

func f3() {
	// Imprime a pilha de execução dessa função
	// Geralmente casos de erro, é bem comum que junto da mensagem de erro seja exibida a pilha de execução.
	debug.PrintStack()
}

func f2() {
	f3()
}

func f1() {
	f2()
}

// A pilha de execução funciona como uma torre de livros que vão se encadeando ou removendo
// Cada vez que uma nova função é chamada, incluimos esse metodo na pilha
// Por outro lado, quando é terminada a execução de uma função, o metodo é removido da pilha

// Conceito muito importante para a recursidade
// --> Para não haver estouro da pilha de execução
// --> Necessário ter uma condição de parada clara em problemas recursivos
// --> Linguagens possuem recursos para otimização em processos com recursividade
func main() {
	f1()

}
