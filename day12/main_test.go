package main

import (
	"fmt"
	"testing"
)

func BenchmarkDay12(b *testing.B) {
	var instructions []instruction
	resultA := 0
	resultB := 0

	b.Run("Parse", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			data := readLines("input.txt")
			instructions = parse(data)
		}
	})

	b.Run("PartA", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			s := ship{
				x:         0,
				y:         0,
				direction: 'E',
				wx:        10,
				wy:        1,
			}
			for _, i := range instructions {
				s.step(i)
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
			for _, i := range instructions {
				s.stepw(i)
			}
			resultB = abs(s.x) + abs(s.y)
		}
	})

	fmt.Printf("Result A: %d\n", resultA)
	fmt.Printf("Result B: %d\n", resultB)
}