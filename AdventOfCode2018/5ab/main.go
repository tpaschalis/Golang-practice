package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var Alphabet string

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var input string
	for scanner.Scan() {
		input = scanner.Text()      // Only one line of input
	}
	Alphabet = "abcdefghijklmnopqrstuvwxyz"

	fmt.Println("Pt 1 : Fully reacting the initial polymer gives us :", len(react(input)))

	min := len(input)
	for i := 0; i < len(Alphabet); i++ {
		l := Alphabet[i : i+1]
		candidate := input
		candidate = strings.Replace(candidate, l, "", -1)
		candidate = strings.Replace(candidate, strings.ToUpper(l), "", -1)
		candidateSize := len(react(candidate))
		if candidateSize < min { min = candidateSize }
	}
	fmt.Println("Pt 2 : Minimum size achieved by removing just one 'letter' :", min)
}

func react(s string) string {
	sInit := s
	for i := 0; i < len(Alphabet); i++ {
		l := Alphabet[i : i+1]
		s = strings.Replace(s, l+strings.ToUpper(l), "", -1)
		s = strings.Replace(s, strings.ToUpper(l)+l, "", -1)
	}
	if sInit == s {
		return sInit
	} else {
		return react(s)
	}
}
