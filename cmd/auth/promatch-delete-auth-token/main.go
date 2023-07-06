package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/promatch/pkg/database"
	"github.com/promatch/pkg/utils/response"
	"github.com/promatch/structs"
)

func DeleteAuthToken(db *sql.DB, token string) error {
	stmt, err := db.Prepare("DELETE FROM auth_tokens WHERE token = ?")
	if err != nil {
		return fmt.Errorf("DeleteAuthToken: %w", err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(token)
	if err != nil {
		return fmt.Errorf("DeleteAuthToken: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("DeleteAuthToken: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("DeleteAuthToken: Token not found")
	}

	return nil
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	db, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer database.CloseDB()

	token := request.QueryStringParameters["token"]

	err = DeleteAuthToken(db, token)
	if err != nil {
		return response.ApiResponse(http.StatusBadRequest, structs.ErrorBody{
			ErrorMsg: aws.String(err.Error()),
		})
	}
	return response.ApiResponse(http.StatusOK, token)
}

func main() {
	lambda.Start(handler)
}
