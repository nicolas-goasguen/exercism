package booking

import (
	"fmt"
	"time"
)

var openingDate = time.Date(2012, 9, 15, 0, 0, 0, 0, time.UTC)

const (
	scheduleInLayout     = "1/2/2006 15:04:05"
	hasPassedInLayout    = "January 2, 2006 15:04:05"
	isAfternoonInLayout  = "Monday, January 2, 2006 15:04:05"
	descriptionInLayout  = scheduleInLayout
	descriptionOutLayout = "You have an appointment on Monday, January 2, 2006, at 15:04."
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	parsed, err := time.Parse(scheduleInLayout, date)
	if err != nil {
		panic("invalid schedule date format")
	}
	return parsed
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	parsed, err := time.Parse(hasPassedInLayout, date)
	if err != nil {
		panic(fmt.Sprintf("invalid 'has passed' date format: %s", date))
	}
	return parsed.Before(time.Now().In(openingDate.Location()))
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	parsed, err := time.Parse(isAfternoonInLayout, date)
	if err != nil {
		panic(fmt.Sprintf("invalid 'is afternoon' date format: %s", date))
	}
	hour := parsed.Hour()
	return hour >= 12 && hour < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	parsed, err := time.Parse(descriptionInLayout, date)
	if err != nil {
		panic(fmt.Sprintf("invalid 'description' date format: %s", date))
	}
	return parsed.Format(descriptionOutLayout)
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	year := time.Now().In(openingDate.Location()).Year() // avoid error in timezone
	return time.Date(year, openingDate.Month(), openingDate.Day(), 0, 0, 0, 0, openingDate.Location())
}
