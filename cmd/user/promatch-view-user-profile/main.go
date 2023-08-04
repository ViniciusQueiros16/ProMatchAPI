package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/promatch/pkg/database"
	"github.com/promatch/pkg/utils/response"
	"github.com/promatch/structs"
)

func FetchUserProfile(db *sql.DB, id int64) (structs.UserProfile, error) {
	var userProfile structs.UserProfile

	query := `
		SELECT u.id, u.name, u.username, u.email, p.avatar, p.birthdate, p.company, p.gender
		FROM users u
		LEFT JOIN profile p ON u.id = p.id_user
		WHERE u.id = ?  LIMIT 1;
	`

	row := db.QueryRow(query, id)

	var birthdate sql.NullTime

	if err := row.Scan(
		&userProfile.UserID, &userProfile.Name, &userProfile.Username, &userProfile.Email,
		&userProfile.Avatar, &birthdate, &userProfile.Company, &userProfile.Gender,
	); err != nil {
		if err == sql.ErrNoRows {
			return structs.UserProfile{}, nil
		}
		return structs.UserProfile{}, fmt.Errorf("FetchUserProfile: %w", err)
	}

	if birthdate.Valid {
		userProfile.Birthdate = birthdate.Time
	}

	return userProfile, nil
}

func handler(ctc context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	db, err := database.InitDB()
	if err != nil {
		return response.ApiResponse(http.StatusInternalServerError, structs.ErrorBody{
			ErrorMsg: aws.String(err.Error()),
		})
	}
	defer database.CloseDB()

	userIDStr := request.QueryStringParameters["userID"]
	userID, err := strconv.ParseInt(userIDStr, 10, 64)

	if err != nil {
		return response.ApiResponse(http.StatusBadRequest, structs.ErrorBody{
			ErrorMsg: aws.String("Invalid userID format: must be an integer"),
		})
	}

	result, err := FetchUserProfile(db, userID)
	if err != nil {
		return response.ApiResponse(http.StatusBadRequest, structs.ErrorBody{
			ErrorMsg: aws.String(err.Error()),
		})
	}

	fmt.Printf("User Profile found: %+v\n", result)

	return response.ApiResponse(http.StatusOK, result)
}

func main() {
	lambda.Start(handler)
}
