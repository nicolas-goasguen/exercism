package anagram

import (
	"strings"
)

func Detect(subject string, candidates []string) []string {
	var anagrams []string

	for _, c := range candidates {
		if isAnagram(subject, c) {
			anagrams = append(anagrams, c)
		}
	}

	return anagrams
}

func isAnagram(a, b string) bool {
	a = strings.ToLower(a)
	b = strings.ToLower(b)

	if a == b {
		return false
	}

	runesCounter := make(map[rune]int)

	for _, r := range a {
		runesCounter[r]++
	}

	for _, r := range b {
		if _, ok := runesCounter[r]; !ok {
			return false
		}
		runesCounter[r]--
		if runesCounter[r] == 0 {
			delete(runesCounter, r)
		}
	}

	return len(runesCounter) == 0
}
