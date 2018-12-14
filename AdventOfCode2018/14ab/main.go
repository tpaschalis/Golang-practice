package main

import "fmt"
import "strings"

func main() {

	scoreboard := []int{3, 7, 1, 0}
	var elf1, elf2, cook1, cook2 int
	elf1, elf2 = 0, 1

	// To simplify things, I hope that the `input` sequence
	//will not appear before `input` amount of iterations :P
	input := 920831
	i := 0
	for {
		i++
		cook1 = scoreboard[elf1]
		cook2 = scoreboard[elf2]
		if cook1+cook2 < 10 {
			scoreboard = append(scoreboard, cook1+cook2)
		} else {
			scoreboard = append(scoreboard, (int(cook1+cook2)/10)%10)
			scoreboard = append(scoreboard, (cook1+cook2)%10)
		}
		elf1 = (1 + elf1 + scoreboard[elf1]) % (len(scoreboard))
		elf2 = (1 + elf2 + scoreboard[elf2]) % (len(scoreboard))

		if i > 10 {
			// Silly, but if you start reading the scoreboard sooner, you'll get out of slice
			// Maybe comparing array-to-array would be faster, or concatenating to a big int? I'll have to see
			current := strings.Trim(strings.Replace(fmt.Sprint(scoreboard[i-6:i]), " ", "", -1), "[]")
			if current == "920831" {
				fmt.Println("Part Two :", i-6)
				break
			}
		}
	}
	fmt.Println("Part One :", scoreboard[input:input+10])
}
