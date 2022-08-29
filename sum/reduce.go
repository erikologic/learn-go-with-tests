package sum


func Reduce[A, B any](collection []A, accumulator func(B, A) B, initialValue B) B {
	result := initialValue
	for _, number := range collection {
		result = accumulator(result, number)
	}
	return result
}