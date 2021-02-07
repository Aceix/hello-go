package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := 3000

	// serve static files
	startFileServer()
	
	// /hello
	http.HandleFunc(
		"/hello",
		helloHandler,
	)

	// /form
	http.HandleFunc("/form", formHandler)

	fmt.Println("Starting server at port", port)
	if err := http.ListenAndServe(
		fmt.Sprintf(":%d", port), nil); err != nil {
		log.Fatalln(err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/hello", "hello":
		switch r.Method {
		case "":
			fallthrough
		case http.MethodGet:
			fmt.Fprintln(w, "Hello world!")
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	default:
		http.Error(w, "Not found", http.StatusNotFound)
	}
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/form", "form":
		switch r.Method {
		case "":
			fallthrough
		case http.MethodPost:
			if err := r.ParseForm(); err != nil {
				fmt.Fprintf(w, "Error: %v\n", err)
				http.Error(w, "Bad request", http.StatusBadRequest)
				return
			}

			name := r.FormValue("name")
			age := r.FormValue("age")
			
			fmt.Fprintf(w, "Name: %v\n", name)
			fmt.Fprintf(w, "Age: %v\n", age)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	default:
		http.Error(w, "Not found", http.StatusNotFound)
	}
}

func startFileServer() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)	
}
