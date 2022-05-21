package main

import (
	"encoding/json"
	"log"
	"net/url"
	"strings"

	"github.com/ch0b3/selector/auth"
	"github.com/ch0b3/selector/filtering"
	"github.com/ch0b3/selector/selection"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type ResponseBody struct {
	ResponseType string `json:"response_type"`
	Text         string `json:"text"`
}

// Reference: https://github.com/aws/aws-lambda-go/blob/main/events/lambda_function_urls.go
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if err := auth.SecretsVerify(request.Body, request.Headers); err != nil {
		return buildResponse("認証情報に間違いがあります。", err)
	}

	body, err := url.QueryUnescape(request.Body)
	if err != nil {
		return buildResponse("処理が失敗しました。", err)
	}
	str_body := string(body)
	log.Println(str_body)

	text := filtering.FilterText(str_body)

	params, err := selection.TextToStruct(text)
	if err != nil {
		return buildResponse("処理が失敗しました。", err)
	}
	log.Println(params)

	selected := selection.SelectByCount(&params)
	log.Println(selected)

	responseBody := ResponseBody{
		ResponseType: "in_channel",
		Text:         strings.Join(selected, "\n"),
	}
	jsonData, _ := json.Marshal(responseBody)

	return buildResponse(string(jsonData), nil)
}

func buildResponse(messageBody string, err error) (events.APIGatewayProxyResponse, error) {
	log.Println(err)
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       messageBody,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
