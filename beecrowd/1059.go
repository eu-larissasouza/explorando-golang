package main

import "fmt"

func main() {

	for num := 2; num <= 100; num += 2 {
		if num%2 == 0 {
			fmt.Println(num)
		}
	}
}
