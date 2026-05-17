package piglatin

import (
	"fmt"
	"regexp"
	"strings"
)

const (
	insensitiveFlag = `(?i)`
	startFlag       = `^`
	vowels          = `aeiou`
	consonants      = `bcdfghjklmnpqrstvwxyz`
	consonantsNoY   = `bcdfghjklmnpqrstvwxz`
	letters         = `a-z`
)

var (
	matchRule1 = fmt.Sprintf(`(([%s]+|xr|yt)[%s]*)`, vowels, letters)
	matchRule2 = fmt.Sprintf(`([%s]+)`, consonants)
	matchRule3 = fmt.Sprintf(`([%s]*qu)`, consonants)
	matchRule4 = fmt.Sprintf(`([%s]+)y`, consonantsNoY)

	rule1 = regexp.MustCompile(insensitiveFlag + startFlag + matchRule1)
	rule2 = regexp.MustCompile(insensitiveFlag + startFlag + matchRule2)
	rule3 = regexp.MustCompile(insensitiveFlag + startFlag + matchRule3)
	rule4 = regexp.MustCompile(insensitiveFlag + startFlag + matchRule4)

	rules = []*regexp.Regexp{rule1, rule4, rule3, rule2}
)

func applyRule(word string, rule *regexp.Regexp) (string, bool) {
	match := rule.FindStringSubmatch(word)
	if match == nil {
		return "", false
	}
	return word[len(match[1]):] + match[1] + "ay", true
}

func Sentence(sentence string) string {
	words := strings.Fields(sentence)
	result := make([]string, 0, len(words))
	for _, w := range words {
		translated := w
		for _, r := range rules {
			s, ok := applyRule(w, r)
			if ok {
				translated = s
				break
			}
		}
		result = append(result, translated)
	}
	return strings.Join(result, " ")
}
