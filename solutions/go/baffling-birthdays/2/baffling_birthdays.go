package bafflingbirthdays

import (
	"math/rand"
	"time"
)

const year = 2026
const totalDays = 365

func SharedBirthday(dates []time.Time) bool {
	var days = make(map[[2]int]struct{})
	for _, d := range dates {
		key := [2]int{int(d.Month()), d.Day()}
		if _, found := days[key]; found {
			return true
		}
		days[key] = struct{}{}
	}
	return false
}

func RandomBirthdates(size int) []time.Time {
	birthdates := make([]time.Time, size)
	for i := range birthdates {
		birthdates[i] = time.Date(year, 1, 1+rand.Intn(totalDays), 0, 0, 0, 0, time.UTC)
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
