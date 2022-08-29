package sum

import "testing"

func makeRange(min, max int) []int {
    a := make([]int, max-min+1)
    for i := range a {
        a[i] = min + i
    }
    return a
}

var numbers = makeRange(0, 99999)

func BenchmarkReduce(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum(numbers)
	}
}

func BenchmarkForLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumForLoop(numbers)
	}
}

func SumForLoop(numbers []int) int {
	res := 0
	for _, number := range numbers {
		res += number
	}
	return res
}