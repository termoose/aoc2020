package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

const (
	Empty = iota
	Tree
)

type world struct {
	forest [][]int
}

func parse(input []string) world {
	var result world
	result.forest = make([][]int, len(input))

	for y, line := range input {
		result.forest[y] = make([]int, len(line))
		for x, char := range line {
			if char == '.' {
				result.forest[y][x] = Empty
			} else if char == '#' {
				result.forest[y][x] = Tree
			}
		}
	}

	return result
}

func (w world) walk(x, y int) (int, bool) {
	width := len(w.forest[0])
	height := len(w.forest)

	return w.forest[y % height][x % width], y >= height
}

func (w world) treesInSlope(dx, dy int) int {
	x := 0
	y := 0
	trees := 0
	for {
		x += dx
		y += dy
		thing, wrapped := w.walk(x, y)

		if wrapped == true {
			break
		}

		if thing == Tree {
			trees++
		}
	}

	return trees
}

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

func main() {
	parsing := time.Now()
	data := readLines("input.txt")
	f := parse(data)
	fmt.Printf("Parsing: %s\n", time.Since(parsing))

	part1 := time.Now()
	trees := f.treesInSlope(3, 1)
	fmt.Printf("A) Trees %d, Time: %s\n", trees, time.Since(part1))

	part2 := time.Now()
	trees1 := f.treesInSlope(1, 1)
	trees2 := f.treesInSlope(3, 1)
	trees3 := f.treesInSlope(5, 1)
	trees4 := f.treesInSlope(7, 1)
	trees5 := f.treesInSlope(1, 2)
	fmt.Printf("B) Prod: %d, Time: %s\n",
		trees1 * trees2 * trees3 * trees4 * trees5, time.Since(part2))
}
