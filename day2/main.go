package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type policy struct {
	letter byte
	max    int
	min    int
}

type password struct {
	pass  string
	rules policy
}

// Just Go Things
func bool2int(b bool) int {
	if b {
		return 1
	}
	return 0
}

func int2bool(i int) bool {
	if i == 0 {
		return false
	}
	return true
}

func (p password) validA() bool {
	count := strings.Count(p.pass, string(p.rules.letter))
	return count <= p.rules.max && count >= p.rules.min
}

func (p password) validB() bool {
	first := bool2int(p.pass[p.rules.min-1] == p.rules.letter)
	second := bool2int(p.pass[p.rules.max-1] == p.rules.letter)
	return int2bool(first ^ second)
}

func parse(line string) (password, error) {
	var rules policy
	var pass string
	_, err := fmt.Sscanf(line, "%d-%d %c: %s", &rules.min, &rules.max, &rules.letter, &pass)

	if err != nil {
		return password{}, fmt.Errorf("error parsing: %w", err)
	}

	return password{pass: pass, rules: rules}, nil
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
	lines := readLines("input.txt")
	var passwords []password
	for _, line := range lines {
		pass, _ := parse(line)
		passwords = append(passwords, pass)
	}
	fmt.Printf("Parsing: %s\n", time.Since(parsing))

	start := time.Now()
	validA := 0
	for _, pass := range passwords {
		if pass.validA() {
			validA++
		}
	}
	fmt.Printf("Valid A: %d in: %s\n", validA, time.Since(start))

	start = time.Now()
	validB := 0
	for _, pass := range passwords {
		if pass.validB() {
			validB++
		}
	}
	fmt.Printf("Valid B: %d in: %s\n", validB, time.Since(start))
}
