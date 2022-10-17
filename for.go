package main

import "fmt"

type User struct {
	firstname string
	no        int
}

func main() {

	var n int
	fmt.Println("How many to u want to add")
	fmt.Scan(&n)
	var animal []string
	var nanimal string
	//var i int
	for i := 0; i <= n; i++ {

		fmt.Println("Enter Animal name")
		fmt.Scan(&nanimal)

		animal = append(animal, nanimal)

	}
	for i, animal := range animal {
		fmt.Println("Animal List is:- ", i, animal)

	}
	// created one sclice and struct
	var user []User

	user = append(user, User{firstname: "pramod", no: 23})
	for _, l := range user {
		fmt.Println(l.firstname)
	}

}
