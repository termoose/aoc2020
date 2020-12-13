package main

import (
	"bufio"
	"fmt"
	"math"
	"math/big"
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

type travel struct {
	earliest int
	buses    []int
}

func parse(d []string) travel {
	earliest, _ := strconv.Atoi(d[0])
	var buses []int
	bs := strings.Split(d[1], ",")
	for _, b := range bs {
		if b == "x" {
			buses = append(buses, 0)
			continue
		}

		bus, _ := strconv.Atoi(b)
		buses = append(buses, bus)
	}

	return travel{
		earliest: earliest,
		buses: buses,
	}
}

func waitFast(earliest, busId int) int {
	return busId - (earliest % busId)
}

func earliest(t travel) int {
	smallest := math.MaxInt32
	busId := 0

	for _, b := range t.buses {
		if b == 0 {
			continue
		}
		w := waitFast(t.earliest, b)
		if w < smallest {
			smallest = w
			busId = b
		}
	}

	return smallest * busId
}

func inv(elem, ring int) int {
	r := big.NewInt(int64(ring))
	e := big.NewInt(int64(elem))

	return int(r.ModInverse(e, r).Int64())
}

func chinese(t travel) int {
	N := 1
	var ys []int
	var zs []int
	for _, n := range t.buses {
		if n == 0 {
			continue
		}
		N *= n
	}
	for _, n := range t.buses {
		if n == 0 {
			continue
		}
		y := N/n
		z := inv(y, n)
		ys = append(ys, y)
		zs = append(zs, z)
	}

	result := 0
	count := 0
	for i, n := range t.buses {
		if n == 0 {
			continue
		}
		a := n-i-1
		y := ys[count]
		z := zs[count]
		//fmt.Printf("%d * %d * %d\n", a, y, z)
		result += a * y * z
		count++
	}
	return result % N
}

func main() {
	d := readLines("input.txt")
	p := parse(d)

	res := earliest(p)
	fmt.Printf("resultA: %d\n", res)

	c := chinese(p)
	fmt.Printf("resultB: %d\n", c+1)
}
