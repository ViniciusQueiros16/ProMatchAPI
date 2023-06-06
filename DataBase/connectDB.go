package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitDB() (*sql.DB, error) {
	// Capture connection properties.
	cfg := mysql.Config{
		User:      os.Getenv("DBUSER"),
		Passwd:    os.Getenv("DBPASS"),
		Net:       "tcp",
		Addr:      "localhost:3306",
		DBName:    "proposta",
		ParseTime: true,
	}
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, err
	}

	pingErr := db.Ping()
	if pingErr != nil {
		return nil, pingErr
	}

	fmt.Println("Connected to database!")
	return db, nil
}

func CloseDB() {
	if db != nil {
		db.Close()
	}
}
