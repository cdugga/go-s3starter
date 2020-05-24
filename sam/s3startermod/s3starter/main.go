// main.go
package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/events"
	"fmt"
)

func hello() (string, error) {
	// return "Hello ƛ!", nil
	  // some error occured
	return events.APIGatewayProxyResponse{
	Body:       fmt.Sprintf("Hello  ƛ!, %v", "colin"),
	StatusCode: 200,
	}, nil
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(hello)
}