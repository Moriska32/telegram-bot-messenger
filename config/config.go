package config

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

//Connect Connecting to db
func Connect() *sql.DB {

	db, err := sql.Open("postgres", "Your BD")

	if err != nil {
		log.Printf("Reason: %v\n", err)
		os.Exit(100)
	}
	log.Printf("Connected to db")
	return db
}
