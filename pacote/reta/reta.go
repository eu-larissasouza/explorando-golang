package main

import "math"

/*
	AULA - PACOTES & VISIBILIDADE

	Com Go, podemos trabalhar a questão de visibilidade,
	- O pacote é a forma que temos para modularizar o nosso código,
	de modo que um conjunto de módulos será o que compõe o nosso programa

	- Além disso, temos como criar bibliotecas (libs ou pacotes) que podem
	ser usados no código, tanto de forma local, como remota.

	- A visibilidade em Go é feita através da letra maiúscula ou minúscula.
*/

// Iniciando com letra maiúscula é PUBLICO (visível dentro e fora do pacote)!
// Iniciando com letra minúscula é PRIVADO (visível apenas dentro do pacote)!

// Exemplo:
// Ponto - publico
// ponto ou _ponto - privado

// Ponto - representa uma coordenada no plano cartesiano
type Ponto struct {
	x float64
	y float64
}

func catetos(a, b Ponto) (cx, cy float64) {
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)
	return
}

// Distância é responsável por calcular a distância linear entre dois pontos
func Distancia(a, b Ponto) float64 {
	cx, cy := catetos(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}
