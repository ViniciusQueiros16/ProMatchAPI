package authToken

import (
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/promatch/pkg/database"
	"github.com/promatch/pkg/utils/response"
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
		return "", fmt.Errorf("GenerateAuthToken: %w", err)
	}
	token := base64.URLEncoding.EncodeToString(tokenBytes)
	return token, nil
}

func CreateAuthToken(db *sql.DB, userID int64) (structs.AuthToken, error) {
	token, err := GenerateAuthToken()
	if err != nil {
		return structs.AuthToken{}, fmt.Errorf("CreateAuthToken: %w", err)
	}

	expiresAt := time.Now().Add(time.Hour * tokenExpiryHours)

	stmt, err := db.Prepare("INSERT INTO auth_tokens(user_id, token, expires_at) VALUES (?, ?, ?)")
	if err != nil {
		return structs.AuthToken{}, fmt.Errorf("CreateAuthToken: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(userID, token, expiresAt)
	if err != nil {
		return structs.AuthToken{}, fmt.Errorf("CreateAuthToken: %w", err)
	}

	authToken := structs.AuthToken{
		Token:     token,
		ExpiresAt: expiresAt,
	}

	return authToken, nil
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	db, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer database.CloseDB()

	var user structs.Users
	err = json.Unmarshal([]byte(request.Body), &user)
	if err != nil {
		return response.ApiResponse(http.StatusBadRequest, structs.ErrorBody{
			ErrorMsg: aws.String(err.Error()),
		})
	}

	authToken, err := CreateAuthToken(db, user.ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Token created: %v\n", authToken.Token)

	return response.ApiResponse(http.StatusCreated, authToken)
}

func main() {
	lambda.Start(handler)
}
