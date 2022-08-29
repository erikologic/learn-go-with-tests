package sum


func Reduce[T any](collection []T, accumulator func(T, T) T, initialValue T) T {
	result := initialValue
	for _, number := range collection {
		result = accumulator(result, number)
	}
	return result
}