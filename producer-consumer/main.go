package main

import (
	"fmt"
)

var fin = make(chan bool)
var stream = make(chan int)

func produce() {
	for i := 0; i < 100000; i++ {
		stream <- i
	}
	fin <- true
}

func consume() {
	for {
		data := <-stream
		fmt.Println(data)
	}
}

func main() {
	go produce()
	go consume()
	<-fin
}
