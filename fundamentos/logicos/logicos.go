package main

import "fmt"

// AULA - OPERADORES LÓGICOS
// Operadores binários, pois envolvem dois operandos

// Parametros em parenteses
// Retornos entre parenteses
func compras(trab1, trab2 bool) (bool, bool, bool) {
	// Operador lógico AND
	// Retorna true se os dois valores forem true
	comprarTv50 := trab1 && trab2

	// Operador lógico XOR
	// Retorna true se e somente se um dos operandos for true e o outro false
	comprarTv32 := trab1 != trab2

	// Operador lógico OR
	// Retorna true se pelo menos um dos operandos for true
	comprarSorvete := trab1 || trab2

	return comprarTv50, comprarTv32, comprarSorvete
}

func main() {
	// Saudável se não tiver sorvete
	tv50, tv32, sorvete := compras(false, true)
	fmt.Printf("Tv50: %t, Tv32: %t, Sorvete: %t, Saudavel: %t\n", tv50, tv32, sorvete, !sorvete)
}
