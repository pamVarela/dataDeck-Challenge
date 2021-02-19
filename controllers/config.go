package controllers

import (
	"database/sql"
	"log"
	
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func GetConnectionDB() *sql.DB {

	log.Println("Connecting DB...")
	if db != nil {
		return db
	}

	var err error

	db, err := sql.Open("sqlite3", "./jrdd.db")
	if err != nil {
		log.Println(err)
	}
	return db
}

