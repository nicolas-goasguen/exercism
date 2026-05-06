package bafflingbirthdays

import (
	"math/rand"
	"strconv"
	"time"
)

var timeFormat = "2006-01-02"
var timeStart, _ = time.Parse(timeFormat, "2026-01-01")

const totalDays = 365

func SharedBirthday(dates []time.Time) bool {
	var days = make(map[string]struct{})
	for _, d := range dates {
		date := d.Month().String() + strconv.Itoa(d.Day())
		if _, found := days[date]; found {
			return true
		}
		days[date] = struct{}{}
	}
	return false
}

func RandomBirthdates(size int) []time.Time {
	birthdates := make([]time.Time, 0, size)
	for range size {
		randDate := rand.Intn(totalDays)
		birthdates = append(birthdates, timeStart.AddDate(0, 0, randDate))
	}
	return birthdates
}

func EstimatedProbability(size int) float64 {
	proba := 1.0
	for i := range size {
		proba *= float64(365-i) / 365
	}
	return (1 - proba) * 100
}
