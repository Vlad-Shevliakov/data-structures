package chp_5

/*
	Selection sort takes about half the number of steps Bubble Sort does
	N elements		Max steps Bubble sort		Max steps Selection Sort
		5						20							14
		10						90							54
		20						380							199
		40						1560						819

	But Selection sort described in Big O as O(N2), just like Bubble sort
*/

func selectionSort(src []int) []int {
	dst := make([]int, len(src))
	copy(dst, src)

	// ir doesn't need to run for the last value itself
	// since the array will be fully sorted bu that point
	for i := 0; i < len(dst)-1; i++ {
		big := i

		for j := i + 1; j < len(dst); j++ {
			if dst[j] < dst[big] {
				big = j
			}
		}

		if big != i {
			temp := dst[i]
			dst[i] = dst[big]
			dst[big] = temp
		}
	}

	return dst
}
