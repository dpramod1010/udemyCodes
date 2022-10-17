package main

import "fmt"

func main() {

	a := 3

	sa(a)
	ad(&a)
}

func sa(v int) {
	fmt.Println(&v)
	v *= v
	fmt.Println(&v, v)
}

func ad(a *int) {
	*a *= *a
	fmt.Println(&a, *a)
}
