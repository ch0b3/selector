package auth_test

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"log"
	"os"
	"selector/auth"
	"strconv"
	"testing"
	"time"
)

func TestSecretsVerify(t *testing.T) {
	os.Setenv("SLACK_SIGNING_SECRET", "8f742231b10e8888abcd99yyyzzz85a5")
	test_body := "token=testtesttesttest"
	now := strconv.FormatInt(time.Now().Unix(), 10)

	sig_basestring := "v0:" + now + ":" + test_body

	mac := hmac.New(sha256.New, []byte(os.Getenv("SLACK_SIGNING_SECRET")))
	mac.Write([]byte(sig_basestring))
	signature := hex.EncodeToString(mac.Sum(nil))

	test_signature := map[string]string{"X-Slack-Signature": "v0=" + signature, "X-Slack-Request-Timestamp": now}

	response := auth.SecretsVerify(test_body, test_signature)
	if response != nil {
		log.Fatal(response)
		t.Errorf("Fail")
	}
}
