package main

import "fmt"

/*
	AULA - STRUCTS
	São tipos de dados que permitem armazenar diferentes valores em uma mesma variável
*/

type Produto struct {
	nome     string
	preco    float64
	desconto float64
}

// Metodo: função com receiver (receptor)
// receiver é o parâmetro que define em qual struct a função irá operar
func (p Produto) precoComDesconto() float64 {
	return p.preco * (1 - p.desconto)
}

func main() {
	var produto1 Produto
	produto1 = Produto{
		nome:     "Lápis",
		preco:    1.80,
		desconto: 0.05,
	}
	fmt.Println(produto1, produto1.precoComDesconto())

	// outra forma de declarar uma struct
	produto2 := Produto{"Notebook", 1989.90, 0.10}
	fmt.Println(produto2.precoComDesconto())
}
