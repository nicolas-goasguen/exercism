package meetup

import "time"

// Define the WeekSchedule type here.
type WeekSchedule int

const (
	First WeekSchedule = iota
	Second
	Third
	Fourth
	Last
	Teenth
)

func positiveOffset(from, to time.Weekday) int {
	return (7 + int(to) - int(from)) % 7
}

func negativeOffset(from, to time.Weekday) int {
	return (7 + int(from) - int(to)) % 7
}

func Day(wSched WeekSchedule, wDay time.Weekday, month time.Month, year int) int {
	firstDay := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)

	offset := positiveOffset(firstDay.Weekday(), wDay)
	firstWeekDay := 1 + offset

	switch wSched {
	case First:
		return firstWeekDay
	case Second:
		return firstWeekDay + 7
	case Third:
		return firstWeekDay + 14
	case Fourth:
		return firstWeekDay + 21
	case Last:
		lastDay := time.Date(year, month+1, 0, 0, 0, 0, 0, time.UTC)
		lastOffset := negativeOffset(lastDay.Weekday(), wDay)
		return lastDay.Day() - lastOffset
	case Teenth:
		firstTeenthDay := time.Date(year, month, 13, 0, 0, 0, 0, time.UTC)
		teenthOffset := positiveOffset(firstTeenthDay.Weekday(), wDay)
		return 13 + teenthOffset
	default:
		return 0
	}
}
