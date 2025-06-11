package main

// AULA - IF INIT

import (
	"fmt"
	"math/rand"
	"time"
)

func numeroAleatorio() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(10)
}

func main() {
	// Bloco de inicialização
	// -- Define uma variavel que vai ser utilizada somente no contexto do bloco condicional
	// -- Também suportado no switch (case) e for
	if i := numeroAleatorio(); i > 5 {
		fmt.Println("Ganhou!")
	} else {
		fmt.Println("Perdeu!")
	}

}
