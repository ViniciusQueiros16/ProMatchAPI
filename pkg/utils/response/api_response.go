package response

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
)

func ApiResponse(status int, body interface{}) (events.APIGatewayProxyResponse, error) {
	resp := events.APIGatewayProxyResponse{
		Headers: map[string]string{
			"Content-Type":                 "application/json",
			"Access-Control-Allow-Origin":  "*",
			"Access-Control-Allow-Headers": "Content-Type, Access-Control-Allow-Headers, X-Requested-With",
			"Access-Control-Allow-Methods": "GET, POST, OPTIONS, PUT, DELETE",
		},
		StatusCode: status,
	}

	stringBody, _ := json.Marshal(body)
	resp.Body = string(stringBody)
	fmt.Println(resp)
	return resp, nil
}
