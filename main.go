package main

import (
	"log"
	"net/url"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Reference: https://github.com/aws/aws-lambda-go/blob/main/events/lambda_function_urls.go
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	body, err := url.QueryUnescape(request.Body)
	if err != nil {
		log.Fatal(err)
	}
	str_body := string(body)
	log.Println(str_body)

	// TODO: bodyからテキストを抽出する

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body: "ok",
	}, nil
}

func main() {
	lambda.Start(Handler)
}
