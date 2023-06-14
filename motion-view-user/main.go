package main

import (
	"context"
	"database/sql"
	database "db"
	"fmt"
	"log"
	"net/http"
	"response"
	"structs"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
)

type ErrorBody struct {
	ErrorMsg *string `json:"error,omitempty"`
}

func UserByUsers(db *sql.DB, name string) ([]structs.Users, error) {
	var findUsers []structs.Users

	rows, err := db.Query("SELECT * FROM users WHERE name = ?", name)
	if err != nil {
		return nil, fmt.Errorf("userByUsers %q: %v", name, err)
	}
	defer rows.Close()

	for rows.Next() {
		var user structs.Users
		var createdAt []uint8
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &createdAt); err != nil {
			return nil, fmt.Errorf("userByUsers %q: %v", name, err)
		}

		if err != nil {
			return nil, fmt.Errorf("userByUsers %q: %v", name, err)
		}
		findUsers = append(findUsers, user)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("userByUsers %q: %v", name, err)
	}
	return findUsers, nil
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	db, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer database.CloseDB()

	name := request.QueryStringParameters["name"]

	result, err := UserByUsers(db, name)
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
