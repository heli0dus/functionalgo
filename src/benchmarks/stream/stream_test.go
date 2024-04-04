package stream

import (
	"github.com/heli0dus/functionalgo/src/stream"
	"strconv"
	"testing"
)

// Re-implementing isEven, square, and add for completeness
func isEven(i int) bool {
	return i%2 == 0
}

func square(i int) int {
	return i * i
}

func add(a int, b int) int {
	return a + b
}

var coolNumbers = []string{"42", "69", "13", "666", "1337", "1", "0", "-5", "12"}

// Benchmark for classic approach
func BenchmarkClassic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var ints []int
		for _, str := range coolNumbers {
			num, _ := strconv.Atoi(str)
			if isEven(num) {
				ints = append(ints, square(num))
			}
		}

		sum := 0
		for _, num := range ints {
			sum = add(sum, num)
		}
	}
}

func BenchmarkClassicFilterOnly(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var result []int
		for _, str := range coolNumbers {
			i, err := strconv.Atoi(str)
			if err == nil && isEven(i) {
				result = append(result, i)
			}
		}
	}
}

func BenchmarkClassicMapOnly(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var result []int
		for _, str := range coolNumbers {
			i, err := strconv.Atoi(str)
			if err == nil {
				result = append(result, square(i))
			}
		}
	}
}

func BenchmarkClassicReduceOnly(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sum := 0
		for _, str := range coolNumbers {
			num, _ := strconv.Atoi(str)
			sum = add(sum, num)
		}
	}
}

func BenchmarkStream(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := stream.AsStream(coolNumbers)
		s = s.
			Fmap(strconv.Atoi).
			Filter(isEven).
			Fmap(square).
			Reduce(add, 0)
	}
}

func BenchmarkStreamFilterOnly(b *testing.B) {
	s := stream.AsStream(coolNumbers)
	s = s.Fmap(strconv.Atoi)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s = s.Filter(isEven)
	}
}

func BenchmarkStreamMapOnly(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := stream.AsStream(coolNumbers)
		s = s.Fmap(strconv.Atoi)
	}
}
func BenchmarkStreamReduceOnly(b *testing.B) {
	s := stream.AsStream(coolNumbers)
	s = s.Fmap(strconv.Atoi)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s = s.Reduce(add, 0)
	}
}
