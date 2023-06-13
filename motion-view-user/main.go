package main

import (
	"database/sql"
	database "db"
	"fmt"
	"log"
	"structs"

	"github.com/aws/aws-lambda-go/lambda"
)

func UserByUsers(db *sql.DB, name string) ([]structs.Users, error) {
	var findUsers []structs.Users

	rows, err := db.Query("SELECT * FROM users WHERE name = ?", name)
	if err != nil {
		return nil, fmt.Errorf("userByUsers %q: %v", name, err)
	}
	defer rows.Close()

	for rows.Next() {
		var user structs.Users
		var createdAt []uint8
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &createdAt); err != nil {
			return nil, fmt.Errorf("userByUsers %q: %v", name, err)
		}

		if err != nil {
			return nil, fmt.Errorf("userByUsers %q: %v", name, err)
		}
		findUsers = append(findUsers, user)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("userByUsers %q: %v", name, err)
	}
	return findUsers, nil
}

func handler() {
	db, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer database.CloseDB()

	userByName, err := UserByUsers(db, "Tony")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Users found: %v\n", userByName)
}

func main() {
	lambda.Start(handler)
}
