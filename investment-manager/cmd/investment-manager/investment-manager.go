package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	_ "github.com/MarcoADP/Investment-Manager/docs"
	"github.com/MarcoADP/Investment-Manager/internal/db"
	v1 "github.com/MarcoADP/Investment-Manager/pkg/api/v1"
)

func main() {

	godotenv.Load()
	log.Println(os.Getenv("DB_URL"))
	/*


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
	*/

	cfg := db.LoadConfig(context.Background())

	database, err := db.ConnectDatabase(cfg)
	if err != nil {
		log.Fatal("cannot connect to database:", err)
	}
	log.Println("Connected to", database.Config.Name())

	router := v1.CreateRoutes(database)

	log.Fatal(http.ListenAndServe(":8080", router))

}
