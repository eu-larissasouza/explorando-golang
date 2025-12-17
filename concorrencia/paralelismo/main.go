/*
	AULA - CONCORRÊNCIA VS PARALELISMO

	É um tema central na linguagem Golang e foi uma das linguagens desenvolvidas após o upgrade de computadores com múltiplos Cores em processadores

	Paralelismo diz respeito a fazer duas tarefas exatamente ao mesmo tempo, só pode ser feito quando temos múltiplos processadores

	Concorrência é a capacidade de administrar múltiplas tarefas, portanto permite que seja feita dentro de um mesmo processador, como é o caso de escalonadores
	- Em um computador com apenas um core

	Paralelismo é mais custoso, pois existe um overhead de controle de tarefas e, isso ocorre para que todos os dados sejam recebidos e unidos ao trabalhar com múltiplos processadores.
	———————————
	Golang é uma linguagem concorrente

	Paralelismo - executar código simultaneamente em processadores físicos diferentes

	Concorrência - ato de intercalar (administrar) vários processos ao mesmo tempo

	Concorrência viabiliza o paralelismo, porém é possível que a concorrência seja melhor em determinadas circunstâncias.

	Concorrência - é a forma de estruturar o seu programa.
	É a composição de processos (tipicamente funções) que executam de forma independente.

*/

package main
