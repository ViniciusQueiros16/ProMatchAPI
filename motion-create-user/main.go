package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"response"
	"structs"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	database "db"
)

type ErrorBody struct {
	ErrorMsg *string `json:"error,omitempty"`
}

func CreateUser(db *sql.DB, user structs.Users) (int64, error) {
	stmt, err := db.Prepare("INSERT INTO users(name, email, password, created_at) VALUES (?, ?, ?, ?)")
	if err != nil {
		return 0, fmt.Errorf("CreateUser: %w", err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(user.Name, user.Email, user.Password, user.CreatedAt)
	if err != nil {
		return 0, fmt.Errorf("CreateUser: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("CreateUser: %w", err)
	}

	return id, nil
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	db, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer database.CloseDB()

	user := structs.Users{
		Name:      "Sergio",
		Email:     "sergio@example.com",
		Password:  "senha6",
		CreatedAt: time.Now(),
	}

	userID, err := CreateUser(db, user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID of added user: %v\n", userID)

	return response.ApiResponse(http.StatusCreated, userID)
}

func main() {
	lambda.Start(handler)
}
