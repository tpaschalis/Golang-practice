package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "strconv"
)

func main() {
    var x, y []int
    scanner := bufio.NewScanner(os.Stdin)

    // Code is pretty awful, and I'm sorely missing map/reduce in Go.
    // Still tho...

    xGrid, yGrid := 0,0
    for scanner.Scan() {
        in := strings.Split(scanner.Text(), ", ")
        xs, _ := strconv.Atoi(in[0])
        ys, _ := strconv.Atoi(in[1])
        if xs > xGrid { xGrid = xs }
        if ys > yGrid { yGrid = ys }
        x, y = append(x, xs), append(y, ys)
    }

    grid := make([][]int, yGrid+1)
    for i := range grid {
        grid[i] = make([]int, xGrid+1)
    }

    areas := make(map[int]int)
    var regn, currId, max int

    for i := range grid {
        for j:= range grid[0] {
            currId = findNearest(i, j, x, y)
            grid[i][j] = currId
        }
    }

    for i:= range grid {
        for j := range grid[0] {
            // Part 1
            currId = grid[i][j]
            areas[currId]++
            // Part 2
            currDist := findTotal(i, j, x, y)
            if currDist < 10000 {
                regn ++
            }
        }
    }
    for i:= range grid {
        for j := range grid[0] {
            if (i == 0 ) || (j == 0) || (i == len(grid)) || (j == len(grid)) {
                currId = grid[i][j]
                areas[currId] = -1
            }
        }
    }
    for _, v := range areas {
        if v > max { max = v }
    }

    fmt.Println("It might be dangerous!! Largest Region is", max)
    fmt.Println("But, on the other hand...... Total Region is", regn)
}

func findNearest (x, y int, xAll, yAll []int) int {
    min := 10000
    var res int
    for i:=0; i < len(xAll); i++ {
        dist := Abs(x - xAll[i]) + Abs(y - yAll[i])
        if dist < min {
            min = dist
            res = i
            continue
        }
        if dist == min { res = -1 }
    }
    return res
}

func findTotal (x, y int, xAll, yAll []int) int {
    var total int
    for i:=0; i < len(xAll); i++ {
        dist := Abs(x - xAll[i]) + Abs(y - yAll[i])
        total = total + dist
    }
    return total
}

func Abs (i int) int {
    if i < 0 {
        return -i
    } else {
        return i
    }
}
