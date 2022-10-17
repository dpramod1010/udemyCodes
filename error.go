package main

import (
	"fmt"
	"net/http"
)

func divide(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is divide program")
	op := div(10, 2)
	fmt.Fprintf(w, "Dividetion is:%v", op)
	if op == 5 {
		fmt.Fprintf(w, "This is Correct")
	}
}

func div(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println("lets go ")
	http.HandleFunc("/", divide)
	http.ListenAndServe(":8085", nil)
}
