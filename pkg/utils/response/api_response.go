package response

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

func ApiResponse(status int, body interface{}) (events.APIGatewayProxyResponse, error) {
	resp := events.APIGatewayProxyResponse{
		Headers: map[string]string{
			"Access-Control-Allow-Origin":  "*",
			"Access-Control-Allow-Headers": "Content-Type, Access-Control-Allow-Headers, X-Requested-With",
			"Access-Control-Allow-Methods": "GET, POST, OPTIONS, PUT, DELETE",
		},
		StatusCode: status,
	}

	responseBody, err := json.Marshal(body)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	resp.Body = string(responseBody)
	return resp, nil
}
