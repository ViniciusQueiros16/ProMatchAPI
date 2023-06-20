package users

import (
	"database/sql"
	"fmt"
)

func DeleteUser(db *sql.DB, userID int64) error {
	stmt, err := db.Prepare("DELETE FROM users WHERE id = ?")
	if err != nil {
		return fmt.Errorf("DeleteUser: %v", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(userID)
	if err != nil {
		return fmt.Errorf("DeleteUser: %v", err)
	}

	return nil
}
