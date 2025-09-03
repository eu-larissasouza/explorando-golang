package main

import "fmt"

/*
	AULA - PSEUDO-HERANÇA EM GOLANGx
*/

type carro struct {
	nome            string
	velocidadeAtual int
}

// o type ferrari "herda" os campos e métodos do type carro
// isso é chamado de pseudo-herança, pois, na prática, é uma composição de structs
type ferrari struct {
	carro      // campo anonimo do tipo carro e o que há em carro, estará disponível em ferrari
	tudoLigado bool
}

func main() {
	f := ferrari{}
	f.nome = "F40"
	f.velocidadeAtual = 0
	f.tudoLigado = true

	fmt.Printf("A ferrari %s está com tudo ligado? %v\n", f.nome, f.tudoLigado)
	fmt.Println(f)
}
