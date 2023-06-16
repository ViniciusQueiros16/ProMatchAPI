package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"response"
	"structs"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"

	database "db"
)

type ErrorBody struct {
	ErrorMsg *string `json:"error,omitempty"`
}

func FetchUser(db *sql.DB, name string) ([]structs.Users, error) {
	var findUser []structs.Users

	rows, err := db.Query("SELECT * FROM users WHERE name = ?", name)
	if err != nil {
		return nil, fmt.Errorf("FetchUser: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var user structs.Users
		var createdAt []uint8
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &createdAt); err != nil {
			return nil, fmt.Errorf("FetchUser: %w", err)
		}

		findUser = append(findUser, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("FetchUser: %w", err)
	}

	return findUser, nil
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	db, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer database.CloseDB()

	name := request.QueryStringParameters["name"]

	result, err := FetchUser(db, name)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Users found: %v\n", result)

	if err != nil {
		return response.ApiResponse(http.StatusBadRequest, ErrorBody{
			aws.String(err.Error()),
		})
	}

	return response.ApiResponse(http.StatusOK, result)
}

func main() {
	lambda.Start(handler)
}
