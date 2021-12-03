package chp4

func bubbleSort(arr []int) []int {

	unsortedUntilIndex := len(arr) - 1
	bound := false

	src := make([]int, len(arr))

	copy(src, arr)

	for !bound {
		bound = true

		for i := 0; i < unsortedUntilIndex; i++ {
			if src[i] > src[i+1] {
				src[i], src[i+1] = src[i+1], src[i]
				bound = false
			}
		}

		unsortedUntilIndex -= 1

	}

	return src
}
