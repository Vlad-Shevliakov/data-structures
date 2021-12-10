package chp_5

import (
	"reflect"
	"testing"
)

var unsorted = []int{13, 10, 7, 4, 3, 4, 2, 1}

func TestSelectionSort(t *testing.T) {
	want := []int{1, 2, 3, 4, 4, 7, 10, 13}

	got := selectionSort(unsorted)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wrong int selection sort\n")
	}
}
