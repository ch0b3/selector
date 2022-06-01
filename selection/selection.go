package selection

import (
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Params struct {
	Members []string
	Count   int
	Mode string
}

var rep = regexp.MustCompile(`\[.*?\]`)

func TextToStruct(text string) (Params, error) {
	response := Params{Members: make([]string, 0), Count: 0, Mode: "default"}

	// []があったら中を取り出す
	results := rep.FindAllStringSubmatch(text, -1)
	for _, member := range results {
		// []を削除
		s := strings.ReplaceAll(member[0], "[", "")
		s = strings.ReplaceAll(s, "]", "")
		response.Members = append(response.Members, s)
	}

	// membersを削る
	text = rep.ReplaceAllString(text, "")
	text = strings.TrimSpace(text)

	// 残りを半角文字で区切る
	texts := strings.Split(text, " ")

	for _, text := range texts {
		if text == "--split" {
			response.Mode = "split"
		}
	}

	if num, err := strconv.Atoi(texts[0]); err == nil {
		response.Count = num
	} else {
		return Params{}, err
	}

	return response, nil
}

func SelectByCount(params *Params) []string {
	selectedMembers := make([]string, 0)

	for i := 0; i < params.Count; i++ {
		rand.Seed(time.Now().UnixNano())
		i := rand.Intn(len(params.Members))

		selectedMember := formattingMember(params.Members[i])
		selectedMembers = append(selectedMembers, selectedMember)
		// 選ばれたものはMembersから削除する
		params.Members = append(params.Members[:i], params.Members[i+1:]...)
	}

	return selectedMembers
}

func formattingMember(selectedMember string) string {
	pattern := "@"
	if strings.HasPrefix(selectedMember, pattern) {
		selectedMember = "<" + selectedMember + ">"
	}
	return selectedMember
}
