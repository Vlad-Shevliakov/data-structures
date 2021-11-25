package chp_1

import "testing"

var orderedArray = []int{1, 7, 12, 19, 29, 31, 76, 83, 90}

// Assumed: 4 steps => found
func TestLinearSearhOfOrderedArray_exist(t *testing.T) {
	want := 3

	foundIndex := linearOrderedArraySearch(orderedArray, 19)

	if foundIndex != want {
		t.Fatalf("incorrect result: expected %d, got = %d\n", want, foundIndex)
	}
}

// Assumed: 5 steps => on 4 index it doesn't make sense to continue
func TestLinearSearhOfOrderedArray_non_exist(t *testing.T) {
	want := -1

	foundIndex := linearOrderedArraySearch(orderedArray, 20)

	if foundIndex != want {
		t.Fatalf("incorrect result: expected %d, got = %d\n", want, foundIndex)
	}
}
