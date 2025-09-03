package main

import "fmt"

/*
	AULA - USANDO INTERFACES #01
	-> Esse é um conceito de OO que Golang implementa, mas de uma forma diferente.
	-> Uma interface é um contrato que define um conjunto de metodos que um tipo deve implementar.

	-> OBS: Interfaces são implementadas implicitamente, ou seja, não é necessario declarar que um tipo implementa uma interface.
	-> Quando eu crio todos os metodos que pertence a uma interface, eu ja estou implementando essa interface.
	-> Uma interface pode ser usada como um tipo,
*/

type imprimivel interface {
	// define o metodo que toda interface deve ter
	toString() string
}

type pessoa struct {
	nome      string
	sobrenome string
}

type produto struct {
	nome  string
	preco float64
}

func (p pessoa) toString() string {
	return p.nome + " " + p.sobrenome
}

func (p produto) toString() string {
	return fmt.Sprintf("%s: %.2f", p.nome, p.preco)
}

func imprimir(x imprimivel) {
	fmt.Println(x.toString())
}

func main() {
	var coisa imprimivel = pessoa{"Roberto", "Soares"}
	imprimir(coisa)

	// a variavel pode receber qualquer tipo que implemente a interface imprimivel
	coisa = produto{"Celular", 1490.00}
	fmt.Println(coisa.toString())

}
