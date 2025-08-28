package main

import "fmt"

/*
	AULA - STRUCT ANINHADA
*/

type item struct {
	produtoID int
	qtde      int
	preco     float64
}

type pedido struct {
	userID int
	itens  []item // Um pedido pode ter vários itens
}

// receiver p do tipo pedido
func (p pedido) valorTotal() float64 {
	total := 0.0
	for _, item := range p.itens {
		total += item.preco * float64(item.qtde)
	}
	return total
}

func main() {
	// criando um pedido com 3 itens
	// em structs, prefira nomear os campos, pois assim não precisa se preocupar com a ordem e o código fica mais legível
	pedido := pedido{
		userID: 1,
		itens: []item{
			{1, 2, 12.1},
			{2, 1, 23.49},
			{11, 100, 3.13},
		},
	}
	fmt.Printf("Valor total do pedido: R$ %.2f\n", pedido.valorTotal())
}
