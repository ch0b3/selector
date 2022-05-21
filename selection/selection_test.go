package selection_test

import (
	"github.com/ch0b3/selector/selection"
	set "github.com/deckarep/golang-set"
	"reflect"
	"testing"
)

func TestTextToStruct(t *testing.T) {
	text := "[member1][member2] 1"
	got, err := selection.TextToStruct(text)
	want := selection.Params{Members: []string{"member1", "member2"}, Count: 1}

	if err != nil {
		t.Errorf("Fail")
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Fail")
	}
}

func TestSelectByCount(t *testing.T) {
	members := []string{"member1", "member2", "member3"}
	count := 2

	inputInterfaces := []interface{}{}
	for _, m := range members {
		inputInterfaces = append(inputInterfaces, m)
	}

	params := selection.Params{Members: members, Count: count}
	gotMembers := selection.SelectByCount(&params)

	gotInterfaces := []interface{}{}
	for _, m := range gotMembers {
		gotInterfaces = append(gotInterfaces, m)
	}

	paramsSet := set.NewSetFromSlice(inputInterfaces)
	gotSet := set.NewSetFromSlice(gotInterfaces)

	// inputとgotの積集合の数がcountになっていればOK
	if paramsSet.Intersect(gotSet).Cardinality() != count {
		t.Errorf("Fail")
	}
}
