package filtering

import (
	"testing"
)

func TestFilterText(t *testing.T) {
	testString := "token=test&team_id=test&team_domain=test-test&channel_id=test&channel_name=directmessage&user_id=test&user_name=test&command=/selector&text=hoge&api_app_id=test&is_enterprise_install=false"
	got := FilterText(testString)
	want := "hoge"
	if got != want {
		t.Errorf("Fail")
	}
}
