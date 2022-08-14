package revisiting_arrays_and_slices

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}

	got := Sum(numbers)
	want := 15

	if got != want {
		t.Errorf("got '%d' but want '%d', %v", got, want, numbers)
	}
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{0, 7}, []int{1, 2})
	want := []int{7, 3}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v but wanted %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v but want %v", got, want)
		}
	}

	t.Run("non-empty slices", func(t *testing.T) {

		got := SumAllTails([]int{2, 5}, []int{4, 1})
		want := []int{5, 1}
		checkSums(t, got, want)
	})
	t.Run("contains empty slices", func(t *testing.T) {

		got := SumAllTails([]int{}, []int{4, 1})
		want := []int{0, 1}
		checkSums(t, got, want)
	})

}

func TestReduce(t *testing.T) {
	t.Run("multiplying", func(t *testing.T) {
		multiply := func(x, y int) int {
			return x * y
		}
		AssertEqual(t, Reduce([]int{1, 2, 3}, multiply, 1), 6)
	})
	t.Run("concatenation", func(t *testing.T) {
		concatenate := func(str1, str2 string) string {
			return fmt.Sprintf("%s%s", str1, str2)
		}
		AssertEqual(t, Reduce([]string{"1", "2", "3"}, concatenate, ""), "123")
	})
}
