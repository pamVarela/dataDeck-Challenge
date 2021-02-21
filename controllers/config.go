package controllers

import (
	"log"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func GetConnectionDB() *sql.DB {

	log.Println("Connecting DB...")
	if db != nil {
		return db
	}

	db, err := sql.Open("sqlite3", "./jrdd.db")
	
	if err != nil {
		log.Println("Failed to connect to DB...")
	}

	return db
}

