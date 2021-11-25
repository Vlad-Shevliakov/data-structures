package chp_1

func linearOrderedArraySearch(array []int, value int) int {
	for ind, v := range array {
		if v == value {
			return ind
		}
		if v > value {
			break
		}

	}

	return -1
}
