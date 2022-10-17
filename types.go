package main

import "log"

var s = "This is s1"

func main() {

	var a string
	a = "This is a"
	log.Println(a)
	log.Println(s)

	z, x := new(s)

	log.Println(z, x)

}

func new(s string) (string, string) {
	return s, "Hello"
}
