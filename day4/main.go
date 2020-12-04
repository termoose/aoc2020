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

type passport struct {
	fields map[string]string
}

func NewPassport() passport {
	return passport{
		fields: make(map[string]string),
	}
}

func (p *passport) add(key, val string) {
	p.fields[key] = val
}

func byrValid(byr string) bool {
	i, _ := strconv.Atoi(byr)
	return i >= 1920 && i <= 2002
}

func iyrValid(iyr string) bool {
	i, _ := strconv.Atoi(iyr)
	return i >= 2010 && i <= 2020
}

func eyrValid(eyr string) bool {
	i, _ := strconv.Atoi(eyr)
	return i >= 2020 && i <= 2030
}

func hgtValid(hgt string) bool {
	if len(hgt) < 4 {
		return false
	}
	cm := strings.Contains(hgt, "cm")
	height := hgt[:len(hgt)-2]
	i, _ := strconv.Atoi(height)

	if cm {
		return i >= 150 && i <= 193
	}

	return i >= 59 && i <= 76
}

func hclValid(hcl string) bool {
	if len(hcl) > 0 && hcl[0] == '#' {
		match, _ := regexp.MatchString("[a-f0-9_]{6}$", hcl)
		return match
	}
	return false
}

func eclValid(ecl string) bool {
	colors := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
	for _, c := range colors {
		found := strings.Contains(ecl, c)
		if found {
			return true
		}
	}

	return false
}

func pidValid(pid string) bool {
	if len(pid) != 9 {
		return false
	}

	match, _ := regexp.MatchString("[0-9_]{9}$", pid)
	return match
}

func (p passport) valid2() bool {
	byr := byrValid(p.fields["byr"])
	iyr := iyrValid(p.fields["iyr"])
	eyr := eyrValid(p.fields["eyr"])
	hgt := hgtValid(p.fields["hgt"])
	hcl := hclValid(p.fields["hcl"])
	ecl := eclValid(p.fields["ecl"])
	pid := pidValid(p.fields["pid"])
	return byr && iyr && eyr && hgt && hcl && ecl && pid
}

func (p passport) valid1() bool {
	_, byr := p.fields["byr"]
	_, iyr := p.fields["iyr"]
	_, eyr := p.fields["eyr"]
	_, hgt := p.fields["hgt"]
	_, hcl := p.fields["hcl"]
	_, ecl := p.fields["ecl"]
	_, pid := p.fields["pid"]
	exist := byr && iyr && eyr && hgt && hcl && ecl && pid
	return exist
}

func parse(data []string) []passport {
	p := NewPassport()
	var result []passport

	for _, l := range data {
		if l == "" {
			result = append(result, p)
			p = NewPassport()
			continue
		}
		reqs := strings.Split(l, " ")

		for _, r := range reqs {
			var key string
			var val string
			fmt.Sscanf(r, "%3s:%s", &key, &val)
			p.add(key, val)
		}
	}
	result = append(result, p)

	return result
}

func main() {
	data := readLines("input.txt")
	passports := parse(data)

	validA := 0
	for _, p := range passports {
		if p.valid1() {
			validA++
		}
	}
	fmt.Printf("Valid A: %d\n", validA)

	validB := 0
	for _, p := range passports {
		if p.valid2() {
			validB++
		}
	}
	fmt.Printf("Valid A: %d\n", validB)
}
