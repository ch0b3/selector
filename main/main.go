package main

import (
	"encoding/json"
	"log"
	"net/url"
	"strconv"
	"strings"

	"github.com/ch0b3/selector/auth"
	"github.com/ch0b3/selector/filtering"
	"github.com/ch0b3/selector/selection"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type ResponseBody struct {
	ResponseType string `json:"response_type"`
	Blocks       []Block
}

type Block struct {
	Type   string `json:"type"`
	Fields []Field
}

type Field struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

// Reference: https://github.com/aws/aws-lambda-go/blob/main/events/lambda_function_urls.go
// Reference: https://github.com/aws/aws-lambda-go/blob/main/events/README_ApiGatewayEvent.md
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

	rooms := selection.SelectMembersByMode(&params)
	log.Println(rooms)

	responseBody := ResponseBody{
		ResponseType: "in_channel",
		Blocks:       serializeRooms(rooms),
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

func serializeRooms(rooms []*selection.Room) []Block {
	var text string
	var blocks []Block

	for idx, room := range rooms {
		text = ""
		roomStrings := []string{"*", strconv.Itoa(idx + 1), ".*\n", strings.Join(room.Members, "\n")}
		text += strings.Join(roomStrings, "")

		var fields []Field
		fields = append(fields, Field{Type: "mrkdwn", Text: text})
		blocks = append(blocks, Block{Type: "section", Fields: fields})
	}

	return blocks
}

func main() {
	lambda.Start(Handler)
}
