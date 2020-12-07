package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
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

type bag struct {
	color    string
	count    int
	contains map[string]int
}

func parse(d []string) map[string]bag {
	result := make(map[string]bag)

	for _, line := range d {
		parts := strings.Split(strings.Trim(line, "."), " contain ")
		color := parts[0][0 : len(parts[0])-5] // remove " bags"
		inside := strings.Split(parts[1], ", ")

		//fmt.Printf("%s CONTAINS %v\n", color, inside)
		newBag := bag{
			color:    color,
			contains: make(map[string]int),
		}

		for _, bagInside := range inside {
			re, _ := regexp.Compile(`(\d+) (.+) (bag|bags)`)
			nrColor := re.FindStringSubmatch(bagInside)
			if len(nrColor) > 1 {
				count, _ := strconv.Atoi(nrColor[1])
				newBag.contains[nrColor[2]] = count
				//fmt.Printf("Count: %d Type: %s\n", count, nrColor[2])
			}
		}

		result[color] = newBag
	}

	return result
}

var result = make(map[string]bool)

func find(luggage map[string]bag, color string) string {
	for top, bag := range luggage {
		_, contains := bag.contains[color]

		if contains {
			result[top] = true
			find(luggage, top)
		}
	}

	return ""
}

func counter(luggage map[string]bag, target string) int {
	total := 0

	for color, count := range luggage[target].contains {
		total += count * (1 + counter(luggage, color))
	}

	return total
}

func main() {
	d := readLines("input.txt")
	res := parse(d)
	_ = find(res, "shiny gold")
	fmt.Printf("Result A: %d\n", len(result))

	sum := counter(res, "shiny gold")
	fmt.Printf("Result B: %d\n", sum)
}
