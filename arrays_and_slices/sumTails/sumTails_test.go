package sumTails

import (
	"reflect"
	"testing"
)

func TestSumTails(t *testing.T) {

	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v but want %v", got, want)
		}
	}

	t.Run("non-empty slices", func(t *testing.T) {

		got := SumTails([]int{2, 5}, []int{4, 1})
		want := []int{5, 1}
		checkSums(t, got, want)
	})
	t.Run("contains empty slices", func(t *testing.T) {

		got := SumTails([]int{}, []int{4, 1})
		want := []int{0, 1}
		checkSums(t, got, want)
	})

}
