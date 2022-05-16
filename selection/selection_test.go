package selection_test

import (
	"reflect"
	"selector/selection"
	"testing"
)

func TestTextToStruct(t *testing.T) {
	text := "[member1][member2] 1"
	got := selection.TextToStruct(text)
	want := selection.Params{Members: []string{"member1", "member2"}, Count: 1}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Fail")
	}
}
