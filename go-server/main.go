package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileserver := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileserver)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Starting server at port 8080.")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func helloHandler(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/hello" {
		http.Error(writer, "404 Not Found", http.StatusNotFound)
		return
	}

	if request.Method != "GET" {
		http.Error(writer, "Method is not supported.", http.StatusNotFound)
		return
	}

	fmt.Fprint(writer, "Hello!")
}

func formHandler(writer http.ResponseWriter, request *http.Request) {
	if err := request.ParseForm(); err != nil {
		fmt.Fprintf(writer, "ParseForm err(): %v", err)
		return
	}

	fmt.Fprintf(writer, "POST request successful!\n")
	name := request.FormValue("name")
	address := request.FormValue("address")

	fmt.Fprintf(writer, "Name = %s\n", name)
	fmt.Fprintf(writer, "Address = %s\n", address)
}
