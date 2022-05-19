package auth

import (
	"github.com/ch0b3/selector/auth/dummy"
	"log"
	"testing"
)

func TestSecretsVerify(t *testing.T) {
	testBody := "token=testtesttest"

	test_signature := dummy.SetDummySignatures(testBody)
	response := SecretsVerify(testBody, test_signature)
	if response != nil {
		log.Fatal(response)
		t.Errorf("Fail")
	}
}
