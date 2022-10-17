package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const portNumber = ":8081"

// Home is the handler for the home page
func home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.tmpl")
	//fmt.Fprintf(w, "Home Page")
}

// About is the handler for the about page
func about(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.tmpl")
	//fmt.Fprintf(w, "About Page...")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./template/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template:", err)
	}
}

// main is the main function
func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)

	//fmt.Fprintf(fmt.Sprintf("Staring application on port %s", portNumber))
	fmt.Printf("Starting on Server %s", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}

//go get -u github.com/gorilla/mux
//go mod vendor
//go mod tidy
