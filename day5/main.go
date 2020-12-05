package main

import (
	"bufio"
	"fmt"
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

func rowToBin(s string) string {
	var result []byte
	for _, c := range s {
		if c == 'B' {
			result = append(result, '1')
		} else if c == 'F' {
			result = append(result, '0')
		}
	}
	return string(result)
}

func columnToBin(s string) string {
	var result []byte
	for _, c := range s {
		if c == 'R' {
			result = append(result, '1')
		} else if c == 'L' {
			result = append(result, '0')
		}
	}
	return string(result)
}

func toDec(s string) int {
	i, _ := strconv.ParseInt(s, 2, 64)
	return int(i)
}

type pos struct {
	col int
	row int
}

func parse(data []string) {
	highest := 0
	list := make(map[int]pos)
	for i := 1; i < 126; i++ {
		for j := 0; j < 8; j++ {
			list[i * 8 + j] = pos{col: i, row: j}
		}
	}

	for _, pass := range data {
		row := rowToBin(pass)
		rowNr := toDec(row)
		col := columnToBin(pass)
		colNr := toDec(col)
		id := rowNr * 8 + colNr

		delete(list, id)

		if id > highest {
			highest = id
		}

		fmt.Printf("%d %d id: %d\n", rowNr, colNr, rowNr * 8 + colNr)
	}

	fmt.Printf("highest: %d\n", highest)
	for key, val := range list {
		_, left := list[key-1]
		_, right := list[key+1]
		if !left && !right {
			fmt.Printf("key: %v val: %v\n", key, val)
		}
	}
}

func main() {
	data := readLines("input.txt")
	parse(data)
	//fmt.Printf("%s\n", data)
}
