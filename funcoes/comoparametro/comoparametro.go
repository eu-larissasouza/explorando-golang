package main

/*
	AULA - FUNÇÕES COMO PARÂMETRO

	Em essência, Golang não é uma linguagem funcional, mas possui recursos que permitem
	aplicar principios das linguagens funcionais.

	Uma função pode ser passada como parâmetro para outra função, o que permite criar
	abstrações e reutilizar código de forma eficiente.

*/

func multiplicar(a, b int) int {
	return a * b
}

// Tem que definir a assinatura da função que será passada como parâmetro
func executar(funcao func(int, int) int, p1, p2 int) int {
	return funcao(p1, p2)
}

func main() {
	resultado := executar(multiplicar, 4, 3)
	println("Resultado:", resultado) // Resultado: 12
}
