package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/promatch/pkg/database"
	"github.com/promatch/pkg/utils/response"
	"github.com/promatch/structs"
)

func UpdateUser(db *sql.DB, user structs.Users) error {
	stmt, err := db.Prepare("UPDATE users SET name = ?, email = ?, password = ?, created_at = ? WHERE id = ?")
	if err != nil {
		return fmt.Errorf("UpdateUser: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Email, user.Password, user.CreatedAt, user.ID)
	if err != nil {
		return fmt.Errorf("UpdateUser: %w", err)
	}

	return nil
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	db, err := database.InitDB()
	if err != nil {
		return response.ApiResponse(http.StatusInternalServerError, structs.ErrorBody{
			ErrorMsg: aws.String(err.Error()),
		})
	}
	defer database.CloseDB()

	var user structs.Users
	err = json.Unmarshal([]byte(request.Body), &user)
	if err != nil {
		return response.ApiResponse(http.StatusBadRequest, structs.ErrorBody{
			ErrorMsg: aws.String(err.Error()),
		})
	}

	err = UpdateUser(db, user)
	if err != nil {
		return response.ApiResponse(http.StatusBadRequest, structs.ErrorBody{
			ErrorMsg: aws.String(err.Error()),
		})
	}

	return response.ApiResponse(http.StatusOK, user)
}

func main() {
	lambda.Start(handler)
}
