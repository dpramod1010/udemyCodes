package main

import (
	"fmt"
	"net/http"
)

var portnum = ":8081"

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is home Page:-")
}

func About(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "This Next PAge:-")
	sum := Addval(2, 2)
	fmt.Fprintf(w, " Addition  is:-%v", sum)
}

func Addval(a int, b int) int {
	return a + b
}

func main() {
	fmt.Println("Program Is running on server 8081")
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	_ = http.ListenAndServe(portnum, nil)
	fmt.Println("Program Is running on server %v", portnum)
}
