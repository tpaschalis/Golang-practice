package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	if scanner.Err() != nil {
		fmt.Println("It's done")
	}

	const size = 1000
	var grid [size][size]int

	var total int
	for i := 0; i < len(input); i++ {
		current := input[i]
		parts := strings.FieldsFunc(current, Split)
		xpos, err := strconv.Atoi(strings.TrimSpace(parts[1]))
		ypos, err := strconv.Atoi(strings.TrimSpace(parts[2]))
		xdim, err := strconv.Atoi(strings.TrimSpace(parts[3]))
		ydim, err := strconv.Atoi(strings.TrimSpace(parts[4]))
		if err != nil {
			panic(err)
		}
		fmt.Println(parts, xdim)

		for j := 0; j < xdim; j++ {
			for k := 0; k < ydim; k++ {
				if grid[ypos+k][xpos+j] == 1 {
					total++
				}
				grid[ypos+k][xpos+j]++
			}
		}
	}

	for i := 0; i < len(input); i++ {
		var overlapping = false
		current := input[i]
		parts := strings.FieldsFunc(current, Split)
		xpos, err := strconv.Atoi(strings.TrimSpace(parts[1]))
		ypos, err := strconv.Atoi(strings.TrimSpace(parts[2]))
		xdim, err := strconv.Atoi(strings.TrimSpace(parts[3]))
		ydim, err := strconv.Atoi(strings.TrimSpace(parts[4]))
		if err != nil {
			panic(err)
		}
		for j := 0; j < xdim; j++ {
			for k := 0; k < ydim; k++ {
				if grid[ypos+k][xpos+j] > 1 {
					overlapping = true
				}
			}
		}
		if overlapping == false {
			fmt.Println(parts[0])
		}
	}

}

func Split(r rune) bool {
	return r == '@' || r == ':' || r == ',' || r == 'x'
}
