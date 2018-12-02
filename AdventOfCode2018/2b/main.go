package main

import (
	"bufio"
	"fmt"
	"os"
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
	for i := 0; i < len(input); i++ {
		str1 := input[i]
		for j := i; j < len(input); j++ {
			if i == j {
				continue
			}
			str2 := input[j]

			differences := 0
			for k := 0; k < len(str2); k++ {
				if str1[k:k+1] != str2[k:k+1] {
					differences++
				}
			}
			if differences == 0 || differences == 1 {
				fmt.Println(input[i], input[j])
			}
		}

	}
}
