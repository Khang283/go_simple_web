package student

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Student struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Age    int    `json:"age"`
}

type StudentDeleteRequest struct {
	Id string `json:"id"`
}

var DATABASE_URL string = "postgresql://postgres:admin@localhost:5432/testdb"

func StudentRoute(r *mux.Router) {
	r.HandleFunc("/student", handleGetStudent).Methods("GET")
	r.HandleFunc("/student/{id}", handleGetStudentById).Methods("GET")
	r.HandleFunc("/student", handleCreateStudent).Methods("POST")
	r.HandleFunc("/student", handleDeleteStudent).Methods("DELETE")
	r.HandleFunc("/student", handleUpdateStudent).Methods("PUT")
}

func GetStudentList() []Student {
	dbpool, err := pgxpool.New(context.Background(), DATABASE_URL)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer dbpool.Close()
	students := make([]Student, 0)
	query := "select id, name, age, gender from students"
	rows, err := dbpool.Query(context.Background(), query)
	if err != nil {
		log.Fatal(err.Error())
	}
	for rows.Next() {
		var student Student
		if err = rows.Scan(&student.Id, &student.Name, &student.Age, &student.Gender); err != nil {
			log.Fatal(err.Error())
		}
		students = append(students, student)
	}
	if err != nil {
		log.Fatal(err.Error())
	}
	return students
}

func GetStudentById(id string) Student {
	dbpool, err := pgxpool.New(context.Background(), DATABASE_URL)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer dbpool.Close()
	var student Student
	query := `
  select id, name, age ,gender
  from students
  where id = $1
  `
	err = dbpool.QueryRow(context.Background(), query, id).Scan(&student.Id, &student.Name, &student.Age, &student.Gender)
	if err != nil {
		log.Fatal(err.Error())
	}
	return student
}

func CreateStudent(student *Student) bool {
	dbpool, err := pgxpool.New(context.Background(), DATABASE_URL)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer dbpool.Close()
	query := `
    insert into students(id, name, age, gender) values ($1, $2, $3, $4)
    `
	tag, err := dbpool.Exec(context.Background(), query, student.Id, student.Name, student.Age, student.Gender)
	if err != nil {
		log.Fatal(err.Error())
	}
	if tag.RowsAffected() != 0 {
		return true
	}
	return false
}

func DeleteStudent(id string) bool {
	query := `
    delete from students where id = $1
  `
	ctx := context.Background()
	dbpool, err := pgxpool.New(context.Background(), DATABASE_URL)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer dbpool.Close()
	tag, err := dbpool.Exec(ctx, query, id)
	if err != nil {
		log.Fatal(err.Error())
	}
	if tag.RowsAffected() != 0 {
		return true
	}
	return false
}

func UpdateStudent(student Student) bool {
	ctx := context.Background()
	query := `
    update students set name = $1, age = $2, gender = $3 where id = $4
  `
	dbpool, err := pgxpool.New(ctx, DATABASE_URL)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer dbpool.Close()
	tag, err := dbpool.Exec(ctx, query, student.Name, student.Age, student.Gender, student.Id)
	if err != nil {
		log.Fatal(err.Error())
	}
	if tag.RowsAffected() != 0 {
		return true
	}
	return false
}

func handleGetStudent(w http.ResponseWriter, r *http.Request) {
	student := GetStudentList()
	json.NewEncoder(w).Encode(student)
}

func handleGetStudentById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	student := GetStudentById(id)
	json.NewEncoder(w).Encode(student)
}

func handleCreateStudent(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	var s Student
	err := dec.Decode(&s)
	if err != nil {
		log.Fatal(err.Error())
	}
	s.Id = uuid.New().String()
	isInsert := CreateStudent(&s)
	if isInsert {
		json.NewEncoder(w).Encode(s)
	}
}

func handleDeleteStudent(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	var request StudentDeleteRequest
	err := dec.Decode(&request)
	if err != nil {
		log.Fatal(err.Error())
	}
	isDeleted := DeleteStudent(request.Id)
	if isDeleted {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func handleUpdateStudent(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	var student Student
	err := dec.Decode(&student)
	if err != nil {
		log.Fatal(err.Error())
	}
	isUpdated := UpdateStudent(student)
	if isUpdated {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}
