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

func wait(earliest, busId int) int {
	time := 0
	for {
		if time >= earliest {
			return time - earliest
		}
		time += busId
	}
}

func earliest(t travel) int {
	smallest := math.MaxInt32
	busId := 0

	for _, b := range t.buses {
		if b == 0 {
			continue
		}
		fmt.Printf("mod: %d MOD %d = %d\n", t.earliest, b, t.earliest % b)
		//mod := t.earliest % b
		w := wait(t.earliest, b)
		if w < smallest {
			smallest = w
			busId = b
		}
		fmt.Printf("res: %d\n", wait(t.earliest, b))
	}

	return smallest * busId
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

// LCM???
func (t travel) cong(step int) bool {
	for i, n := range t.buses {
		if n == 0 {
			continue
		}
		// step + i == 0 (mod n)
		// step == (n-1)*i (mod n)
		// step == ni-i (mod n)
		//c := (step + i) % n

		// step == n-i-1 (mod n)
		c := step % n
		// pos % n == 0
		//c := (n + step) % l
		//if c != i {
		//	return false
		//}
		if c != (n-i-1) {
			return false
		}
	}

	for i, n := range t.buses {
		if n == 0 {
			continue
		}
		//c := step % n
		fmt.Printf("%d == %d (mod %d)\n", step, n-i-1, n)
		//fmt.Printf("%d MOD %d == %d\n", step, n-i-1, c)
	}

	return true
}

func GCDRecursive(p, q int) int {
	if q == 0 {
		return p
	}

	r := p % q
	return GCDRecursive(q, r)
}

//def egcd(a, b):
//if a == 0:
//return (b, 0, 1)
//else:
//gcd, x, y = egcd(b % a, a)
//return (gcd, y - (b//a) * x, x)
func egcd(a, b int) (int, int, int) {
	if a == 0 {
		return b, 0, 1
	}

	gcd, x, y := egcd(b % a, a)
	return gcd, y - (b/a)*x, x
}

func inv(elem, ring int) int {
	fmt.Printf("inverse %d MOV %d\n", elem, ring)
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
	fmt.Printf("ys: %v\n", ys)
	fmt.Printf("zs: %v\n", zs)
	result := 0
	count := 0
	for i, n := range t.buses {
		if n == 0 {
			continue
		}
		a := n-i-1
		y := ys[count]
		z := zs[count]
		fmt.Printf("%d * %d * %d\n", a, y, z)
		result += a * y * z
		count++
	}
	return result % N
}

func main() {
	d := readLines("input.txt")

	p := parse(d)
	res := earliest(p)
	fmt.Printf("result: %d\n", res)
	fmt.Printf("%v\n", p)

	//step := 0
	//for {
	//	res := p.cong(step)
	//	if res {
	//		fmt.Printf("step: %d\n", step)
	//		break
	//	}
	//	step += 1
	//}

	c := chinese(p)
	fmt.Printf("chinese: %d\n", c+1)

	// 19 % 8 == 7
	// 31 % 8 == 6
	// ...

}
