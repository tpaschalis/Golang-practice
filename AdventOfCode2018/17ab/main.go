package main

import (
	"fmt"
	"strconv"
	"strings"
)

var ground [2200][2200]string
var ymax, xmax int

func main() {
	data := strings.Split(input, "\n")
	ground = initialize(data)

	flow(500, 0)

	flowing, watered := 0, 0
	for x := 0; x < 2100; x++ {
		for y := 0; y < 2100; y++ {
			if ground[x][y] == "|" {
				flowing++
			} else if ground[x][y] == "~" {
				watered++
			}
		}
	}

	fmt.Println("Total water :", watered+flowing+1) // +1 for source
	fmt.Println("Water at rest :", watered)
}

func initialize(data []string) [2200][2200]string {
	var res [2200][2200]string
	for l := range data {
		pt1 := strings.Split(data[l], ", ")
		a, _ := strconv.Atoi(pt1[0][2:])
		pt2 := strings.Split(pt1[1][2:], "..")
		b, _ := strconv.Atoi(pt2[0])
		c, _ := strconv.Atoi(pt2[1])

		if pt1[0][0:1] == "x" {
			if c > ymax {
				ymax = c
			}
			for i := b; i <= c; i++ {
				res[a][i] = "#"
			}
		} else {
			if a > ymax {
				ymax = a
			}
			for i := b; i <= c; i++ {
				res[i][a] = "#"
			}
		}
	}
	return res
}

func avail(x, y int) bool {
	// Returns true if water can flow through the area
	return ground[x][y] == "" || ground[x][y] == "|"
}

func flow(x, y int) {
	// Args are starting points of the water fountain
	if y >= ymax {
		//  As the exercise mentions
		// "To prevent counting forever, ignore tiles with a y coordinate smaller than the smallest y"
		return
	} else if !avail(x, y) {
		return
	}
	// Let's specify how the water flows!
	if !avail(x, y+1) {
		leftBorder := x
		// Find out the limits of how 'left' or 'right' the water flows
		// The flow stops either on clay, or on empty space
		for avail(leftBorder, y) && !avail(leftBorder, y+1) {
			ground[leftBorder][y] = "|"
			leftBorder--
		}
		rightBorder := x + 1
		for avail(rightBorder, y) && !avail(rightBorder, y+1) {
			ground[rightBorder][y] = "|"
			rightBorder++
		}
		if avail(leftBorder, y+1) || avail(rightBorder, y+1) {
			// If the square downwards is available, flow the water through
			// the current square
			flow(leftBorder, y)
			flow(rightBorder, y)
		} else if ground[leftBorder][y] == "#" && ground[rightBorder][y] == "#" {
			// If the puddle at a certain depth is constrained left and right with clay
			// denote that there is water at rest
			for i := leftBorder + 1; i < rightBorder; i++ {
				ground[i][y] = "~"
			}
		}
	}
	// If the point in ground is empty, fill it with water and flow one square downwards.
	// If the square downwards has water at rest, flow to the same square (so maybe water
	// fills the sideways squares
	if ground[x][y] == "" {
		ground[x][y] = "|"
		flow(x, y+1)
		if ground[x][y+1] == "~" {
			flow(x, y)
		}
	}
}
