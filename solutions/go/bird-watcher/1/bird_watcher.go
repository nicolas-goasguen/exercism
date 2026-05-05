package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	sum := 0
	for _, v := range birdsPerDay {
		sum += v
	}
	return sum
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	offset := 7 * (week - 1)
	weekBirdsPerDay := birdsPerDay[offset : offset+7]

	sum := 0
	for _, v := range weekBirdsPerDay {
		sum += v
	}

	return sum
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	daysCount := len(birdsPerDay)
	for i := 0; i < daysCount; i += 2 {
		birdsPerDay[i] += 1
	}
	return birdsPerDay
}
