package chp4

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	want := []int{10, 20, 30, 40}

	x := []int{30, 10, 40, 20}

	got := bubbleSort(x)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wrong int bubble sort\n")
	}
}
