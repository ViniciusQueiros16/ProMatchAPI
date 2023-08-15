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

type CreatePostRequest struct {
	UserID        int64  `json:"user_id"`
	Message       string `json:"message"`
	ImageURL      string `json:"image_url"`
	CommunityType string `json:"community_type"`
}

type CreatePostResponse struct {
	PostID int64 `json:"post_id"`
}

func CreatePost(db *sql.DB, req structs.Post) (int64, error) {
	stmt, err := db.Prepare("INSERT INTO posts (user_id, message, image_url, communityType) VALUES (?, ?, ?, ?)")
	if err != nil {
		return 0, fmt.Errorf("CreatePost: %w", err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(req.UserID, req.Message, req.ImageURL, req.CommunityType)
	if err != nil {
		return 0, fmt.Errorf("CreatePost: %w", err)
	}

	postID, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("CreatePost: %w", err)
	}

	return postID, nil
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	db, err := database.InitDB()
	if err != nil {
		return response.ApiResponse(http.StatusInternalServerError, structs.ErrorBody{
			ErrorMsg: aws.String(err.Error()),
		})
	}
	defer database.CloseDB()

	var req structs.Post
	err = json.Unmarshal([]byte(request.Body), &req)
	if err != nil {
		return response.ApiResponse(http.StatusBadRequest, structs.ErrorBody{
			ErrorMsg: aws.String(err.Error()),
		})
	}

	postID, err := CreatePost(db, req)
	if err != nil {
		return response.ApiResponse(http.StatusInternalServerError, structs.ErrorBody{
			ErrorMsg: aws.String(err.Error()),
		})
	}

	responseBody := CreatePostResponse{
		PostID: postID,
	}

	return response.ApiResponse(http.StatusCreated, responseBody)
}

func main() {
	lambda.Start(handler)
}
