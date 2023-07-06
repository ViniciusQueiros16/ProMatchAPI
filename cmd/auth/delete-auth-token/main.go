package main

import (
	"database/sql"
	"fmt"
)

const (
	tokenLength      = 64
	tokenExpiryHours = 24
)


func DeleteAuthToken(db *sql.DB, token string) error {
	stmt, err := db.Prepare("DELETE FROM auth_tokens WHERE token = ?")
	if err != nil {
		return fmt.Errorf("DeleteAuthToken: %v", err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(token)
	if err != nil {
		return fmt.Errorf("DeleteAuthToken: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("DeleteAuthToken: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("DeleteAuthToken: Token not found")
	}

	return nil
}
