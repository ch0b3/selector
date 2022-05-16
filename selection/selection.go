package selection

import (
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Params struct {
	members []string
	count   int
}

var rep = regexp.MustCompile(`\[.*?\]`)

func TextToStruct(text string) Params {
	response := Params{members: make([]string, 0), count: 0}

	// []があったら中を取り出す
	results := rep.FindAllStringSubmatch(text, -1)
	for _, member := range results {
		// []を削除
		s := strings.ReplaceAll(member[0], "[", "")
		s = strings.ReplaceAll(s, "]", "")
		response.members = append(response.members, s)
	}

	// membersを削る
	text = rep.ReplaceAllString(text, "")
	text = strings.TrimSpace(text)

	// 残りを半角文字で区切る
	texts := strings.Split(text, " ")

	// TODO: エラーハンドリング
	if num, err := strconv.Atoi(texts[0]); err == nil {
		response.count = num
	}

	return response
}

func SelectByCount(params *Params) []string {
	selectedMembers := make([]string, 0)

	for i := 0; i < params.count; i++ {
		rand.Seed(time.Now().UnixNano())
		i := rand.Intn(len(params.members))

		selectedMember := formattingMember(params.members[i])
		selectedMembers = append(selectedMembers, selectedMember)
		// 選ばれたものはmembersから削除する
		params.members = append(params.members[:i], params.members[i+1:]...)
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
