package main

import "fmt"

// AULA - MAPS ANINHADOS

func main() {
	// Tenho um map que tem uma string como chave e que tem como valor um outro map
	// Agrupando funcionarios pela inicial do nome
	// Cuidado ao utilizar esse tipo de estrutura.
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela Silva": 15456.78,
			"Guga Pereira":   8456.78,
		},
		"J": {
			"José João": 11325.45,
		},
		"P": {
			"Pedro Junior": 1200.0,
		},
	}

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
	}
}
