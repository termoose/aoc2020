package main

import (
	"testing"
)

func BenchmarkDay9(b *testing.B) {
	data := readLines("input.txt")
	resultA := 0

	b.Run("Part1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			resultA = find(data, 25)
		}
	})

	b.Run("Part2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = findRange(data, resultA)
		}
	})
}