package main

import (
	"bufio"
	"fmt"
	"golang.org/x/tools/container/intsets"
	"os"
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

func evaluate(expr string) int {
	ops := strings.Split(expr, " ")
	acc, _ := strconv.Atoi(ops[0])

	for i := 2; i < len(ops); i += 2 {
		nr, _ := strconv.Atoi(ops[i])

		op := ops[i-1]

		switch op {
		case "+":
			acc += nr
		case "*":
			acc *= nr
		}

	}

	return acc
}

func evaluate2(expr string) int {
	//fmt.Printf("expr: %s\n", expr)
	ops := strings.Split(expr, " * ")

	var prod []int
	for i := 0; i < len(ops); i++ {
		addends := strings.Split(ops[i], " + ")
		//fmt.Printf("%s\n", ops[i])

		sum := 0
		for _, a := range addends {
			num, _ := strconv.Atoi(a)
			sum += num
		}
		//fmt.Printf("%v = %d\n", addends, sum)
		prod = append(prod, sum)
	}

	result := 1
	for _, n := range prod {
		result *= n
	}

	return result
}

func inner(line string) int {
	last := intsets.MaxInt

	for i := 0; i < len(line); i++ {
		idx := len(line) - i - 1

		if line[idx] == ')' {
			last = idx
		}

		if line[idx] == '(' {
			left := line[:idx]
			right := line[last+1:]
			expr := line[idx+1:last]
			val := evaluate2(expr)

			return inner(left + strconv.Itoa(val) + right)
		}
	}

	return evaluate2(line)
}

func main() {
	data := readLines("input.txt")

	//test := evaluate2("1 + 2 * 3 + 4 * 5 + 6")
	//fmt.Printf("test: %d\n", test)
	//test := inner(data[1])
	//fmt.Printf("test: %d\n", test)

	sum := 0
	for _, l := range data {
		sum += inner(l)

	}

	// wrong: 448594014470847
	fmt.Printf("sum: %d\n", sum)
}
