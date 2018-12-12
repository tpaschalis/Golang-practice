package main

import "fmt"

type pattern struct {
	p []int
	r int
}

func main() {

	state := make(map[int]int)
	input := []pattern{                                     // Instead of parsing the input, a simple search and replace
		pattern{p: []int{1, 0, 0, 0, 0}, r: 0}, //  01         provides the necessary structure easily. :)
		pattern{p: []int{1, 0, 0, 1, 1}, r: 1}, //  02         At the expense of loc of course :P
		pattern{p: []int{0, 0, 0, 0, 1}, r: 0}, //  03
		pattern{p: []int{0, 0, 0, 1, 0}, r: 0}, //  04
		pattern{p: []int{0, 0, 0, 1, 1}, r: 1}, //  05
		pattern{p: []int{1, 0, 1, 0, 1}, r: 0}, //  06
		pattern{p: []int{0, 1, 0, 0, 0}, r: 1}, //  07
		pattern{p: []int{1, 1, 0, 1, 0}, r: 0}, //  08
		pattern{p: []int{0, 0, 1, 0, 1}, r: 0}, //  09
		pattern{p: []int{0, 1, 1, 0, 1}, r: 1}, //  10
		pattern{p: []int{1, 1, 1, 0, 1}, r: 1}, //  11
		pattern{p: []int{0, 1, 0, 1, 1}, r: 0}, //  12
		pattern{p: []int{0, 0, 0, 0, 0}, r: 0}, //  13         This is the line of input that makes sure the structure will stabilize
		pattern{p: []int{1, 1, 1, 1, 1}, r: 1}, //  14
		pattern{p: []int{1, 1, 1, 0, 0}, r: 0}, //  15
		pattern{p: []int{1, 1, 0, 0, 1}, r: 1}, //  16
		pattern{p: []int{1, 0, 1, 1, 1}, r: 1}, //  17
		pattern{p: []int{1, 0, 1, 0, 0}, r: 0}, //  18
		pattern{p: []int{0, 0, 1, 1, 1}, r: 0}, //  19
		pattern{p: []int{0, 0, 1, 0, 0}, r: 0}, //  20
		pattern{p: []int{0, 1, 0, 0, 1}, r: 1}, //  21
		pattern{p: []int{0, 1, 1, 0, 0}, r: 1}, //  22
		pattern{p: []int{1, 1, 0, 0, 0}, r: 1}, //  23
		pattern{p: []int{0, 1, 0, 1, 0}, r: 1}, //  24
		pattern{p: []int{0, 1, 1, 1, 0}, r: 1}, //  25
		pattern{p: []int{1, 0, 0, 1, 0}, r: 0}, //  26
		pattern{p: []int{1, 1, 1, 1, 0}, r: 0}, //  27
		pattern{p: []int{0, 1, 1, 1, 1}, r: 1}, //  28
		pattern{p: []int{1, 0, 1, 1, 0}, r: 1}, //  29
		pattern{p: []int{1, 1, 0, 1, 1}, r: 0}, //  30
		pattern{p: []int{0, 0, 1, 1, 0}, r: 0}, //  31
		pattern{p: []int{1, 0, 0, 0, 1}, r: 1}, //  32
	}

	init := `##...#......##......#.####.##.#..#..####.#.######.##..#.####...##....#.#.####.####.#..#.######.##...` // I decided to work with 0/1 instead of ./# don't think it matters so much...
	state = initialize(init)                                                                                       // un-initialized values are zero by default, which is nice
	state[-1], state[-2] = 0, 0

	for i := 0; i < 2000; i++ {                             // Change to 20 for Part One answer
		nstate := invoke(state, input)
		exitCondition := true
		for j, _ := range nstate {
			if nstate[j] != state[j-1] {
				exitCondition = false
			}
		}
		state = nstate
		if exitCondition {
			fmt.Println("Stabilized on generation :", i+1)
			break
		}
		for j := 0; j < 200; j++ {
			fmt.Printf("%d", state[j])
		}
	}
	sum, idx := 0, 0
	for k, v := range state {
		if v == 1 {
			sum += k
			idx++
		}
	}
	fmt.Println("\nSum of first stabilized plant pots : ", sum)
	fmt.Println("\nSum of future 50B-th generation : ", (uint64(50000000000)-99)*uint64(idx)+uint64(sum), idx) // It was stabilized on Generation 99, so that's the offset we add to the stabilized sum
}

func invoke(m map[int]int, pt []pattern) map[int]int {

	max := 0
	for k, _ := range m {
		if (m[k] == 1) && (k > max) { max = k }
	}

	n := make(map[int]int)
	for i := -2; i <= max+2; i++ {
		n[i] = 0
	}

	var cur int
	for i := range n {
		for j := range pt {
			check := pt[j].p
			poss := pt[j].r
			match := true
			for k := -2; k <= 2; k++ {
				cur = i + k
				if m[cur] != check[k+2] { match = false }
			}
			if match == true {
				n[i] = poss
			}
		}
	}
	return n
}

func initialize(input string) map[int]int {
	res := make(map[int]int)
	for i := range input {
		if input[i:i+1] == "#" {
			res[i] = 1
		} else if input[i:i+1] == "." {
			res[i] = 0
		}
	}
	return res
}

// Feels dirty, but this is how I initially validated that the structure was stabilized
// This was inside the 'invoke' function
//min := max
//for k, _ := range m {
//	if (m[k] == 1) && (k < min) {
//		min = k
//	}
//}
//fmt.Println("  Length : ", max-min)

//n := make(map[int]int)
//for i := min - 2; i <= max+2; i++ {
//	n[i] = 0
//}
