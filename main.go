package main

import (
	"log"
	"net/url"
	"strings"

	"selector/filtering"
	"selector/selection"

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

	text := filtering.Filter_text(str_body)

	params := selection.TextToStruct(text)
	log.Println(params)

	selected := selection.SelectByCount(&params)
	log.Println(selected)

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       strings.Join(selected, "\n"),
	}, nil
}

func main() {
	lambda.Start(Handler)
}
