package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/promatch/pkg/utils/uploadS3"
	"github.com/promatch/structs"
)

type ResponseToSend events.APIGatewayProxyResponse

func Handler(request events.APIGatewayProxyRequest) (ResponseToSend, error) {

	bodyRequest := &structs.ImageRequestBody{}
	err := json.Unmarshal([]byte(request.Body), &bodyRequest)

	if err != nil {
		return ResponseToSend{Body: err.Error(), StatusCode: 404}, nil
	}

	resp := uploadS3.ImageUpload(bodyRequest.FileName, bodyRequest.Body)

	response, err := json.Marshal(&resp)
	if err != nil {
		return ResponseToSend{Body: err.Error(), StatusCode: 404}, nil
	}

	respToSend := ResponseToSend{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(response),
	}

	return respToSend, nil
}

func main() {
	lambda.Start(Handler)
}
