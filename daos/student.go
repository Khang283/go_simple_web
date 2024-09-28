package daos

import (
	"context"
	"crud/types"
	"crud/util"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DATABASE_URL *string = &util.DATABASE_URL

func GetStudentList() []types.Student {
	dbpool, err := pgxpool.New(context.Background(), *DATABASE_URL)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer dbpool.Close()
	students := make([]types.Student, 0)
	query := "select id, name, age, gender from students"
	rows, err := dbpool.Query(context.Background(), query)
	if err != nil {
		log.Fatal(err.Error())
	}
	for rows.Next() {
		var student types.Student
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

func GetStudentById(id string) types.Student {
	dbpool, err := pgxpool.New(context.Background(), *DATABASE_URL)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer dbpool.Close()
	var student types.Student
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

func CreateStudent(student *types.Student) bool {
	dbpool, err := pgxpool.New(context.Background(), *DATABASE_URL)
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
	dbpool, err := pgxpool.New(context.Background(), *DATABASE_URL)
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

func UpdateStudent(student types.Student) bool {
	ctx := context.Background()
	query := `
    update students set name = $1, age = $2, gender = $3 where id = $4
  `
	dbpool, err := pgxpool.New(ctx, *DATABASE_URL)
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
