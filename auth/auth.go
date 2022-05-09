package auth

import (
	"log"
	"net/http"
	"os"

	"github.com/slack-go/slack"
)

func SecretsVerify(body string, headers map[string]string) error {
	httpHeaders := convertHeaders(headers)

	verifier, err := slack.NewSecretsVerifier(httpHeaders, os.Getenv("SLACK_SIGNING_SECRET"))
	if err != nil {
		log.Println(err)
		return err
	}

	verifier.Write([]byte(body))

	return verifier.Ensure()
}

func convertHeaders(headers map[string]string) http.Header {
	h := http.Header{}
	for key, value := range headers {
		h.Set(key, value)
	}
	return h
}
