package main

import (
	"github.com/aws/aws-lambda-go/events"
	"selector/auth/dummy"
	"testing"
)

func TestHandler(t *testing.T) {
	testBody := "token=test&team_id=test&team_domain=test-test&channel_id=test&channel_name=directmessage&user_id=test&user_name=test&command=/selector&text=hoge&api_app_id=test&is_enterprise_install=false"

	response := dummy.SetDummySignatures(testBody)
	testRequest := events.APIGatewayProxyRequest{Headers: response, Body: testBody}

	Handler(testRequest)
}
