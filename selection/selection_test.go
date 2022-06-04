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
	want := selection.Params{Members: []string{"member1", "member2"}, Count: 1, Mode: "default"}

	if err != nil {
		t.Errorf("Fail")
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Fail")
	}
}

func TestSelectMembersByMode(t *testing.T) {
	// defaultパターン
	params := selection.Params{Members: []string{"member1", "member2"}, Count: 1, Mode: "default"}
	rooms := selection.SelectMembersByMode(&params)

	if len(rooms[0].Members) != 1 {
		t.Errorf("Fail")
	}

	if rooms[0].Count != 1 {
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

	room := selection.Room{Members: make([]string, 0), Count: count}
	candidates := members
	afterRoom, remainingCandidates := selection.SelectByCount(&room, candidates)

	gotInterfaces := []interface{}{}
	for _, m := range afterRoom.Members {
		gotInterfaces = append(gotInterfaces, m)
	}

	paramsSet := set.NewSetFromSlice(inputInterfaces)
	gotSet := set.NewSetFromSlice(gotInterfaces)

	// inputとgotの積集合の数がcountになっていればOK
	if paramsSet.Intersect(gotSet).Cardinality() != count {
		t.Errorf("Fail")
	}

	// 与えたcandidatesの数からcountが引かれていればOK
	if len(remainingCandidates) != (len(members) - count) {
		t.Errorf("Fail")
	}
}
