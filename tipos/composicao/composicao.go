package main

import "fmt"

/*
	AULA - COMPOSIÇÃO DE INTERFACES
	É muito comum utilizar uma interface que seja composta por outras interfaces.
*/

type esportivo interface {
	ligarTurbo()
}
type luxuoso interface {
	fazerBaliza()
}
type esportivoLuxuoso interface {
	esportivo
	luxuoso
	// adicionar outros metodos
}

type bmw7 struct{}

func (b *bmw7) ligarTurbo() {
	fmt.Println("Turbo...")
}
func (b *bmw7) fazerBaliza() {
	fmt.Println("Baliza...")
}

// OBS: se eu adicionar um novo metodo na bmw, ela deixa de ser uma esportivoLuxuoso
// pois não implementa mais todos os metodos da interface
