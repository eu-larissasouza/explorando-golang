package main

import "fmt"

/*
	AULA - USANDO INTERFACES #02
*/

type esportivo interface {
	ligarTurbo()
}

type ferrari struct {
	modelo          string
	turboLigado     bool
	velocidadeAtual int
}

func (f *ferrari) ligarTurbo() {
	f.turboLigado = true
}

func main() {
	carro1 := ferrari{
		modelo:          "F40",
		turboLigado:     false,
		velocidadeAtual: 0,
	}
	carro1.ligarTurbo()

	// Quando usamod o nivel de interface, precisamos passar o endereço da struct,
	// para que o metodo possa alterar o valor do atributo

	// OBS: É mais comum que se evite trabalhar com metodos que geram efeitos colaterais ao usar interfaces
	var carro2 esportivo = &ferrari{"F40", false, 10}
	carro2.ligarTurbo()

	fmt.Println(carro1, carro2)
}
