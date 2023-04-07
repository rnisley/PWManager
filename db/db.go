package db

import (
	"database/sql"
	"log"
	"os"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	Db *sql.DB
}

var instance *Database
var once sync.Once

func Connect() *Database {
	once.Do(func() {
		mustCreateDb := false
		if _, err := os.Stat("dd.db"); err != nil {
			mustCreateDb = true
		}

		database, err := sql.Open("sqlite3", "dd.db")
		if err != nil {
			log.Fatalf("Error connecting to database: %v", err)
		}

		if mustCreateDb {
			if command, err := os.ReadFile("init.sql"); err == nil {
				_, err := database.Exec(string(command))
				if err != nil {
					log.Fatalf("Error initializing database: %v", err)
				}
			} else {
				log.Fatalf("Error loading sqlite initial schema")
			}
		}

		instance = &Database{Db: database}
	})

	return instance
}
