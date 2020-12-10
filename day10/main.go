package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readLines(filename string) []int {
	file, _ := os.Open(filename)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var result []int
	for scanner.Scan() {
		line := strings.TrimSuffix(scanner.Text(), "\n")
		n, _ := strconv.Atoi(line)
		result = append(result, n)
	}

	return result
}

func find(adaptors []int, start, r int) int {
	target := start
	diffs := []int{0, 0, 0, 1}
	for _, n := range adaptors {
		if n-target <= r {
			diffs[n-target]++
			target = n
		}
	}

	return diffs[1] * diffs[3]
}

func count(adaptors []int) int {
	c := make(map[int]int, len(adaptors)+3)
	c[0] = 1

	for _, a := range adaptors {
		c[a] += c[a-1] + c[a-2] + c[a-3]
	}

	return c[adaptors[len(adaptors)-1]]
}

func fastCount(a map[int]int, latest int) int {
	a[0] = 1

	for key, _ := range a {
		a[key] += a[key-1] + a[key-2] + a[key-3]
	}

	return a[latest]
}

func main() {
	d := readLines("input.txt")
	dataMap := make(map[int]int)
	sort.Ints(d)
	for _, d := range d {
		dataMap[d] = 0
	}

	resultA := find(d, 0, 3)
	resultB := fastCount(dataMap, d[len(d)-1])
	fmt.Printf("result A: %d\n", resultA)
	fmt.Printf("result B: %d\n", resultB)
}
