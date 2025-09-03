package main

import "strings"

/*
	AULA - MÉTODOS EM STRUCTS
*/

type Pessoa struct {
	nome      string
	sobrenome string
}

func (p Pessoa) getNomeCompleto() string {
	return p.nome + " " + p.sobrenome
}

func (p *Pessoa) setNomeCompleto(nomeCompleto string) {
	partes := strings.Split(nomeCompleto, " ")
	p.nome = partes[0]
	p.sobrenome = strings.Join(partes[1:], " ")
}

func main() {
	p1 := Pessoa{"João", "Silva"}
	println(p1.getNomeCompleto())

	p1.setNomeCompleto("Maria Pereira Souza")
	println(p1.getNomeCompleto())
}
