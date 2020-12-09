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

type instruction struct {
	opcode string
	arg    int
}

type program struct {
	ip   int
	acc  int
	code []instruction
}

func NewProgram(data []string) program {
	var code []instruction

	for _, line := range data {
		var i instruction
		_, _ = fmt.Sscanf(line, "%s %d", &i.opcode, &i.arg)

		code = append(code, i)
	}

	return program{
		ip:   0,
		acc:  0,
		code: code,
	}
}

func (p *program) step(ip int) bool {
	i := p.code[ip]

	switch i.opcode {
	case "nop":
		p.ip = ip + 1
	case "acc":
		p.acc += i.arg
		p.ip = ip + 1
	case "jmp":
		p.ip += i.arg
	}

	return p.ip == len(p.code)
}

func (p *program) run() (int, bool) {
	visited := make(map[int]bool)

	for {
		terminate := p.step(p.ip)

		if terminate {
			return p.acc, terminate
		}
		if _, v := visited[p.ip]; v {
			return p.acc, terminate
		}

		visited[p.ip] = true
	}
}

func (p *program) reset() {
	p.ip = 0
	p.acc = 0
}

func (p program) mutate(ip int) program {
	newProgram := program{
		acc:  0,
		ip:   0,
		code: make([]instruction, len(p.code)),
	}
	copy(newProgram.code, p.code)

	switch p.code[ip].opcode {
	case "jmp":
		newProgram.code[ip].opcode = "nop"
	case "nop":
		newProgram.code[ip].opcode = "jmp"
	}

	return newProgram
}

func main() {
	d := readLines("input.txt")
	prog := NewProgram(d)

	res, terminated := prog.run()
	fmt.Printf("Result A: %d Terminated: %t\n", res, terminated)
	prog.reset()

	counter := 0
	for {
		progMutated := prog.mutate(counter)
		res, term := progMutated.run()

		if term {
			fmt.Printf("Result B: %d Terminated: %t\n", res, term)
			break
		}

		prog.reset()
		counter++
	}
}
