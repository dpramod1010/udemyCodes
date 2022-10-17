package main

import "fmt"

func main() {
	var input int

	fmt.Scan(&input)
	if input == 1 {
		fmt.Println("Welcome, You Login Sucess:")
	} else if input == 2 {
		fmt.Printf("Welcome, %v You Login Sucessss:", input)
	} else {
		fmt.Println("Welcome, You Login Sucessss:")
	}

}
