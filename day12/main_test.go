package main

import (
	"fmt"
	"testing"
)

func BenchmarkDay12(b *testing.B) {
	data := readLines("input.txt")
	resultA := 0
	resultB := 0
	b.Run("PartA", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			s := ship{
				x:         0,
				y:         0,
				direction: 'E',
				wx:        10,
				wy:        1,
			}
			for _, l := range data {
				s.step(l)
			}
			resultA = abs(s.x) + abs(s.y)
		}
	})

	b.Run("PartB", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			s := ship{
				x:         0,
				y:         0,
				direction: 'E',
				wx:        10,
				wy:        1,
			}
			for _, l := range data {
				s.stepw(l)
			}
			resultB = abs(s.x) + abs(s.y)
		}
	})

	fmt.Printf("Result A: %d\n", resultA)
	fmt.Printf("Result B: %d\n", resultB)
}