package main

import (
	"fmt"
)

type calculator interface {
	add() float64
	mul() float64
}

type data struct {
	first, second float64
}

func (d data) add() float64 {
	return d.first + d.second
}
func (d data) mul() float64 {
	return d.first * d.second
}

func measure(c calculator) {

	fmt.Println(c.add())
	fmt.Println(c.mul())
}

func main() {
	var input, input1 int
	fmt.Println("Enter Two Numbers:")
	fmt.Scan(&input)
	fmt.Scan(&input1)
	d := data{first: float64(input), second: float64(input1)}

	measure(d)

}
