// Go program to illustrate the
// concept of appending slices.
package main

import "fmt"

func main() {
	var sclice1 []int
	sclice1 = append(sclice1, 1, 2, 3, 4)

	// Displaying slice
	fmt.Println("Slice1 : ", sclice1)

	// Appending slice
	var sclice2 []int
	sclice2 = append(sclice2, 5)
	sclice2 = append(sclice2, 6)
	sclice2 = append(sclice2, sclice1...)

	// Displaying slice

	fmt.Println("Slice2 : ", sclice2)
}
