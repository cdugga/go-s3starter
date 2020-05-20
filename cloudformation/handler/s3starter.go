package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/sbstjn/go-lambda-example/repository"
)


func handleRequest(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	
	return events.APIGatewayProxyResponse{Body: string("success"), StatusCode: 200}, nil
}

func main() {
	lambda.Start(handleRequest)
}