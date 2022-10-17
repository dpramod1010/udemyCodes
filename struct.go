package main

import (
	"log"
	"time"
)

type User struct {
	firstName   string
	lastName    string
	phoneNumber string
	birthDate   time.Time
}

func main() {
	user := User{firstName: "Ptamod", lastName: "Dhayarkar"}
	log.Println("first ", user.firstName)
}
