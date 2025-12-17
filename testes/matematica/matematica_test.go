package matematica

import "testing"

/*
	AULA - TESTES UNITARIOS BASICOS EM GO
	Observação: O arquivo de testes deve finalizar com _test.go

	Comandos:
	-> go test (executa os testes no pacote atual)
 	-> go test ./... (todos os testes em todos os pacotes dentro da pasta atual)
	-> go test -v (verbose) (detalha cada teste executado)
	-> go test -cover (mostra a cobertura de código pelos testes)
	-> go test -coverprofile=cobertura.out (gera um arquivo com detalhes da cobertura)
	-> go tool cover -html=cobertura.out (abre um relatório HTML da cobertura)
*/

const erroPadrao = "Valor esperado: %v, Valor retornado: %v"

// A assinatura de um teste é padronizada: Test + NomeDaFuncao
// testing.T é uma estrutura auxiliar para o teste, como um helper
func TestMedia(t *testing.T) {
	valorEsperado := 7.28
	valor := Media(7.2, 9.5, 6.5, 5.9)

	if valor != valorEsperado {
		t.Errorf(erroPadrao, valor, valorEsperado)
	}
}
