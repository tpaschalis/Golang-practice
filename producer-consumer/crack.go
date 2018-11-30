package main

import (
    "fmt"
    "os"
    "bytes"
    "unsafe"
    "sync"
    _"runtime"
)


// #cgo LDFLAGS: -lcrypt
// #define _GNU_SOURCE
// #include <crypt.h>
// #include <stdlib.h>
import "C"

// crypt wraps C library crypt_r
func crypt(key, salt string) string {
    data := C.struct_crypt_data{}
    ckey := C.CString(key)
    csalt := C.CString(salt)
    out := C.GoString(C.crypt_r(ckey, csalt, &data))
    C.free(unsafe.Pointer(ckey))
    C.free(unsafe.Pointer(csalt))
    return out
}

func produce(in chan string) {
    defer close(in)
//    for i:=0; i<1000; i++ {
//        in <- fmt.Sprintf("%d", i)
//    }

    var buffer bytes.Buffer
    for a:=0; a<len(possibleLetters); a++ {
        for b:=0; b<len(possibleLetters); b++ {
            for c:= 0; c<len(possibleLetters); c++ {
                for d:=0; d<len(possibleLetters); d++ {
                    for e:=0; e<len(possibleLetters); e++ {
                        buffer = bytes.Buffer{}
                        buffer.WriteString(possibleLetters[a])
                        buffer.WriteString(possibleLetters[b])
                        buffer.WriteString(possibleLetters[c])
                        buffer.WriteString(possibleLetters[d])
                        buffer.WriteString(possibleLetters[e])

                        in <- buffer.String()
                    }
                }
            }
        }
    }
}

func consume(in, out chan string, wg *sync.WaitGroup, salt string) {
    defer wg.Done()
    for s:= range in {
        currentHash := crypt(s, salt)
        if currentHash == hashToCrack {
            fmt.Println(currentHash,"\n")
            out <- s
        }
    }
}

func stop(out chan string, wg *sync.WaitGroup){
    wg.Wait()
    close(out)
}


var possibleLetters = [57]string{ "","a","A","b","B","c","C","d","D","e","E","f","F","g","G","h","H","i","I","j","J","k","K","l","L","m","M","n","N","o","O","p","P","q","Q","r","R","s","S","t","T","u","U","v","V","w","W","x","X","y","Y","z","Z"}

var hashToCrack string

func main() {
    numOfArgs := len(os.Args)
    if numOfArgs != 2 {
        fmt.Println("Sorry, code requires exactly one argument")
        os.Exit(1)
    }
    hashToCrack = os.Args[1]
    firstTwo := hashToCrack[0:2]
    fmt.Println(hashToCrack)


    in, out := make(chan string), make(chan string)
    wg := &sync.WaitGroup{}
    go produce(in)

    //for i:=0; i<runtime.NumCPU();i++ {
    for i:=0; i<20; i++ {
        wg.Add(1)
        go consume(in, out, wg, firstTwo)
    }
    go stop (out, wg)
    fmt.Println(<-out)


}

// List of possible improvements
//
// Check if recursion would be faster
// Benchmark usage of buffer.WriteString, or maybe use
// some other utility, such as StringBuilder
//

