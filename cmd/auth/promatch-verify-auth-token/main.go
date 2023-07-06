package main

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
)

const (
	tokenLength      = 64
	tokenExpiryHours = 24
)



func VerifyAuthToken(db *sql.DB, token string) (int64, error) {
	var userID int64
	var expiresAt time.Time

	err := db.QueryRow("SELECT user_id, expires_at FROM auth_tokens WHERE token = ? LIMIT 1", token).Scan(&userID, &expiresAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, errors.New("VerifyAuthToken: Invalid token")
		}
		return 0, fmt.Errorf("VerifyAuthToken: %v", err)
	}

	if expiresAt.Before(time.Now()) {
		return 0, errors.New("VerifyAuthToken: Token has expired")
	}

	return userID, nil
}

