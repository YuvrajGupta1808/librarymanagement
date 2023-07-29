package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file. Make sure it exists or set the environment variables manually.")
	}
}

func DatabaseConnection() (*sql.DB, error) {
	loadEnv()
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbname := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Println("Failed to connect to the database")
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
