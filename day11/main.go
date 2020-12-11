package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readLines(filename string) []string {
	file, _ := os.Open(filename)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var result []string
	for scanner.Scan() {
		line := strings.TrimSuffix(scanner.Text(), "\n")
		result = append(result, line)
	}

	return result
}

func getSeat(seats [][]byte, x, y int) byte {
	if x < 0 || x >= len(seats[0]) {
		return 0
	}
	if y < 0 || y >= len(seats) {
		return 0
	}

	return seats[y][x]
}

func adjacent(seats [][]byte, x, y int) []byte {
	return []byte{
		getSeat(seats, x-1, y-1),
		getSeat(seats, x+0, y-1),
		getSeat(seats, x+1, y-1),
		getSeat(seats, x+1, y-0),
		getSeat(seats, x+1, y+1),
		getSeat(seats, x-0, y+1),
		getSeat(seats, x-1, y+1),
		getSeat(seats, x-1, y+0)}
}

func ray(seats [][]byte, x, y, dx, dy int) byte {
	posx := x
	posy := y
	for {
		posx += dx
		posy += dy

		seat := getSeat(seats, posx, posy)

		if seat != '.' {
			return seat
		}

		if seat == 0 {
			return '.'
		}
	}
}

func canSee(seats [][]byte, x, y int) []byte {
	return []byte{
		ray(seats, x, y, -1, -1),
		ray(seats, x, y, +0, -1),
		ray(seats, x, y, +1, -1),
		ray(seats, x, y, +1, -0),
		ray(seats, x, y, +1, +1),
		ray(seats, x, y, -0, +1),
		ray(seats, x, y, -1, +1),
		ray(seats, x, y, -1, +0)}
}

func nrOccupied(adj []byte) int {
	result := 0
	for _, a := range adj {
		if a == '#' {
			result++
		}
	}
	return result
}

func equal(a, b [][]byte) bool {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}

func toBytes(seats []string) [][]byte {
	result := make([][]byte, len(seats))

	for i := 0; i < len(seats); i++ {
		result[i] = make([]byte, len(seats[i]))
		for j := 0; j < len(seats[i]); j++ {
			result[i][j] = seats[i][j]
		}
	}

	return result
}

func mutate2(seats [][]byte) [][]byte {
	result := make([][]byte, len(seats))
	copy(result, seats)

	for y := 0; y < len(seats); y++ {
		result[y] = make([]byte, len(seats[y]))
		copy(result[y], seats[y])

		row := seats[y]
		for x := 0; x < len(row); x++ {
			adj := canSee(seats, x, y)
			if seats[y][x] == 'L' && nrOccupied(adj) == 0 {
				result[y][x] = '#'
			}

			if seats[y][x] == '#' && nrOccupied(adj) >= 5 {
				result[y][x] = 'L'
			}
		}
	}

	return result
}

func mutate(seats [][]byte) [][]byte {
	result := make([][]byte, len(seats))
	copy(result, seats)

	for y := 0; y < len(seats); y++ {
		result[y] = make([]byte, len(seats[y]))
		copy(result[y], seats[y])

		row := seats[y]
		for x := 0; x < len(row); x++ {
			adj := adjacent(seats, x, y)
			if seats[y][x] == 'L' && nrOccupied(adj) == 0 {
				result[y][x] = '#'
			}

			if seats[y][x] == '#' && nrOccupied(adj) >= 4 {
				result[y][x] = 'L'
			}
		}
	}

	return result
}

func print(seats [][]byte) {
	for i := 0; i < len(seats); i++ {
		for j := 0; j < len(seats[i]); j++ {
			fmt.Printf("%c", seats[i][j])
		}
		fmt.Printf("\n")
	}
}

func countOccupied(seats [][]byte) int {
	result := 0
	for i := 0; i < len(seats); i++ {
		for j := 0; j < len(seats[i]); j++ {
			if seats[i][j] == '#' {
				result++
			}
		}
	}
	return result
}

func main() {
	d := readLines("input.txt")
	prev := toBytes(d)

	for {
		next := mutate2(prev)
		//print(next)
		//fmt.Printf("-----\n")
		if equal(next, prev) {
			fmt.Printf("equals! occupied: %d\n", countOccupied(next))
			break
		}
		prev = next
	}

	//fmt.Printf("%v\n", d)
}
