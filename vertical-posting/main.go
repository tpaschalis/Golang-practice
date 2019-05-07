package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Exactly one argument expected!")
		os.Exit(1)
	}
	word := os.Args[1]
	var sb strings.Builder
	_ = sb
	sb.WriteString("```\n")
	for i, _ := range word {
		sb.WriteString(word[i:i+1] + " ")
	}
	sb.WriteString("\n")
	for i, _ := range word {
		if i >= 1 {
			sb.WriteString(word[i:i+1] + strings.Repeat(" ", i*2-1) + word[i:i+1] + "\n")
		}
	}
	sb.WriteString("\n```")
	fmt.Println(sb.String())
}
