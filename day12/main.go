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

type ship struct {
	x         int
	y         int
	direction byte
	wx        int
	wy        int
}

type instruction struct {
	action   byte
	argument int
}

func (s *ship) forward(n int) {
	switch s.direction {
	case 'N':
		s.y += n
	case 'S':
		s.y -= n
	case 'E':
		s.x += n
	case 'W':
		s.x -= n
	}
}

func (s *ship) forwardw(n int) {
	s.x += n * s.wx
	s.y += n * s.wy
}

func (s *ship) left() {
	switch s.direction {
	case 'N':
		s.direction = 'W'
	case 'S':
		s.direction = 'E'
	case 'E':
		s.direction = 'N'
	case 'W':
		s.direction = 'S'
	}
}

func (s *ship) right() {
	switch s.direction {
	case 'N':
		s.direction = 'E'
	case 'S':
		s.direction = 'W'
	case 'E':
		s.direction = 'S'
	case 'W':
		s.direction = 'N'
	}
}

func (s *ship) rotate(dir byte, deg int) {
	nr90s := deg / 90

	for i := 0; i < nr90s; i++ {
		switch dir {
		case 'L':
			s.left()
		case 'R':
			s.right()
		}
	}
}

func sincos(angle int) (int, int) {
	switch angle {
	case 0:
		return 0, 1
	case 90:
		return 1, 0
	case 180:
		return 0, -1
	case 270:
		return -1, 0
	}

	return 0, 0
}

func (s *ship) rot(angle int) {
	sin, cos := sincos(angle)
	newx := s.wx*cos - s.wy*sin
	newy := s.wx*sin + s.wy*cos
	s.wx = newx
	s.wy = newy
}

func (s *ship) step(inst instruction) {
	switch inst.action {
	case 'N':
		s.y += inst.argument
	case 'S':
		s.y -= inst.argument
	case 'E':
		s.x += inst.argument
	case 'W':
		s.x -= inst.argument
	case 'L':
		s.rotate(inst.action, inst.argument)
	case 'R':
		s.rotate(inst.action, inst.argument)
	case 'F':
		s.forward(inst.argument)
	}
}

func (s *ship) stepw(inst instruction) {
	switch inst.action {
	case 'N':
		s.wy += inst.argument
	case 'S':
		s.wy -= inst.argument
	case 'E':
		s.wx += inst.argument
	case 'W':
		s.wx -= inst.argument
	case 'L':
		s.rot(inst.argument)
	case 'R':
		s.rot(360 - inst.argument)
	case 'F':
		s.forwardw(inst.argument)
	}
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func parse(lines []string) []instruction {
	var result []instruction

	for _, l := range lines {
		f := l[0]
		arg, _ := strconv.Atoi(l[1:])
		result = append(result, instruction{
			action:   f,
			argument: arg,
		})
	}

	return result
}

func main() {
	d := readLines("input.txt")
	instructions := parse(d)

	s1 := ship{
		x:         0,
		y:         0,
		direction: 'E',
		wx:        10,
		wy:        1,
	}
	for _, i := range instructions {
		s1.step(i)
	}

	s2 := ship{
		x:         0,
		y:         0,
		direction: 'E',
		wx:        10,
		wy:        1,
	}
	for _, i := range instructions {
		s2.stepw(i)
	}

	fmt.Printf("manhattan A: %d\n", abs(s1.x)+abs(s1.y))
	fmt.Printf("manhattan B: %d\n", abs(s2.x)+abs(s2.y))
}
