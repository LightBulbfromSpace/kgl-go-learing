package revisiting_arrays_and_slices

func Find[A any](collection []A, parameterFunc func(A) bool) (value A, found bool) {
	for _, item := range collection {
		if parameterFunc(item) {
			return item, true
		}
	}
	return
}

func Reduce[A, B any](collection []A, accumulator func(B, A) B, initialValue B) B {
	result := initialValue
	for _, item := range collection {
		result = accumulator(result, item)
	}
	return result
}
