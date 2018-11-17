package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "strconv"
)

func main() {

    // Read pyramid height
	reader := bufio.NewReader(os.Stdin)
    height := -1
    for {
	    fmt.Printf("Height : ")
	    input, err := reader.ReadString('\n')
	    input = strings.TrimSuffix(input, "\n")
	    if err != nil {
		    panic(err)
	    }
        height, err = strconv.Atoi(input)
        if err != nil {
            panic(err)
        }
        if height > 0 {
            break
        }
    }

    spaceString := " "
    blockString := "#"

    for i:=1; i<=height; i++ {
        spaces := strings.Repeat(spaceString, height-i)
        blocks := strings.Repeat(blockString, i)

        fmt.Println(spaces + blocks + spaceString + blocks + spaces)
    }
}
