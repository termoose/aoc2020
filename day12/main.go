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
	newx := s.wx * cos - s.wy * sin
	newy := s.wx * sin + s.wy * cos
	s.wx = newx
	s.wy = newy
}

func (s *ship) step(action string) {
	f := action[0]
	arg, _ := strconv.Atoi(action[1:])

	switch f {
	case 'N':
		s.y += arg
	case 'S':
		s.y -= arg
	case 'E':
		s.x += arg
	case 'W':
		s.x -= arg
	case 'L':
		s.rotate(f, arg)
	case 'R':
		s.rotate(f, arg)
	case 'F':
		s.forward(arg)
	}
}

func (s *ship) stepw(action string) {
	f := action[0]
	arg, _ := strconv.Atoi(action[1:])

	switch f {
	case 'N':
		s.wy += arg
	case 'S':
		s.wy -= arg
	case 'E':
		s.wx += arg
	case 'W':
		s.wx -= arg
	case 'L':
		s.rot(arg)
	case 'R':
		s.rot(360-arg)
	case 'F':
		s.forwardw(arg)
	}
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func main() {
	d := readLines("input.txt")
	s1 := ship{
		x:         0,
		y:         0,
		direction: 'E',
		wx:        10,
		wy:        1,
	}
	for _, l := range d {
		s1.step(l)
	}

	s2 := ship{
		x:         0,
		y:         0,
		direction: 'E',
		wx:        10,
		wy:        1,
	}
	for _, l := range d {
		s2.stepw(l)
	}

	fmt.Printf("manhattan A: %d\n", abs(s1.x)+abs(s1.y))
	fmt.Printf("manhattan B: %d\n", abs(s2.x)+abs(s2.y))
}
