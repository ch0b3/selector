package selection_test

import (
	"github.com/ch0b3/selector/selection"
	"reflect"
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
