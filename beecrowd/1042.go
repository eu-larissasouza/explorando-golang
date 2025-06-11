package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	d, e, f := a, b, c

	if e < d {
		d, e = e, d
	}
	if f < e {
		f, e = e, f
	}
	if f < d {
		d, f = f, d
	}
	if e < d {
		d, e = e, d
	}

	fmt.Printf("%d\n%d\n%d\n\n", d, e, f)
	fmt.Printf("%d\n%d\n%d\n", a, b, c)
}
