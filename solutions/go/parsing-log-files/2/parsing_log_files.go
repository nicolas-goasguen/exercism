package parsinglogfiles

import (
	"regexp"
	"slices"
)

var (
	reValidLine = regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	reSplitLine = regexp.MustCompile(`<[~*=-]*>`)
	rePassword  = regexp.MustCompile(`(?i)".*password.*"`)
	reEndOfLine = regexp.MustCompile(`end-of-line\d+`)
	reUserName  = regexp.MustCompile(`User\s+([a-zA-Z]*\d*)`)
)

func IsValidLine(text string) bool {
	return reValidLine.MatchString(text)
}

func SplitLogLine(text string) []string {
	return reSplitLine.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	count := 0
	for _, l := range lines {
		if rePassword.MatchString(l) {
			count++
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	return reEndOfLine.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	cpy := slices.Clone(lines)
	for i, l := range cpy {
		userName := reUserName.FindStringSubmatch(l)
		if len(userName) > 1 {
			cpy[i] = "[USR] " + userName[1] + " " + cpy[i]
		}
	}
	return cpy
}
