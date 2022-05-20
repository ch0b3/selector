package main

import (
	"encoding/json"
	"net/url"
	"strings"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/ch0b3/selector/auth/dummy"
)

func TestHandler(t *testing.T) {
	testBody := "token=test&team_id=test&team_domain=test-test&channel_id=test&channel_name=directmessage&user_id=test&user_name=test&command=/selector&text=[member1][member2][member3] 2&api_app_id=test&is_enterprise_install=false"
	escapedBody := url.QueryEscape(testBody)
	response := dummy.SetDummySignatures(escapedBody)

	testRequest := events.APIGatewayProxyRequest{Headers: response, Body: escapedBody}

	got, _ := Handler(testRequest)

	if got.StatusCode != 200 {
		t.Errorf("Fail")
	}

	var responseBody ResponseBody
	err := json.Unmarshal([]byte(got.Body), &responseBody)
	if err != nil {
		t.Errorf("Fail")
	}
	gotMembers := strings.Split(responseBody.Text, "\n")

	if responseBody.ResponseType != "in_channel" {
		t.Errorf("Fail")
	}

	for _, m := range gotMembers {
		if m != "member1" && m != "member2" && m != "member3" {
			t.Errorf("Fail")
		}
	}
}
