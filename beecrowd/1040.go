package main

import "fmt"

func main() {
	var n1, n2, n3, n4 float64
	fmt.Scan(&n1, &n2, &n3, &n4)

	media := (n1*2 + n2*3 + n3*4 + n4*1) / 10
	fmt.Printf("Media: %.1f\n", media)

	if media >= 7.0 {
		fmt.Println("Aluno aprovado.")
	} else if media < 5.0 {
		fmt.Println("Aluno reprovado.")
	} else {
		fmt.Println("Aluno em exame.")
		var nExame float64
		fmt.Scan(&nExame)

		fmt.Printf("Nota do exame: %.1f\n", nExame)
		media = (media + nExame) / 2.0

		if media >= 5.0 {
			fmt.Println("Aluno aprovado.")
		} else {
			fmt.Println("Aluno reprovado.")
		}

		fmt.Printf("Media final: %.1f\n", media)
	}
}
