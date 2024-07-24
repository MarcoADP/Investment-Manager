package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func Connect(ctx context.Context) (*sql.DB, error) {

	dbHost := os.Getenv("DB_HOST")
	log.Println(dbHost)
	dbPort := os.Getenv("DB_PORT")
	log.Println(dbPort)
	dbUser := os.Getenv("DB_USER")
	log.Println(dbUser)
	dbPassword := os.Getenv("DB_PASSWORD")
	log.Println(dbPassword)
	dbName := os.Getenv("DB_NAME")
	log.Println(dbName)

	// Construir string de conexão
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	// Conectar ao banco de dados
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Erro ao abrir a conexão com o banco de dados: %v", err)
		return nil, err
	}
	defer db.Close()

	// Verificar se a conexão foi bem-sucedida
	err = db.Ping()
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
		return nil, err
	}

	fmt.Println("Conexão com o banco de dados bem-sucedida!")

	return db, nil

}
