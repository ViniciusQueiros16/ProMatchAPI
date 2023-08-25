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
	"github.com/google/uuid"
	"github.com/promatch/pkg/database"
	"github.com/promatch/pkg/utils/response"
	"github.com/promatch/pkg/utils/token"
	"github.com/promatch/pkg/utils/uploadS3"
	"github.com/promatch/structs"
)

type UpdateProfileRequest struct {
	Avatar    string `json:"avatar"`
	Birthdate string `json:"birthdate"`
	Company   string `json:"company"`
	Gender    string `json:"gender"`
}

func UpdateUserProfile(db *sql.DB, userID int64, request UpdateProfileRequest) error {
	stmt, err := db.Prepare(`
		UPDATE profile
		SET avatar = ?, birthdate = ?, company = ?, gender = ?, updated_at = NOW()
		WHERE user_id = ?;
	`)
	if err != nil {
		return fmt.Errorf("UpdateUserProfile: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(request.Avatar, request.Birthdate, request.Company, request.Gender, userID)
	if err != nil {
		return fmt.Errorf("UpdateUserProfile: %w", err)
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

	tokenString := request.QueryStringParameters["token"]

	// Decode the token to get the user ID
	userID, err := token.DecodeAuthToken(tokenString)
	if err != nil {
		return response.ApiResponse(http.StatusBadRequest, structs.ErrorBody{
			ErrorMsg: aws.String(err.Error()),
		})
	}

	var updateRequest UpdateProfileRequest
	err = json.Unmarshal([]byte(request.Body), &updateRequest)
	if err != nil {
		return response.ApiResponse(http.StatusBadRequest, structs.ErrorBody{
			ErrorMsg: aws.String(err.Error()),
		})
	}

	if updateRequest.Avatar != "" {
		imageUUID := uuid.New().String()
		resp := uploadS3.ImageUpload(imageUUID, updateRequest.Avatar)
		updateRequest.Avatar = resp.Location
	}

	err = UpdateUserProfile(db, userID, updateRequest)
	if err != nil {
		return response.ApiResponse(http.StatusBadRequest, structs.ErrorBody{
			ErrorMsg: aws.String(err.Error()),
		})
	}

	return response.ApiResponse(http.StatusOK, userID)
}

func main() {
	lambda.Start(handler)
}
