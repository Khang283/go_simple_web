package util

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var DATABASE_URL string

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
	}
	DATABASE_URL = os.Getenv("DATABASE_URL")
}
