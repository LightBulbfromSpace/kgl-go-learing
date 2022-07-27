package arrays_and_slices

import "testing"

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}

	got := Sum(numbers)
	want := 15

	if got != want {
		t.Errorf("got '%d' but want '%d', %v", got, want, numbers)
	}
}
