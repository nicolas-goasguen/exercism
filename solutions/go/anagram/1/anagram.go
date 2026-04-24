package anagram

import (
	"maps"
	"strings"
)

func Detect(subject string, candidates []string) []string {
	subLower := strings.ToLower(subject)
	subRunes := []rune(subLower)

	subRunesCount := make(map[rune]int, len(subRunes))
	for _, r := range subRunes {
		subRunesCount[r] += 1
	}

	var anagrams []string
	for _, can := range candidates {
		canLower := strings.ToLower(can)
		canRunes := []rune(canLower)

		if canLower == subLower {
			continue
		}

		canRunesCount := make(map[rune]int, len(subRunesCount))
		maps.Copy(canRunesCount, subRunesCount)

		canInvalid := false
		for _, r := range canRunes {
			rCount, ok := canRunesCount[r]
			if !ok || rCount <= 0 {
				canInvalid = true
				break
			}
			canRunesCount[r]--
		}
		for _, v := range canRunesCount {
			if v > 0 {
				canInvalid = true
				break
			}
		}
		if !canInvalid {
			anagrams = append(anagrams, can)
		}
	}
	return anagrams
}
