package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	authToken "github.com/promatch/cmd/auth/promatch-create-auth-token"
	"github.com/promatch/pkg/database"
	"github.com/promatch/pkg/utils/response"
	"github.com/promatch/structs"
	"golang.org/x/crypto/bcrypt"
)

type UserResponse struct {
	Token    string `json:"token"`
	Username string `json:"username"`
}

func CreateUser(db *sql.DB, user structs.Users) (int64, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, fmt.Errorf("CreateUser: %w", err)
	}

	stmt, err := db.Prepare("INSERT INTO users(username, name, email, password, created_at) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return 0, fmt.Errorf("CreateUser: %w", err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(user.Username, user.Name, user.Email, string(hashedPassword), user.CreatedAt)
	if err != nil {
		return 0, fmt.Errorf("CreateUser: %w", err)
	}

	userID, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("CreateUser: %w", err)
	}

	// Create a profile entry for the user
	profileStmt, err := db.Prepare("INSERT INTO profile(user_id) VALUES (?)")
	if err != nil {
		return 0, fmt.Errorf("CreateUser: %w", err)
	}
	defer profileStmt.Close()

	_, err = profileStmt.Exec(userID)
	if err != nil {
		return 0, fmt.Errorf("CreateUser: %w", err)
	}

	return userID, nil
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

	user.CreatedAt = time.Now()

	userID, err := CreateUser(db, user)
	authToken, err := authToken.CreateAuthToken(db, userID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID of added user: %v\n", userID)

	userResponse := UserResponse{
		Token:    authToken.Token,
		Username: user.Username,
	}

	return response.ApiResponse(http.StatusCreated, userResponse)
}

func main() {
	lambda.Start(handler)
}
