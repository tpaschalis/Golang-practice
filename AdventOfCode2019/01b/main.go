package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	res := 0
	for scanner.Scan() {
		m, err := strconv.Atoi(scanner.Text())
		if err != nil {
			continue
		}
		res += calcTotal(m)
	}
	fmt.Println(res)
}

// Now, we also have to calculate the fuel required
// to launch the fuel itself!
func calcFuel(i int) int {
	return (i / 3) - 2
}

func calcTotal(i int) (res int) {
	cur := calcFuel(i)
	for cur >= 0 {
		res += cur
		cur = calcFuel(cur)
	}
	return res
}
