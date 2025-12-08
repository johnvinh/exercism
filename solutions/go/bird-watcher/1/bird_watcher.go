package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	sum := 0
    for i := 0; i < len(birdsPerDay); i++ {
        sum += birdsPerDay[i]
    }
    return sum
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	startingIndex := 7 * (week - 1)
    endingIndex := startingIndex + 6
    sum := 0
    for i := startingIndex; i < (endingIndex+1); i++ {
        sum += birdsPerDay[i]
    }
    return sum
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
    for i := 0; i < len(birdsPerDay); i+= 2 {
        if i == 0 {
            birdsPerDay[i] += 1
        } else {
            birdsPerDay[i] += 1
        }
    }
    return birdsPerDay
}
