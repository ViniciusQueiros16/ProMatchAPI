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

type UpdateProfileResponse struct {
	UserID int64 `json:"userID"`
}

func UpdateUserProfile(db *sql.DB, userID int64, request structs.UpdateProfileRequest) error {
	// Update the 'profile' table
	profileStmt, err := db.Prepare(`
		UPDATE profile
		SET avatar = ?, birthdate = ?, gender = ?, about = ?, first_name = ?, last_name = ?, cover_photo = ?, phone_number = ?
		WHERE user_id = ?;
	`)
	if err != nil {
		return fmt.Errorf("UpdateUserProfile (profile): %w", err)
	}
	defer profileStmt.Close()

	_, err = profileStmt.Exec(request.Avatar, request.Birthdate, request.Gender, request.About, request.FirstName, request.LastName, request.CoverPhoto, request.PhoneNumber, userID)
	if err != nil {
		return fmt.Errorf("UpdateUserProfile (profile): %w", err)
	}

	// Update the'users' table
	usersStmt, err := db.Prepare(`
		UPDATE users
		SET username = ?, email = ?, user_type_id = ?, privacy_accepted = ?
		WHERE id = ?;
	`)
	if err != nil {
		return fmt.Errorf("UpdateUserProfile (users): %w", err)
	}
	defer usersStmt.Close()

	_, err = usersStmt.Exec(request.Username, request.Email, request.UserTypeID, request.PrivacyAccepted, userID)
	if err != nil {
		return fmt.Errorf("UpdateUserProfile (users): %w", err)
	}

	// Update the 'user_addresses' table
	userAddressesStmt, err := db.Prepare(`
		UPDATE user_addresses
		SET country = ?, street_address = ?, city = ?, state = ?, postal_code = ?
		WHERE user_id = ?;
	`)
	if err != nil {
		return fmt.Errorf("UpdateUserProfile (user_addresses): %w", err)
	}
	defer userAddressesStmt.Close()

	_, err = userAddressesStmt.Exec(request.Country, request.StreetAddress, request.City, request.State, request.PostalCode, userID)
	if err != nil {
		return fmt.Errorf("UpdateUserProfile (user_addresses): %w", err)
	}

	// Update the 'notifications' table
	notificationsStmt, err := db.Prepare(`
		UPDATE notifications
		SET comments = ?, candidates = ?, offers = ?, sms_delivery_option = ?
		WHERE user_id = ?;
	`)
	if err != nil {
		return fmt.Errorf("UpdateUserProfile (notifications): %w", err)
	}
	defer notificationsStmt.Close()

	_, err = notificationsStmt.Exec(request.Comments, request.Candidates, request.Offers, request.SMSDeliveryOption, userID)
	if err != nil {
		return fmt.Errorf("UpdateUserProfile (notifications): %w", err)
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

	var updateRequest structs.UpdateProfileRequest
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
