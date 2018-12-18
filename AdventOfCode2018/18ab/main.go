package main

import "fmt"
import "strings"

//import "time"

var rows, cols int

func main() {
	rows, cols = 50, 50

	data := strings.Split(input, "\n")

	area := make([][]string, rows+3)
	for i := range area {
		area[i] = make([]string, cols+3)
	}

	for i := range data {
		l := data[i]
		for j := range data[i] {
			area[i+1][j+1] = l[j : j+1]
		}
	}

	results := make([]int, 1)

	for t := 0; t < 5000; t++ {
		// Essentially, copy(oldarea, area)
		oldarea := make([][]string, rows+3)
		for i := range oldarea {
			oldarea[i] = make([]string, cols+3)
			copy(oldarea[i], area[i])
		}

		// Simulate the next minute
		for i := 1; i <= rows+1; i++ {
			for j := 1; j <= cols+1; j++ {
				area[i][j] = adjacent(i, j, oldarea)
			}
		}
		results = append(results, checkProduce(area))

		//    Visual Inspection
		//    We notice that early on, not after 100-150 minutes, a cyclical pattern comes forward
		// time.Sleep(50*time.Millisecond)
		// fmt.Println("\033[2J")            // Special code to clear terminal in unix-y environments
		// fmt.Println("Minutes :", t)
		// display(area)

	}
	fmt.Println("Part One :", results[10])
	fmt.Println("The structure oscillates after ~100-150 cycles, with a frequency of", findFrequency(results))
	point := 1000 + (1000000000-1000)%28
	fmt.Println("Part Two: After 1000000000 minutes, it will be", results[point])
}

func display(m [][]string) {
	for i := range m {
		for j := range m[i] {
			fmt.Printf("%s", m[i][j])
		}
		fmt.Printf("\n")
	}
}

func adjacent(x, y int, m [][]string) string {

	open, tree, yard := 0, 0, 0
	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			if i == x && j == y {
				continue
			}
			if m[i][j] == "." {
				open++
			}
			if m[i][j] == "|" {
				tree++
			}
			if m[i][j] == "#" {
				yard++
			}
		}
	}
	//fmt.Println(open, tree, yard)

	if m[x][y] == "." && tree >= 3 {
		return "|"
	} else if m[x][y] == "." {
		return "."
	}

	if m[x][y] == "|" && yard >= 3 {
		return "#"
	} else if m[x][y] == "|" {
		return "|"
	}

	if m[x][y] == "#" && yard >= 1 && tree >= 1 {
		return "#"
	} else if m[x][y] == "#" {
		return "."
	}
	return ""
}

func checkProduce(m [][]string) int {
	wood, lumber := 0, 0
	for i := range m {
		for j := range m[i] {
			if m[i][j] == "|" {
				wood++
			}
			if m[i][j] == "#" {
				lumber++
			}
		}
	}
	return wood * lumber
}

func findFrequency(t []int) int {

	min, freq := 10000000, 0
	for f := 1; f <= 60; f++ {
		sum := 0
		for i := 61; i < len(t)-1; i++ {
			sum = sum + Abs(t[i]-t[i-f])
		}
		if sum < min {
			min = sum
			freq = f
		}
	}
	return freq
}
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
