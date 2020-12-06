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

type group struct {
	answers map[byte]int
	people  int
}

func NewGroup() group {
	return group{
		answers: make(map[byte]int),
		people:  0,
	}
}

func (g *group) add(c byte) {
	g.answers[c]++
}

func (g group) countA() int {
	return len(g.answers)
}

func (g group) countB() int {
	result := 0
	for _, as := range g.answers {
		if as == g.people {
			result++
		}
	}
	return result
}

func (g *group) addPerson() {
	g.people++
}

func parse(data []string) []group {
	var result []group
	g := NewGroup()
	for _, person := range data {
		if person == "" {
			result = append(result, g)
			g = NewGroup()
			continue
		}

		g.addPerson()

		for _, c := range person {
			g.add(byte(c))
		}
	}
	result = append(result, g)
	return result
}

func main() {
	d := readLines("input.txt")
	groups := parse(d)

	countA := 0
	countB := 0
	for _, g := range groups {
		countA += g.countA()
		countB += g.countB()
	}

	fmt.Printf("A: %d\n", countA)
	fmt.Printf("B: %d\n", countB)
}
