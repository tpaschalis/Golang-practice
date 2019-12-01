package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	//scanner := bufio.NewScanner(os.Stdin)
	//var input []string
	//for scanner.Scan() {
	//	input = append(input, scanner.Text())
	//}

	res := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		m, err := strconv.Atoi(scanner.Text())
		if err != nil {
			continue
		}
		res += calcFuel(m)
	}

	fmt.Println(res)
}

// Fuel required to launch a given module is based on its mass.
// Specifically, to find the fuel required for a module, take its mass,
// divide by three, round down, and subtract 2.
func calcFuel(i int) int {
	return (i / 3) - 2
}
