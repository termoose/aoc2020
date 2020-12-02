package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

func read(file string) string {
	dat, _ := ioutil.ReadFile(file)
	return string(dat)
}

func readInts(filename string) []int {
	file, _ := os.Open(filename)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var result []int
	for scanner.Scan() {
		line := strings.TrimSuffix(scanner.Text(), "\n")
		integer, _ := strconv.Atoi(line)
		result = append(result, integer)
	}

	return result
}

func findInt(list []int, target int) (int, error) {
	found := sort.SearchInts(list, target)

	if found < len(list) && list[found] == target {
		return list[found], nil
	}

	return 0, fmt.Errorf("not found")
}

func findInts(list []int, sum int) (int, int, error) {
	for _, val := range list {
		target := sum - val
		found := sort.SearchInts(list, target)

		if found < len(list) && list[found] == target {
			return list[found], val, nil
		}
	}

	return 0, 0, fmt.Errorf("no solution")
}

func main() {
	data := readInts("input.txt")
	sort.Ints(data)
	a, b, _ := findInts(data, 2020)
	fmt.Printf("hello lol: %d+%d=%d, %d*%d = %d\n", a, b, a+b, a, b, a*b)

	for c := 0; c < 2020; c++ {
		_, err := findInt(data, c)
		if err != nil {
			continue
		}

		a, b, err := findInts(data, 2020 - c)
		if err == nil && (a+b+c == 2020) {
			fmt.Printf("%d+%d+%d = %d\n", a, b, c, a+b+c)
			fmt.Printf("Prod: %d\n", a*b*c)
		}
	}
}
