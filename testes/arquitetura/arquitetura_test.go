package arquitetura

/*
	AULA - TESTES DE ARQUITETURA

	Em Golang, é possível criar testes que são específicos para determinadas arquiteturas de CPU.

	Isso é útil quando você tem código que depende de características específicas de uma arquitetura,
*/

import (
	"runtime"
	"testing"
)

func TestDependente(t *testing.T) {
	// Define que este teste pode ser executado em paralelo com outros testes
	// Cria goroutines para executar vários testes ao mesmo tempo
	t.Parallel()

	// Teste para garantir que a função dependente funciona corretamente
	t.Log("Teste de arquitetura - dependente")

	if runtime.GOARCH == "amd64" {
		t.Skip("Não funciona em arquitetura amd64")
	}
	// Aqui você pode adicionar o código do teste que deve ser executado
	t.Fail()
}
