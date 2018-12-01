package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, err := os.OpenFile("input", os.O_RDONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer f.Close()

    var input []string
    var current int
    var value int

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		input = append(input, sc.Text())
	}
	if err := sc.Err(); err != nil {
		panic(err)
	}

    seen := map[int]bool{}
    for i:=0; seen[value] == false; i = (i+1) % len(input) {
        seen[value] = true
        current, err = strconv.Atoi(input[i])
        value = value + current
    }
    fmt.Println(value)

}
