package users

import (
	"database/sql"
	"fmt"

	"github.com/promatch/structs"
)

func UpdateUser(db *sql.DB, user structs.Users) error {
	stmt, err := db.Prepare("UPDATE users SET name = ?, email = ?, password = ?, created_at = ? WHERE id = ?")
	if err != nil {
		return fmt.Errorf("UpdateUser: %v", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Name, user.Email, user.Password, user.CreatedAt, user.ID)
	if err != nil {
		return fmt.Errorf("UpdateUser: %v", err)
	}

	return nil
}
