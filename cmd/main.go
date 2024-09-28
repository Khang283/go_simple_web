package main

import (
	"crud/router"
	"crud/util"
	"fmt"
	"log"
	"net/http"
)

func Init() {
	util.LoadEnv()
}

func main() {
	PORT := ":8080"
	Init()
	router := router.CreateNewRouter()
	// http.Handle("/css/", http.StripPrefix("/css", http.FileServer(http.Dir("./css/"))))
	// http.Handle("/css/", http.FileServer(http.Dir("./css")))
	fmt.Println("Server start at port 8080")
	log.Fatal(http.ListenAndServe(PORT, router))
}
