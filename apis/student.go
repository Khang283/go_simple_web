package apis

import (
	"crud/daos"
	"crud/types"
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func StudentRoute(r *mux.Router) {
	r.HandleFunc("/student", handleGetStudent).Methods("GET")
	r.HandleFunc("/student/{id}", handleGetStudentById).Methods("GET")
	r.HandleFunc("/student", handleCreateStudent).Methods("POST")
	r.HandleFunc("/student", handleDeleteStudent).Methods("DELETE")
	r.HandleFunc("/student", handleUpdateStudent).Methods("PUT")
}

func handleGetStudent(w http.ResponseWriter, r *http.Request) {
	student := daos.GetStudentList()
	json.NewEncoder(w).Encode(student)
}

func handleGetStudentById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	student := daos.GetStudentById(id)
	json.NewEncoder(w).Encode(student)
}

func handleCreateStudent(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	var s types.Student
	err := dec.Decode(&s)
	if err != nil {
		log.Fatal(err.Error())
	}
	s.Id = uuid.New().String()
	isInsert := daos.CreateStudent(&s)
	if isInsert {
		json.NewEncoder(w).Encode(s)
	}
}

func handleDeleteStudent(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	var request types.StudentDeleteRequest
	err := dec.Decode(&request)
	if err != nil {
		log.Fatal(err.Error())
	}
	isDeleted := daos.DeleteStudent(request.Id)
	if isDeleted {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func handleUpdateStudent(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	var student types.Student
	err := dec.Decode(&student)
	if err != nil {
		log.Fatal(err.Error())
	}
	isUpdated := daos.UpdateStudent(student)
	if isUpdated {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}
