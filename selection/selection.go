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
	Mode    string
}

type Room struct {
	Members []string
	Count   int
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

func SelectMembersByMode(params *Params) []*Room {
	rooms := make([]*Room, 0)

	if params.Mode == "split" {
		// params.Members / params.Count
		quotient := len(params.Members) / params.Count
		remainder := len(params.Members) % params.Count

		var rooms []Room
		room := Room{Members: make([]string, 0), Count: (quotient + 1)}
		// rooms = append(rooms, )
		
		// (商+1)人のroom * 余りの数 と (商)人のroom * (count - 余りの数)
	} else {
		room := Room{Members: make([]string, 0), Count: params.Count}
		rooms = append(rooms, &room)
		SelectByCount(&room, params.Members)
	}

	return rooms
}

func SelectByCount(room *Room, candidates []string) (*Room, []string) {
	for i := 0; i < room.Count; i++ {
		rand.Seed(time.Now().UnixNano())
		i := rand.Intn(len(candidates))

		selectedMember := formattingMember(candidates[i])
		room.Members = append(room.Members, selectedMember)
		// 選ばれたものはMembersから削除する
		candidates = append(candidates[:i], candidates[i+1:]...)
	}

	return room, candidates
}

func formattingMember(selectedMember string) string {
	pattern := "@"
	if strings.HasPrefix(selectedMember, pattern) {
		selectedMember = "<" + selectedMember + ">"
	}
	return selectedMember
}
