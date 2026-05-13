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
	firstMonthDay := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)

	offsetToFirst := positiveOffset(firstMonthDay.Weekday(), wDay)
	firstWeekDay := 1 + offsetToFirst

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
		lastMonthDay := time.Date(year, month+1, 0, 0, 0, 0, 0, time.UTC)
		offsetToLast := negativeOffset(lastMonthDay.Weekday(), wDay)
		return lastMonthDay.Day() - offsetToLast
	case Teenth:
		firstTeenthDay := time.Date(year, month, 13, 0, 0, 0, 0, time.UTC)
		offsetToTeenth := positiveOffset(firstTeenthDay.Weekday(), wDay)
		return 13 + offsetToTeenth
	default:
		panic("invalid week schedule")
	}
}
