package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")
	for i := 0; i <= 50; i++ {
		by3 := i%3 == 0
		by5 := i%5 == 0
		if by3 && by5 {
			println("FizzBuzz!")
		} else if by3 {
			println("Fizz")
		} else if by5 {
			println("Buzz")
		} else {
			println(i)
		}
	}
}
