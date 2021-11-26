package chp_1

import "testing"

var binaryArray = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}

func TestBinarySearch_exist(t *testing.T) {
	requiredIndex := 10

	index := binarySearch(binaryArray, 11)

	if index != requiredIndex {
		t.Fatalf("incorrect result: expected %d, got %d\n", requiredIndex, index)
	}
}

func TestBinarySearch_non_exist(t *testing.T) {
	want := -1

	result := binarySearch(binaryArray, 17)

	if result != want {
		t.Fatalf("incorrect result: expected %d, got %d\n", want, result)
	}
}
