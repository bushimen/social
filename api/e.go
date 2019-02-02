package models

import (
	"os"

	"github.com/go-pg/pg"
)

var db *pg.DB

// Connect initialize db connection
func Connect() {
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "localhost:5432"
	}

	db = pg.Connect(&pg.Options{
		Addr:     host,
		User:     "postgres",
		Password: "postgres",
		Database: "social",
	})

	if db == nil {
		panic("DB connection cannot be established")
	}
}

// Disconnect closes the db connection
func Disconnect() {
	db.Close()
}
