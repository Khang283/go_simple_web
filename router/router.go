package router

import (
	"crud/student"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateNewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", handleIndex)
	student.StudentRoute(r)
	// r.HandleFunc("/student", handleGetStudent).Methods("GET")
	// r.HandleFunc("/student/{id}", handleGetStudentById).Methods("GET")
	// r.HandleFunc("/student", handleCreateStudent).Methods("POST")
	return r
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}
