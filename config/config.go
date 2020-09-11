package config

import (
	"database/sql"
	"log"
	"os"
	
	_ "github.com/lib/pq"
)

//Connect Connecting to db
func Connect() *sql.DB {
	//"postgres://kot_user:1qaz@WSX@172.20.0.78:5432/hospital_db"
	db, err := sql.Open("postgres", "postgres://kot_user:1qaz@WSX@172.20.0.78:5432/hospital_db")

	if err != nil {
		log.Printf("Reason: %v\n", err)
		os.Exit(100)
	}
	log.Printf("Connected to db")
	return db
}
