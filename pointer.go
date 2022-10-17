package main

import "log"

func main() {
	var a string
	a = "Green"
	log.Println("Colour is:- ", a)
	say(&a)
	log.Println("After Change :-", a)

}
func say(s *string) {
	var z string
	z = "Red"
	*s = z
	log.Println("Before Z set to:", *s)
}
