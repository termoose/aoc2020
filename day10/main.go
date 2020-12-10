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

func find(adaptors []int, start, r int) ([]int, int, int) {
	var result []int
	target := start
	diff1 := 0
	diff3 := 1
	for _, n := range adaptors {
		//fmt.Printf("From %d range %d check %d\n", target, r, n)
		if n-target <= r {
			//fmt.Printf("n (%d) - target (%d) = %d\n", n, target, n-target)
			result = append(result, n)

			if n-target == 1 {
				diff1++
			}

			if n-target == 3 {
				diff3++
			}

			target = n
		}
	}

	result = append(result, target+3)
	return result, diff1, diff3
}

func count(adaptors []int, r int) int {
	c := make(map[int]int)
	c[0] = 1 // starting adaptor is reached by 1 other

	//// do we have to?
	//c[-1] = 0
	//c[-2] = 0

	for _, a := range adaptors {
		// adaptor n can be reached by n-1, n-2 and n-3 (r = 3)
		//for i := 1; i <= r; i++ {
			//fmt.Printf("adaptor %d reaches %d\n", a, c[a-i])
			//c[a] += c[a-i]
		//}
		c[a] += c[a-1] + c[a-2] + c[a-3]
	}

	// last adaptor N-1 can be reached in this many ways
	//fmt.Printf("res: %d\n", c[adaptors[len(adaptors)-1]])
	return c[adaptors[len(adaptors)-1]]
}

func main() {
	d := readLines("small.txt")
	sort.Ints(d)
	found, diffa, diffb := find(d, 0, 3)
	result := count(d, 3)
	fmt.Printf("result: %d\n", result)
	fmt.Printf("data: %v found: %v! %d %d prod: %d\n", d, found, diffa, diffb, diffa*diffb)
}
