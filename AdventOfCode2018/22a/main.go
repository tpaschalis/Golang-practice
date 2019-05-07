package main

import "fmt"

type rcave struct {
	x       int
	y       int
	erosion int
	ground  int
}

func main() {
	//targetX, targetY := 10, 10
	//depth := 510

	targetX, targetY := 10, 715
	depth := 3339

	cave := make([][]rcave, targetX+1)
	for i := range cave {
		cave[i] = make([]rcave, targetY+1)
	}

	dangerlvl := 0
	// Initialize the cave and sides (first two rules)
	for x := range cave {
		for y := range cave[x] {
			cave[x][y].x = x
			cave[x][y].y = y
			if x == 0 && y == 0 {
				cave[x][y].erosion = (0 + depth) % 20183
				cave[x][y].ground = cave[x][y].erosion % 3
				dangerlvl += cave[x][y].ground
			}
			if y == 0 && x != 0 {
				cave[x][y].erosion = (x*16807 + depth) % 20183
				cave[x][y].ground = cave[x][y].erosion % 3
				dangerlvl += cave[x][y].ground
			}
			if x == 0 && y != 0 {
				cave[x][y].erosion = (y*48271 + depth) % 20183
				cave[x][y].ground = cave[x][y].erosion % 3
				dangerlvl += cave[x][y].ground
			}
		}
	}

	// Filling in the cave (third rule)
	for x := 1; x <= targetX; x++ {
		for y := 1; y <= targetY; y++ {
			if x == targetX && y == targetY {
				break
			}

			cave[x][y].erosion = (cave[x-1][y].erosion*cave[x][y-1].erosion + depth) % 20183
			cave[x][y].ground = cave[x][y].erosion % 3
			dangerlvl += cave[x][y].ground
		}
	}

	// Target coordinates (fourth rule)
	cave[targetX][targetY].erosion = (0 + depth) % 20183
	cave[targetX][targetY].ground = cave[targetX][targetY].erosion % 3
	dangerlvl -= cave[0][0].ground + cave[targetX][targetY].ground

	fmt.Println("Total danger lvl :", dangerlvl)
	//display(cave)
}

func display(c [][]rcave) {

	for x := range c {
		for y := range c[x] {
			fmt.Printf("%s", printchar(c[y][x].ground))
		}
		fmt.Printf("\n")
	}
}

func printchar(n int) string {
	if n == 0 { return "." }
	if n == 1 { return "=" }
	if n == 2 { return "|" }
	panic("n should have been modulo 3")
	return "err"
}
