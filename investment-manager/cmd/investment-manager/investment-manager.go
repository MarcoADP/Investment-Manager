package main

import (
	"context"
	"log"

	"os"

	"github.com/MarcoADP/Investment-Manager/internal/db"

	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()

	dbHost := os.Getenv("DB_HOST")
	log.Println(dbHost)

	database, err := db.Connect(context.Background())
	if err != nil {
		log.Fatalf(
			"Error trying to connect to database, error=%s \n",
			err.Error())
		return
	}
	log.Println(database.Stats())

}
