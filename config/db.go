package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func DatabaseConnection() (*sql.DB, error) {
	host := "localhost"
	port := 5432
	user := "postgres"
	password := "Yuvi1808"
	dbname := "first"
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Println("Failed to connnect database")
		return nil, err
	}
	return db, nil
}

func GetDB() (*sql.DB, error) {
	db, err := DatabaseConnection()
	if err != nil {
		return nil, err
	}
	return db, nil
}
