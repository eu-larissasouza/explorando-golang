package matematica

import (
	"fmt"
	"strconv"
)

/*
	AULA - TESTES UNITARIOS BASICOS EM GO

	Não posso executar um codigo sem um pacote main, mas posso testar e executar o código atraves de testes

	O objetivo do teste unitário é garantir que uma unidade de código (função, método, etc.) funcione conforme o esperado.

	Uma outra pratica é construir o teste antes do código, conhecido como TDD (Test-Driven Development).
	-> Isso garante que seus testes tenham estresse suficiente para validar o comportamento
*/

func Media(numeros ...float64) float64 {
	total := 0.0
	for _, num := range numeros {
		total += num
	}
	media := total / float64(len(numeros))
	mediaArredondada, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", media), 64)
	return mediaArredondada
}
