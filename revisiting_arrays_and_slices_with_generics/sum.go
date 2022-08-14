package revisiting_arrays_and_slices

// Sum calculates the total from a slice of numbers.
func Sum(numbers []int) int {
	sum := func(n1, n2 int) int {
		return n1 + n2
	}
	return Reduce[int](numbers, sum, 0)
}

// SumAll calculates the sums in all given slices.
func SumAll(nums ...[]int) []int {
	sumAll := func(sums, numbers []int) []int {
		return append(sums, Sum(numbers))
	}
	return Reduce[[]int](nums, sumAll, []int{})
}

// SumAllTails calculates the sums of all but the first number given a collection of slices.
func SumAllTails(numsToSum ...[]int) []int {
	sumTail := func(sums, numbers []int) []int {
		if len(numbers) == 0 {
			return append(sums, 0)
		} else {
			tail := numbers[1:]
			return append(sums, Sum(tail))
		}
	}
	return Reduce[[]int](numsToSum, sumTail, []int{})
}
