package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID       string
	Name     string
	Email    string
	Password string
}

func InitDB() *sql.DB {
	path := "tmp/database.db"
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		log.Println(err)
	}
	return db
}
