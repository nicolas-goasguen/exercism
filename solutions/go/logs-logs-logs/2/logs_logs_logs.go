package logs

import (
	"strings"
	"unicode/utf8"
)

var appByRune = map[rune]string{
	'❗': "recommendation",
	'🔍': "search",
	'☀': "weather",
}

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, r := range log {
		if app, ok := appByRune[r]; ok {
			return app
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	return strings.Map(func(r rune) rune {
		if r == oldRune {
			return newRune
		}
		return r
	}, log)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
