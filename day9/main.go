package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
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

func findSum(numbers []int, target int) bool {
	for _, a := range numbers {
		for _, b := range numbers {
			if a != b && a+b == target {
				return true
			}
		}
	}

	return false
}

func find(numbers []int, preamble int) int {
	for i := preamble; i < len(numbers); i++ {
		from := i - preamble
		to := i - 1
		sums := findSum(numbers[from:to+1], numbers[to+1])
		if !sums {
			return numbers[to+1]
		}
	}

	return 0
}

func checkRange(numbers []int, target int) (int, bool) {
	sum := 0
	smallest := math.MaxInt32
	largest := 0

	for _, n := range numbers {
		if n > largest {
			largest = n
		}
		if n < smallest {
			smallest = n
		}
		sum += n
	}

	return smallest + largest, target == sum
}

func findRange(numbers []int, target int) int {
	for r := 2; r < len(numbers); r++ {
		for i := 0; i < len(numbers)-r; i++ {
			from := i
			to := i + r
			res, check := checkRange(numbers[from:to], target)
			if check {
				return res
			}
		}
	}

	return 0
}

func main() {
	d := readLines("input.txt")
	res := find(d, 25)
	fmt.Printf("Result 1: %d\n", res)

	r := findRange(d, res)
	fmt.Printf("Result 2: %d\n", r)
}
