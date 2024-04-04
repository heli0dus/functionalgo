package examples

import "testing"

func BenchmarkStreams(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SolveWithStreams(TextLines, W1, W2)
	}
}

func BenchmarkClassic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SolveClassic(TextLines, W1, W2)
	}
}
