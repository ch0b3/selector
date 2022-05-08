package filtering

import "strings"

// bodyからテキストを抽出する
func Filter_text(base_str string) string {
	params := strings.Split(base_str, "&")
	prefix_func := func(str *string) bool {
		pattern := "text="
		answer := strings.HasPrefix(*str, pattern)
		if answer {
			*str = strings.Replace(*str, pattern, "", -1)
		}
		return answer
	}
	texts := select_map(prefix_func, params)
	text := texts[0]
	return text
}

// sから、f(x)==true なxを返す
func select_map(f func(s *string) bool, strs []string) []string {
	res := make([]string, 0)
	for _, str := range strs {
		if f(&str) {
			res = append(res, str)
		}
	}
	return res
}
