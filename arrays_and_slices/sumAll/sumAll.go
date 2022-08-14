package arrays_and_slices

import s "github.com/LightBulbfromSpace/kld-go-learning/arrays_and_slices/sum"

// SumAll calculates the sums in all given slices.
func SumAll(nums ...[]int) (sums []int) {
	for _, numSingleSlice := range nums {
		sums = append(sums, s.Sum(numSingleSlice))
	}
	return
}
