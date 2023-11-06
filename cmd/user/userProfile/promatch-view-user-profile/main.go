package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/promatch/pkg/database"
	"github.com/promatch/pkg/utils/response"
	"github.com/promatch/pkg/utils/token"
	"github.com/promatch/structs"
)

func FetchUserProfile(db *sql.DB, id int64) (structs.UserProfile, error) {
	var userProfile structs.UserProfile

	query := `
		SELECT 
			u.id,
			u.username,
			u.email,
			u.user_type_id,
			u.verified,
			u.privacy_accepted,
			p.first_name,
			p.last_name,
			p.avatar,
			p.cover_photo,
			p.phone_number,
			p.birthdate,
			p.gender,
			p.about,
			ua.country,
			ua.street_address,
			ua.city,
			ua.state,
			ua.postal_code,
			n.comments,
			n.candidates,
			n.offers,
			n.sms_delivery_option
		FROM users u
		LEFT JOIN profile p ON u.id = p.user_id
		LEFT JOIN user_addresses ua ON u.id = ua.user_id
		LEFT JOIN notifications n ON u.id = n.user_id
		WHERE u.id = ? LIMIT 1;
	`

	row := db.QueryRow(query, id)

	var birthdate sql.NullTime

	if err := row.Scan(
		&userProfile.UserID, &userProfile.Username, &userProfile.Email,
		&userProfile.UserTypeID, &userProfile.Verified, &userProfile.PrivacyAccepted, &userProfile.FirstName, &userProfile.LastName,
		&userProfile.Avatar, &userProfile.CoverPhoto, &userProfile.PhoneNumber, &birthdate,
		&userProfile.Gender, &userProfile.About,
		&userProfile.Country, &userProfile.StreetAddress, &userProfile.City, &userProfile.State, &userProfile.PostalCode,
		&userProfile.Comments, &userProfile.Candidates, &userProfile.Offers, &userProfile.SMSDeliveryOption,
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

	result, err := FetchUserProfile(db, userID)
	if err != nil {
		return response.ApiResponse(http.StatusBadRequest, structs.ErrorBody{
			ErrorMsg: aws.String(err.Error()),
		})
	}

	return response.ApiResponse(http.StatusOK, result)
}
func main() {
	lambda.Start(handler)
}
