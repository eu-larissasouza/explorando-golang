package main

// BONUS - CPUS DISPONÍVEIS

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.NumCPU())
}
