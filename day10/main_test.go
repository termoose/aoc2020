package main

import (
	"fmt"
	"sort"
	"testing"
)

func BenchmarkDay9(b *testing.B) {
	data := readLines("input.txt")
	sort.Ints(data)
	resultA := 0
	resultB := 0
	b.Run("Part1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, diffa, diffb := find(data, 0, 3)
			resultA = diffa * diffb
		}
	})

	b.Run("Part2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			resultB = count(data, 3)
		}
	})

	fmt.Printf("PartA: %d\n", resultA)
	fmt.Printf("PartB: %d\n", resultB)
}
