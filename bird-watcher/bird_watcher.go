package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	var TotalBirdCount int = 0
	for _, v := range birdsPerDay {
		TotalBirdCount += v
	}
	return TotalBirdCount
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	dayIndex := (week - 1) * 7
	return TotalBirdCount(birdsPerDay[dayIndex : dayIndex+7])
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay); i = i + 2 {
		birdsPerDay[i] = birdsPerDay[i] + 1
	}
	return birdsPerDay
}
