package main

// AULA - ZEROS EM GO (VALORES PADRÕES)

import "fmt"

func main() {
	var a int
	var b float64
	var c string
	var d bool
	var e *int // ponteiro para int

	// %q é usado para formatar strings com aspas
	fmt.Printf("%v %v %q %v %v", a, b, c, d, e)

}
