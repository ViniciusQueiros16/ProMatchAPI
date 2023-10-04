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

type MatchResponse struct {
	MatchID int64 `json:"match_id"`
}

func CreateMatch(db *sql.DB, match structs.Matches) (int64, error) {
	stmt, err := db.Prepare(`
        INSERT INTO matches (user_id, matched_user_id)
        SELECT ?, ?
        WHERE NOT EXISTS (
            SELECT 1 FROM matches
            WHERE user_id = ? AND matched_user_id = ?
        )
    `)
	if err != nil {
		return 0, fmt.Errorf("CreateMatch: %w", err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(match.UserID, match.MatchedUserID, match.UserID, match.MatchedUserID)
	if err != nil {
		return 0, fmt.Errorf("CreateMatch: %w", err)
	}

	matchID, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("CreateMatch: %w", err)
	}

	return matchID, nil
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	db, err := database.InitDB()
	if err != nil {
		return response.ApiResponse(http.StatusInternalServerError, structs.ErrorBody{
			ErrorMsg: aws.String(err.Error()),
		})
	}
	defer database.CloseDB()

	var req structs.Matches

	err = json.Unmarshal([]byte(request.Body), &req)
	if err != nil {
		return response.ApiResponse(http.StatusBadRequest, structs.ErrorBody{
			ErrorMsg: aws.String(err.Error()),
		})
	}

	matchID, err := CreateMatch(db, req)
	if err != nil {
		return response.ApiResponse(http.StatusInternalServerError, structs.ErrorBody{
			ErrorMsg: aws.String(err.Error()),
		})
	}

	responseBody := MatchResponse{
		MatchID: matchID,
	}

	return response.ApiResponse(http.StatusCreated, responseBody)
}

func main() {
	lambda.Start(handler)
}
