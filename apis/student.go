package apis

import (
	"crud/daos"
	"crud/types"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strconv"

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
	student := types.GetStudentResponse{
		Data:   daos.GetStudentList(),
		Status: "200",
		Error:  "",
	}
	dir, _ := filepath.Abs("web/")
	tmpl := template.Must(template.ParseFiles(dir+"/template/layout/index.html", dir+"/template/student_page.html"))
	tmpl.Execute(w, student)
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
	var requestBody types.CreateStudentRequest
	err := dec.Decode(&requestBody)
	if err != nil {
		log.Fatal(err.Error())
	}
	age, err := strconv.Atoi(requestBody.Age)
	if err != nil {
		log.Print(err.Error())
	}
	s = types.Student{
		Id:     uuid.New().String(),
		Name:   requestBody.Name,
		Age:    age,
		Gender: requestBody.Gender,
	}
	isInsert := daos.CreateStudent(&s)
	if isInsert {
		w.WriteHeader(http.StatusOK)
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
