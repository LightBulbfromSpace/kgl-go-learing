package sumTails

import (
	obj "arrays_and_slices/sum"
)

func SumTails(numsToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numsToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, obj.Sum(tail))
		}
	}
	return sums
}
