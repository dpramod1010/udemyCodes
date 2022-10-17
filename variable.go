package main

import "fmt"

func main() {
	fmt.Println("Second Code")
	var i int
	i = 20
	fmt.Println("Value Of i is:- ", i)
	var a int
	var b int
	fmt.Println("Enter value of a")
	fmt.Scan(&a)
	fmt.Println("Enter value of b")
	fmt.Scan(&b)
	n := add(a, b)
	m := mul(a, b)
	fmt.Println("Addition of two number is:- ", n)
	fmt.Println("Multiplication is:- ", m)
}
func add(a int, b int) int {
	return a + b
}
func mul(a int, b int) int {
	return a * b
}
