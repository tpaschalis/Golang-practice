package main

import "fmt"

type rcave struct {
    x int
    y int
    erosion int
    ground int
}

// Part 2
type Node struct {
    x, y     int
    gear     int
    previous *Node
}
var graph[10+200][715+200][3]int
const (
    NAKED int = iota
    TORCH
    PICKAXE
)
var directions = [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func main() {
    //targetX, targetY := 10, 10
    //depth := 510

    targetX, targetY := 10, 715
    depth := 3339

    cave := make([][]rcave, targetX+200)
    for i:= range cave {
        cave[i] = make([]rcave, targetY+200)
    }

    // Initialize the cave and sides (first two rules)
    for x := range cave {
        for y:= range cave[x] {
            cave[x][y].x = x
            cave[x][y].y = y
            if x==0 && y==0 {
                cave[x][y].erosion = (0 + depth) % 20183
                cave[x][y].ground = cave[x][y].erosion %3
            }
            if y==0 && x!=0{
                cave[x][y].erosion = (x*16807 + depth) % 20183
                cave[x][y].ground = cave[x][y].erosion %3
            }
                if x==0 && y!=0 {
                    cave[x][y].erosion = (y*48271 + depth) % 20183
                    cave[x][y].ground = cave[x][y].erosion %3
            }
        }
    }

    // Filling in the cave (third rule)
    for x:=1; x<=targetX; x++ {
        for y:=1; y<=targetY; y++ {
            if x == targetX && y == targetY { break }
            cave[x][y].erosion = (cave[x-1][y].erosion * cave[x][y-1].erosion + depth) % 20183
            cave[x][y].ground = cave[x][y].erosion %3
        }
    }

    // Target coordinates (fourth rule)
    cave[targetX][targetY].erosion = (0+depth) % 20183
    cave[targetX][targetY].ground = cave[targetX][targetY].erosion %3

    // Part Two
    // Modifying Dijkstra for each step seems difficult.
    // Maybe modifying the graph is easier?




    queue := []*Node {
        &Node{0, 0, TORCH, nil},
    }
    graph[0][0][TORCH] = 1

    for len(queue) > 0 {
        // Find the minimum weight
        min := -1
        for _, entry := range queue {
            if min < 0 || entry.Weight() < min {
                min = entry.Weight()
            }
        }

        minCount := 0
        for i, entry := range queue {
            if entry.Weight() == min {
                if entry.x == targetX && entry.y == targetY && entry.gear == TORCH {
                    fmt.Println("Shortest path :", entry.Weight()-1)
                    return
                }
                for _, dir := range directions {
                    x := entry.x + dir[0]
                    y := entry.y + dir[1]
                    //fmt.Println(x,y)
                    if x < 0 || y < 0 {
                        continue
                    } else if x>=len(cave) || y>=len(cave[x]) {
                        continue
                    }
                    if entry.gear != cave[x][y].ground {
                        posWeight := &graph[x][y][entry.gear]
                        if *posWeight == 0 || *posWeight > entry.Weight() + 1 {
                            *posWeight = entry.Weight() + 1
                            queue = append(queue, &Node{x, y, entry.gear, entry})
                        }
                    }
                }
                //fmt.Println(entry.x, entry.y)
                gear := 3 ^ (entry.gear ^ cave[entry.x][entry.y].ground)
                //fmt.Println(entry.x, entry.y, gear)
                posWeight := &graph[entry.x][entry.y][gear]
                if *posWeight == 0 || *posWeight > entry.Weight() + 7 {
                    *posWeight = entry.Weight() + 7
                    queue = append(queue, &Node{entry.x, entry.y, entry.gear, entry})
                }
                // Move processed entries to the beginning, to remove them easily
                queue[minCount], queue[i] = queue[i], queue[minCount]
                minCount++
            }
        }
        queue = queue[minCount:]
    }







    fmt.Println("I'm out!")





}

func (e *Node) Weight() int {
    return graph[e.x][e.y][e.gear]
}



func display(c [][]rcave){

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

