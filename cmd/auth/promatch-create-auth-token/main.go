package authToken

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/promatch/pkg/database"
	"github.com/promatch/pkg/utils/response"
	"github.com/promatch/pkg/utils/token"
	"github.com/promatch/structs"
)

const tokenExpiryHours = 24

func CreateAuthToken(db *sql.DB, userID int64) (structs.AuthToken, error) {
	token, err := token.GenerateAuthToken(userID, tokenExpiryHours)
	if err != nil {
		return structs.AuthToken{}, fmt.Errorf("CreateAuthToken: %w", err)
	}

	expiresAt := time.Now().Add(time.Hour * tokenExpiryHours)

	stmt, err := db.Prepare("INSERT INTO auth_tokens(user_id, token, expires_at) VALUES (?, ?, ?)")
	if err != nil {
		return structs.AuthToken{}, fmt.Errorf("CreateAuthToken: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(userID, token, expiresAt)
	if err != nil {
		return structs.AuthToken{}, fmt.Errorf("CreateAuthToken: %w", err)
	}

	authToken := structs.AuthToken{
		Token:     token,
		ExpiresAt: expiresAt,
	}

	return authToken, nil
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	db, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer database.CloseDB()

	var user structs.Users
	err = json.Unmarshal([]byte(request.Body), &user)
	if err != nil {
		return response.ApiResponse(http.StatusBadRequest, structs.ErrorBody{
			ErrorMsg: aws.String(err.Error()),
		})
	}

	authToken, err := CreateAuthToken(db, user.ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Token created: %v\n", authToken.Token)

	return response.ApiResponse(http.StatusCreated, authToken)
}

func main() {
	lambda.Start(handler)
}
