package main

func Sum(numbers []int) int {
	var sum int
	for _, num := range numbers { // _ is "blank identifier" used to ignore index provided by "range"
		sum += num
	}
	return sum
}
