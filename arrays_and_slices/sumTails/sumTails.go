package arrays_and_slices

import s "github.com/LightBulbfromSpace/kld-go-learning/arrays_and_slices/sum"

// SumTails calculates the sums of all but the first number given a collection of slices.
func SumTails(numsToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numsToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, s.Sum(tail))
		}
	}
	return sums
}
