package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	PORT := ":8080"
	mux := http.NewServeMux()
	mux.Handle("/", helloWorld())
	mux.HandleFunc("/name/{name}", helloName)
	fmt.Println("Server start at port 8080")
	log.Fatal(http.ListenAndServe(PORT, mux))
}

func helloWorld() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World")
	}
}

func helloName(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")
	result := fmt.Sprintf("Hello %s", name)
	fmt.Fprintln(w, result)
}
