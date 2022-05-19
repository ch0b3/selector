package auth

import (
	"log"
	"selector/auth/dummy"
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
