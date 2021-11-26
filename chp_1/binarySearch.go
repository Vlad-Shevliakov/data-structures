package chp_1

/*
	array of 100,000 just in 16 steps
	dividing 100,000 by 2 until we get down to ~1
	100,000 -> 50,000 -> 25,000 -> 12,500 ...
*/

func binarySearch(array []int, value int) int {

	// range of indexes
	lowerBound := 0
	upperBound := len(array) - 1

	// If we've narrowed the bounds until they
	// have reached each other, that means
	// that the value is not contained within this array
	for lowerBound <= upperBound {
		midPoint := (upperBound + lowerBound) / 2

		// Inspect the value at the middle
		midValue := array[midPoint]

		if midValue == value {
			return midPoint
		} else if value < midValue {
			upperBound = midPoint - 1
		} else if value > midValue {
			lowerBound = midPoint + 1
		}
	}

	return -1
}
