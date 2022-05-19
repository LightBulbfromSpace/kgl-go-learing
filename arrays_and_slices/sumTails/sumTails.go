package sumTails

func SumTails(arr ...[]int) ([]int sum_of_tails)  {
	sum_of_tail := 0
	for i := 1; i < len(arr); i++ {
		sum_of_tail += arr[i]
	}
	return
}
