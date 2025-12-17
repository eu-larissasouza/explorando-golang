package main

import (
	"fmt"
	"time"
)

// AULA - ESTRUTURA DE CONTROLE SELECT

/*
	Temos essa estrutura de controle para trabalhar com múltiplos canais de forma concorrente.

	Ela é semelhante ao switch, mas cada case dentro do select deve ser uma operação de leitura ou escrita em um canal.

	-> Permite que uma goroutine aguarde múltiplas operações de comunicação.
	-> A função select fica bloqueada até que um de seus casos possa ser executado, então ela executa esse caso.
	Se houver vários prontos, ela escolhe um aleatoriamente.

*/

func oMaisRapido(url1, url2, url3 string) string {
	c1 := titulo(url1)
	c2 := titulo(url2)
	c3 := titulo(url3)

	// estrutura de controle especifica para concorrencia
	select {
	case t1 := <-c1:
		return t1
	case t2 := <-c2:
		return t2
	case t3 := <-c3:
		return t3
	case <-time.After(200 * time.Millisecond):
		return "Todos perderam"
		//default:
		//	return "Sem resposta ainda"
	}
}

func main() {
	// o primeiro resultado dentro do select que chegar será o vencedor
	campeao := oMaisRapido(
		"https://www.youtube.com",
		"https://www.github.com",
		"https://www.google.com",
	)
	fmt.Println(campeao)
}
