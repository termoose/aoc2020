package main

import (
	"fmt"
)

func main() {
	// number -> spoken
	numbers := make(map[int][]int)
	//d := []int{3, 1, 2}
	d := []int{6, 3, 15, 13, 1, 0}
	for i, n := range d {
		numbers[n] = []int{i}
	}

	for key, val := range numbers {
		fmt.Printf("%d spoken: %v\n", key, val)
	}

	for i := len(d); i < 30000000; i++ {
		last := d[i-1]
		nrs, _ := numbers[last]
		//fmt.Printf("%d last %d: %v ok: %v\n", i, last, nrs, ok)

		if len(nrs) < 2 {
			d = append(d, 0)
			numbers[0] = append(numbers[0], i)
		} else {
			prev := nrs[len(nrs)-1]
			prevprev := nrs[len(nrs)-2]
			diff := prev - prevprev
			//fmt.Printf("diff: %d\n", diff)
			d = append(d, diff)
			numbers[diff] = append(numbers[diff], i)
		}
	}

	fmt.Printf("last number: %v\n", d[len(d)-1])
}
