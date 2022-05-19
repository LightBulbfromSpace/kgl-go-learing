package sumTails

import (
	"reflect"
	"testing"
)

func TestSumTails(t *testing.T) {
	got := SumTails([]int{2, 5}, []int{4, 1})
	want := []int{5, 1}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v but want %v", got, want)
	}
}
