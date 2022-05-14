package main

import (
	"encoding/json"
	"log"
	"net/url"
	"strings"

	"selector/auth"
	"selector/filtering"
	"selector/selection"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type ResponseBody struct {
	ResponseType string `json:"response_type"`
	Text         string `json:"text"`
}

// Reference: https://github.com/aws/aws-lambda-go/blob/main/events/lambda_function_urls.go
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if auth.SecretsVerify(request.Body, request.Headers) != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 200,
			Body:       "認証情報に間違いがあります。",
		}, nil
	}

	body, err := url.QueryUnescape(request.Body)
	if err != nil {
		log.Fatal(err)
	}
	str_body := string(body)
	log.Println(str_body)

	text := filtering.FilterText(str_body)

	params := selection.TextToStruct(text)
	log.Println(params)

	selected := selection.SelectByCount(&params)
	log.Println(selected)

	responseBody := ResponseBody{
		ResponseType: "in_channel",
		Text:         strings.Join(selected, "\n"),
	}

	jsonData, _ := json.Marshal(responseBody)

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(jsonData),
	}, nil
}

func main() {
	lambda.Start(Handler)
}
