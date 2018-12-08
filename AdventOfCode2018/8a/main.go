package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "strconv"
)

var Metadata []int

func main() {
    var input string
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        input = scanner.Text() // one line of input
    }
    in := strings.Split(input, " ")

    var tree []int
    for i := range in {
        curr, _ := strconv.Atoi(in[i])
        tree = append(tree, curr)
    }
    _ = parse(tree, Metadata)
    fmt.Println(sumMetadata(Metadata))
}

func parse(data []int, metadata []int) (int) {
    idx := 2
    for i:= data[0]; i>0; i-- {
        addMore := parse(data[idx:], metadata)
        idx += addMore
    }
    currMd := data[idx:idx+data[1]]
    for i := range currMd {
        Metadata = append(Metadata, currMd[i])
    }
    idx += data[1]
    return idx
}

func sumMetadata(in []int) int {
    res := 0
    for i:= range in {
        res += in[i]
    }
    return res
}
