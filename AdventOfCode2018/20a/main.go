package main

import "fmt"
import "strings"

var list []string

func main() {

    //input := []string{`^ENWWW(NEEE|SSE)NN$`}
    //input := []string{`^ENWWW(NEEE|SSE(EE|N)))ASAASSA(MY(AA|BB)QWERTY(CC|DD(EE|FF|GG))STUFF)PPP$`}
    in := []string{`^ENWWW(NEEE|SSE(EE|N))$`}

    list = append(list, in[0])
    branch(in, list)

    fmt.Println("-----------------------------------------------------------------------------------------------")
    fmt.Println(in)
    fmt.Println(list, len(list))
    for k, v := range list {
        fmt.Println(k, v)
    }
}

func branch(input []string, l []string) {
    fmt.Println("branchiiing!")
    fmt.Println(input)
    for s := range input{
        current := input[s]
        fmt.Println(current)
        for i:=0; i<len(current); i++ {
            if current[i:i+1] == "(" {
                //branch([]string{current[i+1 : len(current)]})
                branch([]string{current[i+1 : len(current)]}, l)
                break
            }
        }

        current = "("+current
        currentParen := getParen(current)
        parts := strings.Split(currentParen, "|")
        fmt.Println("Capturing : ", getParen(current), "w/ parts", parts)
        var currentSentence string
        for q := range l {
            if strings.Contains(l[q], getParen(current)) {
            //list = remove(list, current)
                currentSentence = l[q]
                fmt.Println("Contained, ", l[q], getParen(current))
                l = remove(l, l[q])
                for p:= range parts {
                    e := strings.Replace(currentSentence, getParen(currentSentence), parts[p], -1)
                    l = append(l, e)
                }
            }
        }
        //for p := range parts {
        //
        //}

        fmt.Println()
    }
}

func getParen(input string) string {
    lp := 0
    for i:=0; i<len(input); i++ {
        if input[i:i+1] == "(" {lp++}
        if input[i:i+1] == ")" {lp--}
        if lp == 0 {
            return (input[0:i+1])
        }
    }
    //panic("no paren matched")??
    return input
}

func remove(slice []string, s string) []string {
    for i, v := range slice {
        if v == s { return append(slice[:i], slice[i+1:]...) }
    }
    return slice
}
