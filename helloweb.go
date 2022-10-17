package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "First Web Program")
		// n, err := fmt.Fprintf(w, "First Web Program")
		// if err != nil {
		// 	fmt.Println(nil, n)
		// }
	})

	_ = http.ListenAndServe(":8081", nil)
}
