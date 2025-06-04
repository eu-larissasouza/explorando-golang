/*
	AULA - Primeiro Programa em Go

 	Convenção: projetos em Go rodam em um pacote
 	Para ser executável, o pacote precisa se chamar main

 	Pacote é uma forma de organizar o código, separando em módulos que podem ser reutilizados
 	-> Também define o escopo de visibilidade das variáveis, funções e tipos

	COMENTÁRIOS:
	--	Priorize código legível e faça comentários que agreguem valor ao entendimento do código!
	-- 	Evite comentários óbvios ou redundantes

*/

package main

/*
	Importa pacotes necessários para o programa
	"fmt" é um pacote padrão do Go que fornece funções para IOs formatados
*/
import "fmt"

// Função principal do programa
func main() {
	// Ao salvar, o format GO deixa o estilo do código padronizado.
	// Suporta ponto e virgula, mas não é obrigatório
	fmt.Println("Primeiro")
	fmt.Println("Programa em Go")
}
