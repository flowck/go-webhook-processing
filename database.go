package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var DbConn *sql.DB

func SetupDabatase() {
	var err error
	DbConn, err = sql.Open("sqlite3", "database.db")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to the database")
}
