package sumAll

import s "arrays_and_slices/sum"

func SumAll(nums ...[]int) (sums []int) {
	for _, i := range nums {
		sums = append(sums, s.Sum(i))
	}
	return
}
