package main

import (
	"log"
	"net/url"
	"strings"

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

	text := filter_text(str_body)
	log.Println(text)

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "ok",
	}, nil
}

func main() {
	lambda.Start(Handler)
}

// bodyからテキストを抽出する
func filter_text(base_str string) string {
	params := strings.Split(base_str, "&")
	prefix_func := func(str *string) bool {
		pattern := "text="
		answer := strings.HasPrefix(*str, pattern)
		if answer {
			*str = strings.Replace(*str, pattern, "", -1)
		}
		return answer
	}
	texts := select_map(prefix_func, params)
	text := texts[0]
	return text
}

// sから、f(x)==true なxを返す
func select_map(f func(s *string) bool, strs []string) []string {
	res := make([]string, 0)
	for _, str := range strs {
		if f(&str) {
			res = append(res, str)
		}
	}
	return res
}
