package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "Method  not supported", http.StatusNotFound)
		return
	}

	fmt.Println(r.Form)
	fmt.Println("Name :", r.Form["name"][0])
	fmt.Println("Email :", r.Form["email"][0])
	//fmt.Println("Email :", r.Form["email"][1])
	fmt.Fprintf(w, "Name: %s", r.Form["name"][0])
	fmt.Fprintf(w, "Email : %s", r.Form["email"][0])
	//fmt.Fprintf(w, "Email : %s", r.Form["email"][1])
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hello")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8010\n")
	if err := http.ListenAndServe(":8010", nil); err != nil {
		log.Fatal(err)
	}
}
