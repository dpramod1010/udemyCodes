package main

import "fmt"

func main() {
	var n int
	fmt.Println("Enter Number:-")
	fmt.Scan(&n)
	switch n {
	case 1:
		fmt.Println("You Entered One")
	case 2:
		fmt.Println("You Entered Two")
	case 3:
		fmt.Println("You Entered Three")
	case 4:
		fmt.Println("You Entered Four")
	case 5:
		fmt.Println("You Entered Five")
	default:
		fmt.Println("Please Enter Number Between 1 to 5")
	}

}
