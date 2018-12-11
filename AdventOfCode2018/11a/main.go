package main

import "fmt"

const sn = 8868

func main() {

	grid := make([][]int, 300)
	for i := range grid {
		grid[i] = make([]int, 300)
	}

	for i := range grid {
		for j := range grid[0] {
			grid[i][j] = fillGrid(i, j)
		}
	}

	max, xm, ym, nrg, xM, yM, sM := 0, 0, 0, 0, 0, 0, 0
	// Part 1
	for i := 0; i < 298; i++ {
		for j := 0; j < 298; j++ {
			cur := sumReg(grid, i, j)
			if cur > max { max, xm, ym = cur, i, j }
		}
	}

	// Part 2
	max = 0
	for i := 0; i < 299; i++ {
		for j := 0; j < 299; j++ {
			nrg = 0
			upto := Minv(299-i, 299-j)
			for s := 0; s <= upto; s++ {
				nrg -= grid[i+s][j+s] // This will be added twice in the next loop, so we subtract one (it's the SxS box corner)
				for k := 0; k <= s; k++ {
					nrg += grid[i+s][j+k]   // Add the new "row" and "line" values of the larger square
					nrg += grid[i+k][j+s]
					if nrg > max {
						max, xM, yM, sM = nrg, i, j, s
					}
				}
			}
		}
	}

	fmt.Println("Part one :", max, ", at :", xm, ym)
	fmt.Println("Part two :", max, ", at :", xM, yM, sM)
}

func fillGrid(x, y int) int {
	res := (int((((x+10)*y + sn) * (x + 10)) % 1000 / 100)) - 5
	return res
}
func sumReg(g [][]int, x, y int) int {
	res := g[x][y] + g[x][y+1] + g[x][y+2] + g[x+1][y] + g[x+1][y+1] + g[x+1][y+2] + g[x+2][y] + g[x+2][y+1] + g[x+2][y+2]
	return res
}
func Minv(x, y int) int {
	if x < y {
        return x
	}
	return y
}
