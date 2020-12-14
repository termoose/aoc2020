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

type write struct {
	addr  uint64
	value uint64
}

type program struct {
	mask   string
	writes []write
	state  map[uint64]uint64
}

func parse(d []string) program {
	result := program{
		state: make(map[uint64]uint64),
	}

	for i := 0; i < len(d); i++ {
		if d[i][0:4] == "mask" {
			result.mask = d[i][7:]
			continue
		}

		var w write
		fmt.Sscanf(d[i], "mem[%d] = %d", &w.addr, &w.value)

		newVal := result.apply(w.value)
		result.state[w.addr] = newVal

		result.writes = append(result.writes, w)
	}

	return result
}

func (p program) apply(value uint64) uint64 {
	for i := 0; i < len(p.mask); i++ {
		c := p.mask[i]
		bit := len(p.mask) - i - 1

		if c == '1' {
			value = (1 << bit) | value
		} else if c == '0' {
			value = value & (^(1 << bit))
		}
	}

	return value
}

func (p program) apply2(value uint64) string {
	var result []byte
	for i, c := range p.mask {
		bit := len(p.mask) - i - 1
		bitVal := (value >> bit) & 1

		switch c {
		case '1':
			result = append(result, '1')
		case '0':
			result = append(result, strconv.Itoa(int(bitVal))[0])
		case 'X':
			result = append(result, 'X')
		}
	}

	return string(result)
}

func parse2(d []string) program {
	result := program{
		state: make(map[uint64]uint64),
	}

	for i := 0; i < len(d); i++ {
		if d[i][0:4] == "mask" {
			result.mask = d[i][7:]
			continue
		}

		var w write
		fmt.Sscanf(d[i], "mem[%d] = %d", &w.addr, &w.value)
		newAddr := result.apply2(w.addr)
		all := getAll(newAddr)
		addrs := toAddrs(all)

		for _, a := range addrs {
			result.state[a] = w.value
		}
	}

	return result
}

func getAll(addr string) []string {
	addrMod := []byte(addr)

	for i, c := range addr {
		if c == 'X' {
			addrMod[i] = '1'
			first := getAll(string(addrMod))

			addrMod[i] = '0'
			second := getAll(string(addrMod))

			return append(first, second...)
		}
	}

	return []string{string(addrMod)}
}

func toAddr(addr string) uint64 {
	i, _ := strconv.ParseInt(addr, 2, 64)
	return uint64(i)
}

func toAddrs(addrs []string) []uint64 {
	var result []uint64
	for _, a := range addrs {
		result = append(result, toAddr(a))
	}
	return result
}

func (p program) sum() uint64 {
	var result uint64
	for _, s := range p.state {
		result += s
	}
	return result
}

func main() {
	d := readLines("input.txt")
	p := parse(d)
	fmt.Printf("Part A: %v\n", p.sum())

	p2 := parse2(d)
	fmt.Printf("Part B: %d\n", p2.sum())
}
