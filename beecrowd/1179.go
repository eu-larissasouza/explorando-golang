package main

import "fmt"

func exibirVetor(tipo string, vetor [5]int) {
	for id, num := range vetor {
		fmt.Printf("%s[%d] = %d", tipo, id, num)
	}
}

func main() {
	var par, impar [5]int
	var idxPar, idxImpar int

	for i := 0; i < 15; i++ {
		var valor int
		fmt.Scan(&valor)

		if valor%2 == 0 {
			par[idxPar] = valor
			idxPar++
		} else {
			impar[idxImpar] = valor
			idxImpar++
		}

		if idxPar == 4 {
			exibirVetor("par", par)
			idxPar = 0
		}
		if idxImpar == 4 {
			exibirVetor("impar", impar)
			idxImpar = 0
		}
	}

}
