package users

import (
	"database/sql"
	"fmt"

	"structs"
)

func AddUser(db *sql.DB, user structs.Users) (int64, error) {
	stmt, err := db.Prepare("INSERT INTO users(name, email, password, created_at) VALUES (?, ?, ?, ?)")
	if err != nil {
		return 0, fmt.Errorf("AddUser: %v", err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(user.Name, user.Email, user.Password, user.CreatedAt)
	if err != nil {
		return 0, fmt.Errorf("AddUser: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("AddUser: %v", err)
	}

	return id, nil
}
