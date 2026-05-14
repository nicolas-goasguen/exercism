package microblog

func Truncate(phrase string) string {
	runes := make([]rune, 0, 5)
	for _, r := range phrase {
		if len(runes) >= 5 {
			break
		}
		runes = append(runes, r)
	}
	return string(runes)
}
