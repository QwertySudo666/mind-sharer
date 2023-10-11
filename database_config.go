package main

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB

func configureMySQLDatabase() {

	// Capture connection properties.
	cfg := mysql.Config{
		User:   "root",      //os.Getenv("root"),
		Passwd: "localhost", //os.Getenv("localhost"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "mind_sharer",
	}
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
}
