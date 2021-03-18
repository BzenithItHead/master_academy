package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)

	http.ListenAndServe(":8888", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is my webpage")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is about section")
}

func contact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is contact section")
}
