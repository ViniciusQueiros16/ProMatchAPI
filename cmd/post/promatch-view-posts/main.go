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

type PostWithUser struct {
	ID            int64  `json:"id"`
	UserID        int64  `json:"user_id"`
	Username      string `json:"username"`
	Avatar        string `json:"avatar"`
	Company       string `json:"company"`
	Message       string `json:"message"`
	ImageURL      string `json:"image_url"`
	CommunityType string `json:"communityType"`
	Timestamp     string `json:"timestamp"`
}

func FetchPaginatedPosts(db *sql.DB, page int) ([]PostWithUser, error) {
	var posts []PostWithUser

	postsPerPage := 5
	offset := (page - 1) * postsPerPage

	query := `
		SELECT
			po.id, po.user_id, u.username, p.avatar, p.company, po.message, po.image_url, po.communityType, po.timestamp
		FROM
			posts po
		JOIN
			users u ON po.user_id = u.id
		JOIN
			profile p ON po.user_id = p.user_id
		ORDER BY
			po.timestamp DESC
		LIMIT ? OFFSET ?;
	`

	rows, err := db.Query(query, postsPerPage, offset)
	if err != nil {
		return nil, fmt.Errorf("FetchPaginatedPostsWithUsers: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var post PostWithUser
		err := rows.Scan(
			&post.ID, &post.UserID, &post.Username, &post.Avatar, &post.Company, &post.Message,
			&post.ImageURL, &post.CommunityType, &post.Timestamp,
		)
		if err != nil {
			return nil, fmt.Errorf("FetchPaginatedPostsWithUsers: %w", err)
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	db, err := database.InitDB()
	if err != nil {
		return response.ApiResponse(http.StatusInternalServerError, structs.ErrorBody{
			ErrorMsg: aws.String(err.Error()),
		})
	}
	defer database.CloseDB()

	// Get page number from query string
	pageStr := request.QueryStringParameters["page"]
	page, _ := strconv.Atoi(pageStr)
	if page == 0 {
		page = 1
	}

	// Fetch paginated posts with user information
	paginatedPostsWithUsers, err := FetchPaginatedPosts(db, page)
	if err != nil {
		return response.ApiResponse(http.StatusBadRequest, structs.ErrorBody{
			ErrorMsg: aws.String(err.Error()),
		})
	}

	return response.ApiResponse(http.StatusOK, paginatedPostsWithUsers)
}

func main() {
	lambda.Start(handler)
}
