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

	var counter int
    var current int
	var line string
	fmt.Println(line)
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := sc.Text()
        current, err = strconv.Atoi(line)
        counter = counter + current
	}
	if err := sc.Err(); err != nil {
		panic(err)
	}

	fmt.Println(counter)
}
