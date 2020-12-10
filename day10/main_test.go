package main

import (
	"fmt"
	"sort"
	"testing"
)

func BenchmarkDay10(b *testing.B) {
	var data []int
	dataMap := make(map[int]int)
	b.Run("Init and sort", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			data = readLines("input.txt")
			sort.Ints(data)
			for _, d := range data {
				dataMap[d] = 0
			}
		}
	})

	resultA := 0
	resultB := 0
	b.Run("Part1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			resultA = find(data, 0, 3)
		}
	})

	b.Run("Part2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			resultB = count(data)
		}
	})

	b.Run("Part2Fast", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			resultB = fastCount(dataMap, data[len(data)-1])
		}
	})

	fmt.Printf("PartA: %d\n", resultA)
	fmt.Printf("PartB: %d\n", resultB)
}
