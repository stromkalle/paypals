package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func New() *sql.DB {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	fmt.Println("Connected to PostgreSQL successfully")
	return db
}
