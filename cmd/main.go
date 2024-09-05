package main

import (
	"crud/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	PORT := ":8080"
	router := router.CreateNewRouter()
	fmt.Println("Server start at port 8080")
	log.Fatal(http.ListenAndServe(PORT, router))
}
