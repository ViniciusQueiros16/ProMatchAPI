package main

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
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

