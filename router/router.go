package router

import (
	"crud/apis"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateNewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", handleIndex)
	r.HandleFunc("/test", HandlerHello)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static", http.FileServer(http.Dir("./web/static/"))))
	apis.StudentRoute(r)
	// r.HandleFunc("/student", handleGetStudent).Methods("GET")
	// r.HandleFunc("/student/{id}", handleGetStudentById).Methods("GET")
	// r.HandleFunc("/student", handleCreateStudent).Methods("POST")
	return r
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./web/static/index.html")
}

func HandlerHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
	w.WriteHeader(http.StatusOK)
}
