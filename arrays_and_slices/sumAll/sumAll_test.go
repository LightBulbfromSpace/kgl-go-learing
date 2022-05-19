package sumAll

import (
	"reflect"
	"testing"
)

func TestSumAll(t *testing.T) {
	got := SumAll([]int{0, 7}, []int{1, 2})
	want := []int{7, 3}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v but wanted %v", got, want)
	}
}
