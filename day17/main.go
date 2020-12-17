package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

type point struct {
	x int
	y int
	z int
	w int
}

type world struct {
	space   map[point]bool
	bbStart point
	bbStop  point
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

func charToBool(c byte) bool {
	switch c {
	case '.':
		return false
	case '#':
		return true
	}

	return false
}

func boolToChar(c bool) string {
	if c {
		return string('#')
	}

	return string('.')
}

func parse(d []string) world {
	result := world{
		space: make(map[point]bool),
	}
	for y, line := range d {
		//maxX = len(line)
		for x, c := range line {
			p := point{
				x: x,
				y: y,
				z: 0,
				w: 0,
			}

			result.space[p] = charToBool(byte(c))
		}
	}

	result.findBoundingBox()

	return result
}

func (ww *world) findBoundingBox() {
	start := point{math.MaxInt16, math.MaxInt16, math.MaxInt16, math.MaxInt16}
	stop := point{-math.MaxInt16, -math.MaxInt16, -math.MaxInt16, -math.MaxInt16}

	for p, _ := range ww.space {
		if p.x < start.x {
			start.x = p.x
		}
		if p.y < start.y {
			start.y = p.y
		}
		if p.z < start.z {
			start.z = p.z
		}
		if p.w < start.w {
			start.w = p.w
		}

		if p.x > stop.x {
			stop.x = p.x
		}
		if p.y > stop.y {
			stop.y = p.y
		}
		if p.z > stop.z {
			stop.z = p.z
		}
		if p.w > stop.w {
			stop.w = p.w
		}
	}

	ww.bbStart = point{start.x - 1, start.y - 1, start.z - 1, start.w - 1}
	ww.bbStop = point{stop.x + 1, stop.y + 1, stop.z + 1, stop.w + 1}
}

func (ww world) activeInactive(p point) (int, int) {
	active := 0
	inactive := 0

	for x := p.x - 1; x <= p.x+1; x++ {
		for y := p.y - 1; y <= p.y+1; y++ {
			for z := p.z - 1; z <= p.z+1; z++ {
				for w := p.w - 1; w <= p.w+1; w++ {
					if p.x == x && p.y == y && p.z == z && p.w == w {
						continue
					}

					pos := point{x, y, z, w}
					val := ww.space[pos]

					if val {
						active++
					} else {
						inactive++
					}
				}
			}
		}
	}

	return active, inactive
}

func (ww world) simulate() world {
	result := world{
		space: make(map[point]bool),
	}

	for z := ww.bbStart.z; z <= ww.bbStop.z; z++ {
		for y := ww.bbStart.y; y <= ww.bbStop.y; y++ {
			for x := ww.bbStart.x; x <= ww.bbStop.x; x++ {
				for w := ww.bbStart.w; w <= ww.bbStop.w; w++ {
					p := point{x, y, z, w}
					active, _ := ww.activeInactive(p)
					val := ww.space[p]
					result.space[p] = val

					if val && (active >= 2 && active <= 3) {
						result.space[p] = true
					} else {
						result.space[p] = false
					}

					if !val && active == 3 {
						result.space[p] = true
					} else {
						//result.space[p] = false
					}
				}
			}
		}
	}

	result.findBoundingBox()
	return result
}

func (ww world) print() {
	for w := ww.bbStart.w; w <= ww.bbStop.w; w++ {
		for z := ww.bbStart.z; z <= ww.bbStop.z; z++ {
			fmt.Printf("z = %d, w = %d\n", z, w)
			for y := ww.bbStart.y; y <= ww.bbStop.y; y++ {
				for x := ww.bbStart.x; x <= ww.bbStop.x; x++ {
					p := point{x, y, z, w}
					val := ww.space[p]
					fmt.Printf("%s", boolToChar(val))
				}
				fmt.Printf("\n")
			}
			fmt.Printf("\n")
		}
	}
}

func (ww world) active() int {
	result := 0
	for z := ww.bbStart.z; z <= ww.bbStop.z; z++ {
		for y := ww.bbStart.y; y <= ww.bbStop.y; y++ {
			for x := ww.bbStart.x; x <= ww.bbStop.x; x++ {
				for w := ww.bbStart.w; w <= ww.bbStop.w; w++ {
					p := point{x, y, z, w}
					val := ww.space[p]
					if val {
						result++
					}
				}
			}
		}
	}
	return result
}

func main() {
	data := readLines("input.txt")
	fmt.Printf("%v\n", data)

	res := parse(data)
	fmt.Printf("%v\n", res)

	for i := 0; i < 6; i++ {
		res = res.simulate()
	}

	res.print()
	fmt.Printf("active: %d\n", res.active())
}
