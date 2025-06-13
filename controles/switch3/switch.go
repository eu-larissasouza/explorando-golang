package main

import (
	"fmt"
	"time"
)

// AULA - SWITCH #3

// Tipo interface{} - permite trabalhar com tipos genéricos
func tipo(i interface{}) string {
	switch i.(type) {
	case int:
		return "inteiro"
	case float32, float64:
		return "real"
	case string:
		return "string"
	case bool:
		return "boolean"
	case func():
		return "função"
	default:
		return "unknown"
	}
}

func main() {
	fmt.Println(tipo(2.3))
	fmt.Println(tipo(func() {}))
	fmt.Println(tipo(false))
	fmt.Println(tipo(time.Now()))
}
