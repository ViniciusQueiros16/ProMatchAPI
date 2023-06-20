package main

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"errors"
	"fmt"
	"time"

	"github.com/promatch/structs"
)

const (
	tokenLength      = 64
	tokenExpiryHours = 24
)

func GenerateAuthToken() (string, error) {
	tokenBytes := make([]byte, tokenLength)
	_, err := rand.Read(tokenBytes)
	if err != nil {
		return "", fmt.Errorf("GenerateAuthToken: %v", err)
	}
	token := base64.URLEncoding.EncodeToString(tokenBytes)
	return token, nil
}

func CreateAuthToken(db *sql.DB, userID int64) (structs.AuthToken, error) {
	token, err := GenerateAuthToken()
	if err != nil {
		return structs.AuthToken{}, fmt.Errorf("CreateAuthToken: %v", err)
	}

	expiresAt := time.Now().Add(time.Hour * tokenExpiryHours)

	stmt, err := db.Prepare("INSERT INTO auth_tokens(user_id, token, expires_at) VALUES (?, ?, ?)")
	if err != nil {
		return structs.AuthToken{}, fmt.Errorf("CreateAuthToken: %v", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(userID, token, expiresAt)
	if err != nil {
		return structs.AuthToken{}, fmt.Errorf("CreateAuthToken: %v", err)
	}

	authToken := structs.AuthToken{
		Token:     token,
		ExpiresAt: expiresAt,
	}

	return authToken, nil
}

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
