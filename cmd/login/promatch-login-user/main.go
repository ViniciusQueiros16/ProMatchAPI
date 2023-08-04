package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	authToken "github.com/promatch/cmd/auth/promatch-create-auth-token"
	"github.com/promatch/pkg/database"
	"github.com/promatch/pkg/utils/response"
	"github.com/promatch/structs"
	"golang.org/x/crypto/bcrypt"
)

func AuthenticateUser(db *sql.DB, usernameOrEmail, password string) (int64, error) {
	var user structs.Users
	query := "SELECT id, password FROM users WHERE username = ? OR email = ?"
	err := db.QueryRow(query, usernameOrEmail, usernameOrEmail).Scan(&user.ID, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, fmt.Errorf("User not found")
		}
		return 0, fmt.Errorf("AuthenticateUser: %w", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return 0, fmt.Errorf("Incorrect password")
	}

	return user.ID, nil
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	db, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer database.CloseDB()

	var authRequest structs.AuthRequest
	err = json.Unmarshal([]byte(request.Body), &authRequest)
	if err != nil {
		return response.ApiResponse(http.StatusBadRequest, structs.ErrorBody{
			ErrorMsg: aws.String(err.Error()),
		})
	}

	userID, err := AuthenticateUser(db, authRequest.UsernameOrEmail, authRequest.Password)
	if err != nil {
		return response.ApiResponse(http.StatusUnauthorized, structs.ErrorBody{
			ErrorMsg: aws.String(err.Error()),
		})
	}

	authToken, err := authToken.CreateAuthToken(db, userID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Token created: %v\n", authToken.Token)

	return response.ApiResponse(http.StatusCreated, authToken)
}

func main() {
	lambda.Start(handler)
}
