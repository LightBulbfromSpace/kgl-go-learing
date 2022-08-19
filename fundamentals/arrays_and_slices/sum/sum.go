package arrays_and_slices

// Sum calculates the total from a slice of numbers.
func Sum(numbers []int) int {
	var sum int
	for _, num := range numbers { // _ is "blank identifier" used to ignore index provided by "range"
		sum += num
	}
	return sum
}
