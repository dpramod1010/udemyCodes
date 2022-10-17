package main

import "fmt"

func main() {
	fmt.Println("Odd Even Program...")
	fmt.Println("Enter Number")
	var num int

	fmt.Scan(&num)
	if num%2 == 0 {
		fmt.Println("Odd Number")
	} else {
		fmt.Println("Even NUmber")
	}

}
