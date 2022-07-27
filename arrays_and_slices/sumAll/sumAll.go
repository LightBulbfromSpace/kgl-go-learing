package arrays_and_slices

import s "github.com/LightBulbfromSpace/kld-go-learning/arrays_and_slices/sum"

func SumAll(nums ...[]int) (sums []int) {
	for _, i := range nums {
		sums = append(sums, s.Sum(i))
	}
	return
}
