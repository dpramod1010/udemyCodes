package main

import "log"

type mystruct struct {
	firstname string
}

func main() {

	var val1 mystruct
	val1.firstname = "ABC"
	log.Println("Val 1 is:- ", val1)

	val2 := mystruct{
		firstname: "pramod",
	}

	log.Println("Val1 is:-", val2)

}

// func (s *mystruct) printfirstname() string {
// 	return s.firstname
// }
