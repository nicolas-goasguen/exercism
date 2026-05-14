package swiftscheduling

import (
	"strconv"
	"time"
)

const (
	now   = "NOW"
	asap  = "ASAP"
	eow   = "EOW"
	fmtDT = "2006-01-02T15:04:05"
)

var quarterEndMonth = map[int]time.Month{
	1: time.March,
	2: time.June,
	3: time.September,
	4: time.December,
}

func positiveWeekdayOffset(from, to time.Weekday) int {
	return (7 + int(to) - int(from)) % 7
}

func positiveWorkdayOffset(from time.Weekday) int {
	switch from {
	case time.Saturday:
		return 2
	case time.Sunday:
		return 1
	}
	return 0
}

func negativeWorkdayOffset(from time.Weekday) int {
	switch from {
	case time.Saturday:
		return 1
	case time.Sunday:
		return 2
	}
	return 0
}

func parseFixedDelivery(meetingDT time.Time, delivery string) (string, bool) {
	switch delivery {
	case now:
		return time.Date(meetingDT.Year(), meetingDT.Month(), meetingDT.Day(), meetingDT.Hour()+2, 0, 0, 0, meetingDT.Location()).Format(fmtDT), true
	case asap:
		if meetingDT.Hour() < 13 {
			return time.Date(meetingDT.Year(), meetingDT.Month(), meetingDT.Day(), 17, 0, 0, 0, meetingDT.Location()).Format(fmtDT), true
		} else {
			return time.Date(meetingDT.Year(), meetingDT.Month(), meetingDT.Day()+1, 13, 0, 0, 0, meetingDT.Location()).Format(fmtDT), true
		}
	case eow:
		weekD := meetingDT.Weekday()
		if weekD >= time.Monday && weekD <= time.Wednesday {
			offsetToFriday := positiveWeekdayOffset(weekD, time.Friday)
			return time.Date(meetingDT.Year(), meetingDT.Month(), meetingDT.Day()+offsetToFriday, 17, 0, 0, 0, meetingDT.Location()).Format(fmtDT), true
		} else if weekD >= time.Thursday && weekD <= time.Friday {
			offsetToSunday := positiveWeekdayOffset(weekD, time.Sunday)
			return time.Date(meetingDT.Year(), meetingDT.Month(), meetingDT.Day()+offsetToSunday, 20, 0, 0, 0, meetingDT.Location()).Format(fmtDT), true
		} else {
			panic("invalid meeting start day")
		}
	default:
		return "", false
	}
}

func parseMonthDelivery(meetingDT time.Time, delivery string) (string, bool) {
	length := len(delivery)
	if length < 2 || length > 3 || delivery[length-1] != 'M' {
		return "", false
	}
	deliveryM, err := strconv.Atoi(delivery[:length-1])
	if err != nil || deliveryM < 1 || deliveryM > 12 {
		return "", false
	}
	var deliveryDT time.Time
	if int(meetingDT.Month()) < deliveryM {
		deliveryDT = time.Date(meetingDT.Year(), time.Month(deliveryM), 1, 0, 0, 0, 0, meetingDT.Location())
	} else {
		deliveryDT = time.Date(meetingDT.Year()+1, time.Month(deliveryM), 1, 0, 0, 0, 0, meetingDT.Location())
	}
	offset := positiveWorkdayOffset(deliveryDT.Weekday())
	return time.Date(deliveryDT.Year(), deliveryDT.Month(), deliveryDT.Day()+offset, 8, 0, 0, 0, meetingDT.Location()).Format(fmtDT), true
}

func parseQuarterDelivery(meetingDT time.Time, delivery string) (string, bool) {
	length := len(delivery)
	if length != 2 || delivery[0] != 'Q' {
		return "", false
	}
	deliveryQ, err := strconv.Atoi(delivery[length-1:])
	if err != nil {
		return "", false
	}
	quarterLastMonth, ok := quarterEndMonth[deliveryQ]
	if !ok {
		return "", false
	}
	var deliveryTime time.Time
	if int(meetingDT.Month()) <= int(quarterLastMonth) {
		deliveryTime = time.Date(meetingDT.Year(), quarterLastMonth+1, 0, 0, 0, 0, 0, meetingDT.Location())
	} else {
		deliveryTime = time.Date(meetingDT.Year()+1, quarterLastMonth+1, 0, 0, 0, 0, 0, meetingDT.Location())
	}
	offset := negativeWorkdayOffset(deliveryTime.Weekday())
	return time.Date(deliveryTime.Year(), deliveryTime.Month(), deliveryTime.Day()-offset, 8, 0, 0, 0, meetingDT.Location()).Format(fmtDT), true
}

func DeliveryDate(start, delivery string) string {
	meetingDT, err := time.Parse(fmtDT, start)
	if err != nil {
		panic("invalid date format, must be: " + fmtDT)
	}
	if parsed, ok := parseFixedDelivery(meetingDT, delivery); ok {
		return parsed
	}
	if parsed, ok := parseMonthDelivery(meetingDT, delivery); ok {
		return parsed
	}
	if parsed, ok := parseQuarterDelivery(meetingDT, delivery); ok {
		return parsed
	}
	return ""
}
