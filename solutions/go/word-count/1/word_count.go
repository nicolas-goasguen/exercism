package wordcount

import (
	"strings"
	"unicode"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	freq := make(map[string]int)

	phraseSeq := strings.FieldsFuncSeq(phrase, func(r rune) bool {
		return !unicode.IsLetter(r) && r != '\'' && !unicode.IsDigit(r)
	})

	for w := range phraseSeq {
		w = strings.ToLower(strings.Trim(w, "'"))
		if len(w) > 0 {
			freq[w] += 1
		}
	}
	return freq
}
