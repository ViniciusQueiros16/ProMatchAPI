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
	Avatar      string `json:"avatar"`
	Birthdate   string `json:"birthdate"`
	Gender      string `json:"gender"`
	About       string `json:"about"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	CoverPhoto  string `json:"cover_photo"`
	PhoneNumber string `json:"phone_number"`
}

type UpdateProfileResponse struct {
	UserID int64 `json:"userID"`
}

func UpdateUserProfile(db *sql.DB, userID int64, request UpdateProfileRequest) error {
	stmt, err := db.Prepare(`
	UPDATE profile
    SET avatar = ?, birthdate = ?, gender = ?, about = ?, first_name = ?, last_name = ?, cover_photo = ?, phone_number = ?
    WHERE user_id = ?;
	`)
	if err != nil {
		return fmt.Errorf("UpdateUserProfile: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(request.Avatar, request.Birthdate, request.Gender, request.About, request.FirstName, request.LastName, request.CoverPhoto, request.PhoneNumber, userID)
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

	responseBody := UpdateProfileResponse{
		UserID: userID,
	}
	return response.ApiResponse(http.StatusOK, responseBody)
}

func main() {
	lambda.Start(handler)
}
