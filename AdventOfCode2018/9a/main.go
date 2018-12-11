package main

import (
    "fmt"
)


var elves []int

func main() {

    for i:=0; i <=464; i++ {
        elves = append(elves, 0)
    }

    circle := []int {0, 2, 1, 3}
    fmt.Println(circle)


    currentMarble := 3
    nextMarble := 4
    player := 2
    for i:=4; i < 7173000; i++ {
        circle, nextMarble, currentMarble, player = nextstep(circle, nextMarble, currentMarble, player)
    }
    fmt.Println(Max(elves))

}

func insert(slice []int, pos, val int) []int {
    s := []int{}
    s = append(s, val)
    slice = append(slice[:pos], append(s, slice[pos:]...)...)
    return slice
}

func remove(slice []int, pos int) []int {
    res := append(slice[:pos], slice[pos+1:]...)
    return res
}

func nextstep(circle []int, nm, cur, p int) ([]int, int, int, int) {
    if (nm)%23 == 0 {
        bc := backcur(circle, cur)
        elves[p] += nm
        elves[p] += circle[bc]
        circle = remove(circle,bc)
        cur = bc
        nm++
        //elves[p] += 
    } else {
        cur = (cur + 2) % len(circle)
        circle = insert(circle, cur, nm)
        nm++
    }
    p = (p+1)%464
    return circle, nm, cur, p
}

func backcur(circle []int, i int) int {
    if i - 7 >= 0 {
        return i - 7
    } else {
        return (len(circle)+(i-7))
    }
}

func Max(s []int) int {
    max := 0
    for i:= range s {
        if s[i] > max {
            max = s[i]
        }
    }
    return max
}
