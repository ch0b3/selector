package main

import (
	"log"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"math/rand"
	"time"

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

	params := text_to_struct(text)
	log.Println(params)

	selected := select_by_count(&params)
	log.Println(selected)

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

// TODO: 場所を整理する
var rep = regexp.MustCompile(`\<.*?\>`)

type Params struct {
	members []string
	count   int
}

func text_to_struct(text string) Params {
	response := Params{members: make([]string, 0), count: 0}

	// <>があったら中を取り出す
	results := rep.FindAllStringSubmatch(text, -1)
	for _, member := range results {
		// <>を削除
		s := strings.ReplaceAll(member[0], "<", "")
		s = strings.ReplaceAll(s, ">", "")
		response.members = append(response.members, s)
	}

	// membersを削る
	text = rep.ReplaceAllString(text, "")
	text = strings.TrimSpace(text)

	// 残りを半角文字で区切る
	texts := strings.Split(text, " ")

	// TODO: エラーハンドリング
	if num, err := strconv.Atoi(texts[0]); err == nil {
		response.count = num
	}

	return response
}

func select_by_count(params *Params) []string {
	selected := make([]string, 0)

	for i := 0; i < params.count; i++ {
		rand.Seed(time.Now().UnixNano())
		i := rand.Intn(len(params.members))
		
		selected = append(selected, params.members[i])
		// 選ばれたものはmembersから削除する
		params.members = append(params.members[:i], params.members[i+1:]...)
	}

	return selected
}
