package main

import (
	"fmt"
	"testing"
)

func BenchmarkDay13(b *testing.B) {
	var p travel

	b.Run("Parse", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			d := readLines("input.txt")
			p = parse(d)
		}
	})

	resultA := 0
	resultB := 0

	b.Run("PartA", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			resultA = earliest(p)
		}
	})

	b.Run("PartB", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			resultB = chinese(p)
		}
	})

	fmt.Printf("Result A: %d\n", resultA)
	fmt.Printf("Result B: %d\n", resultB + 1)
}
