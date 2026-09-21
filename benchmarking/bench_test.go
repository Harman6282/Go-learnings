package main

import "testing"

func BenchmarkAddNumbers(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		AddNumbers(23, 52)

	}

}
