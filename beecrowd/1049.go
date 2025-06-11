package main

import "fmt"

func main() {
	var ossos, tipo, alimentacao string
	fmt.Scan(&ossos)
	fmt.Scan(&tipo)
	fmt.Scan(&alimentacao)

	switch ossos {
	case "vertebrado":
		if tipo == "ave" {
			if alimentacao == "carnivoro" {
				fmt.Println("aguia")
			} else {
				fmt.Println("pomba")
			}
		} else {
			if alimentacao == "onivoro" {
				fmt.Println("homem")
			} else {
				fmt.Println("vaca")
			}
		}
	case "invertebrado":
		if tipo == "inseto" {
			if alimentacao == "hematofago" {
				fmt.Println("pulga")
			} else {
				fmt.Println("lagarta")
			}
		} else {
			if alimentacao == "hematofago" {
				fmt.Println("sanguessuga")
			} else {
				fmt.Println("minhoca")
			}
		}
	}
}
