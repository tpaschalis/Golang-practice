package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var input []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	seen := map[string]int{}
	var hashTwice, hashThrice int

	for i := 0; i < len(input); i++ {
		seenTwice, seenThrice := false, false
		seen = make(map[string]int)
		for j := 0; j < len(input[i]); j++ {
			current := input[i]
			seen[current[j:j+1]]++
		}
		for _, value := range seen {
			if value == 2 {
				seenTwice = true
			}
			if value == 3 {
				seenThrice = true
			}
		}
		if seenTwice {
			hashTwice++
		}
		if seenThrice {
			hashThrice++
		}
		//fmt.Println(seen)
	}
	fmt.Println(hashTwice)
	fmt.Println(hashThrice)
	fmt.Println(hashTwice * hashThrice)

	//fmt.Println(input)
}
