package main

import (
	"encoding/json"
	"fmt"
)

/*
	AULA - CONVERTENDO STRUCT EM JSON E VICE VERSA

	É bastante conveniente trabalhar com mapeamento de dados em JSON,
	principalmente em aplicações web, para apis REST, por exemplo.

*/

type produto struct {
	// fazemos o mapeamento dos nomes dos campos
	// há uma convenção para structs json, com os campos iniciando com letra maiúscula
	Id    int      `json:"id"`
	Nome  string   `json:"nome"`
	Preco int      `json:"preco"`
	Tags  []string `json:"tags"`
}

func main() {
	// struct -> json
	p1 := produto{1, "Caneta", 2, []string{"Papelaria", "Escrita"}}
	p1Json, _ := json.Marshal(p1) // segundo parametro é o erro

	fmt.Println(string(p1Json))

	// json -> struct
	var p2 produto
	jsonString := `{"id":2,"nome":"Lapis","preco":3,"tags":["Papelaria","Desenho"]}`

	// slice de bytes - para ele conseguir interpretar o json
	// referencia para p2
	json.Unmarshal([]byte(jsonString), &p2)
	fmt.Println(p2.Tags[1])

}
