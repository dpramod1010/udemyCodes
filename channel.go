package main

import "fmt"

func main() {
	message := make(chan string, 2)
	go func() {
		message <- "One"
		message <- "two"

	}()
	msg := <-message
	fmt.Println(msg)

}
