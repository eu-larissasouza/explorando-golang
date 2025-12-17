package tabela

import (
	"strings"
	"testing"
)

// AULA - CRIANDO UM DATASET PARA TESTES
// Vamos testar a função index do pacote string do Go

// Criamos uma tabela que permite diferentes entradas e saídas esperadas para o mesmo teste
// -> Cenários positivos e negativos são testados

const msgIndex = "%s (parte: %s) - índices: esperado (%d) <> encontrado (%d)."

func TestIndex(t *testing.T) {
	t.Parallel()

	testes := []struct {
		texto    string
		parte    string
		esperado int
	}{
		{
			texto:    "Google Chrome é um navegador.",
			parte:    "Google",
			esperado: 0,
		},
		{
			texto:    "",
			parte:    "",
			esperado: 0,
		},
		{
			texto:    "",
			parte:    "opa",
			esperado: -1,
		},
		{
			texto:    "teste",
			parte:    "s",
			esperado: 2,
		},
	}

	for _, test := range testes {
		atual := strings.Index(test.texto, test.parte)

		if atual != test.esperado {
			t.Errorf(msgIndex, test.texto, test.parte, test.esperado, atual)
		}
	}
}
