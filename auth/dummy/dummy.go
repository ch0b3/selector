package dummy

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"os"
	"strconv"
	"time"
)

func SetDummySignatures(testBody string) map[string]string {
	os.Setenv("SLACK_SIGNING_SECRET", "8f742231b10e8888abcd99yyyzzz85a5")

	now := strconv.FormatInt(time.Now().Unix(), 10)

	sigBasestring := "v0:" + now + ":" + testBody

	mac := hmac.New(sha256.New, []byte(os.Getenv("SLACK_SIGNING_SECRET")))
	mac.Write([]byte(sigBasestring))
	signature := hex.EncodeToString(mac.Sum(nil))

	return map[string]string{"X-Slack-Signature": "v0=" + signature, "X-Slack-Request-Timestamp": now}
}
