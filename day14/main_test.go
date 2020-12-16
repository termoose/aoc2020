package main

import (
	"fmt"
	"testing"
)

func BenchmarkDay14(b *testing.B) {
	d := readLines("input.txt")
	var resultA uint64
	var resultB uint64

	b.Run("PartA", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			p := parse(d)
			resultA = p.sum()
		}
	})

	b.Run("PartB", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			p := parse2(d)
			resultB = p.sum()
		}
	})

	fmt.Printf("ResultA: %d\n", resultA)
	fmt.Printf("ResultB: %d\n", resultB)
}
