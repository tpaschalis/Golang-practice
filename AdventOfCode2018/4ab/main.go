package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	// We're lucky that a simple alphabetical sort
	// is also the correct time sort :P
    // Another awesome solution at https://github.com/seankhliao/adventofcode2018/blob/master/04/main.go
	sort.Strings(input)
	r, _ := regexp.Compile("#[0-9]*")
	q, _ := regexp.Compile(":[0-9]*")

	var guards []string
	var currentGuard string
	var st, en string
	var startTime, endTime int

	watchLog := make(map[string][60]int)

	for i := 0; i < len(input); i++ {
		if r.FindString(input[i]) != "" {
			guards = append(guards, r.FindString(input[i]))
			currentGuard = r.FindString(input[i])
		}

		// I assume a single guard at a time, and that
		// he will "fall asleep" and "wake up" in matching ordered pairs
		if strings.Contains(input[i], "falls") {
			st = q.FindString(input[i])[1:]
			startTime, _ = strconv.Atoi(st)
		}
		if strings.Contains(input[i], "wakes") {
			en = q.FindString(input[i])[1:]
			endTime, _ = strconv.Atoi(en)

			// The map we built is not addressable, so we're using this temp variable
			var temp [60]int
			temp = watchLog[currentGuard]
			for j := startTime; j < endTime; j++ {
				temp[j]++
			}
			watchLog[currentGuard] = temp
		}
	}

	// Strategy 1
	// Who sleeps the most?
	// Find the guard that has the most minutes asleep. What minute does that guard spend asleep the most?
	var sleepyHead string
	max, idx := 0, 0
	for _, v := range guards {
		if sum(watchLog[v]) > max { max, sleepyHead = sum(watchLog[v]), v }
	}
	max = 0
	for i, v := range watchLog[sleepyHead] {
		if v > max { max, idx = v, i }
	}
	fmt.Println("Sleepy Head guard", sleepyHead, "spent most days sleeping on 00 :", idx)

	// Strategy 2
	// Of all guards, which guard is most frequently asleep on the same minute?
	max = 0
	for _, v := range guards {
		for i, m := range watchLog[v] {
			if m > max { max, sleepyHead, idx = m, v, i	}
		}
	}
	fmt.Println("Guard", sleepyHead, "kept falling asleep on minute", idx)
}

func sum(input [60]int) int {
	sum := 0
	for _, i := range input {
		sum += i
	}
	return sum
}
