package main

import (
	"database/sql"
	"fmt"
	"log"
	"login"
	"os"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {

	// Capture connection properties.
	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "localhost:3306",
		DBName: "proposta",
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

	//----------------------------------

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := login.Hello("Vinny")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
