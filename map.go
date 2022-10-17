// package main

// import "log"

// type user struct {
// 	firstname string
// 	lastname  string
// }

// func main() {
// 	// log.Println("Map Praogram:--")
// 	// m := make(map[int]int)
// 	// m[1] = 10
// 	// m[2] = 20
// 	// log.Println("Prints First Value", m[1])
// 	// log.Println("Prints All Value From map:- ", m)

// 	//map using Struct

// 	m := make(map[string]user)

// 	me := user{
// 		firstname: "pramod",
// 		lastname:  "Dhayarkar",
// 	}

// 	m["me"] = me

// 	log.Println("Firstname", m["me"].firstname)

// }

package main

import (
	"fmt"
)

type user struct {
	firstname string
	lastname  string
}

func main() {
	m := make(map[string]user)

	me := user{
		firstname: "pranod",
		lastname:  "dhayarkar",
	}
	m["me"] = me
	//var n string
	// fmt.Scan("Enter What You want", &n)

	fmt.Println("firstname is:- ", m["me"].firstname)
	fmt.Println("All value is:- ", m["me"])

	// fmt.Println("firstname is:- ", m[""].n)

}
