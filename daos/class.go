package daos

import (
	"context"
	"crud/types"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func GetClassList() []types.Class {
	dbpool, err := pgxpool.New(context.Background(), *DATABASE_URL)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer dbpool.Close()
	classses := make([]types.Class, 0)
	query := "select id, name, description, type, code from class"
	rows, err := dbpool.Query(context.Background(), query)
	if err != nil {
		log.Fatal(err.Error())
	}
	for rows.Next() {
		var class types.Class
		if err = rows.Scan(&class.Id, &class.Name, &class.Description, &class.Type, &class.Code); err != nil {
			log.Fatal(err.Error())
		}
		classses = append(classses, class)
	}
	if err != nil {
		log.Fatal(err.Error())
	}
	return classses
}
