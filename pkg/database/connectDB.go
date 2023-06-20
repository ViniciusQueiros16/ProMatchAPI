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
		User:                 os.Getenv("DBUSER"),
		Passwd:               os.Getenv("DBPASS"),
		Net:                  "tcp",
		Addr:                 "proposta.cynfirggau4l.us-east-2.rds.amazonaws.com:3306",
		DBName:               "proposta",
		ParseTime:            true,
		AllowNativePasswords: true,
	}

	// Get a database handle.
	var openErr error
	db, openErr = sql.Open("mysql", cfg.FormatDSN())
	if openErr != nil {
		return nil, openErr
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
