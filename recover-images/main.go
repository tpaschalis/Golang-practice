package main

import (
    "fmt"
    "io/ioutil"
    "strconv"
)


func main() {
    data, err := ioutil.ReadFile("card.raw")
    if err != nil {
        panic(err)
    }

    bufferSize := 512
    j := 0

    var jpgStartPoints []int
    for i:=0; i<len(data)/bufferSize; i++ {
        currentBuffer := data[i*bufferSize:(i+1)*bufferSize]
        if currentBuffer[0]==255 && currentBuffer[1]==216 && currentBuffer[2]==255 {
            j = j + 1
            jpgStartPoints = append(jpgStartPoints, i)
        }
    }

    // Close point for last picture
    jpgStartPoints = append(jpgStartPoints,len(data)-1)

    for i:=0; i<len(jpgStartPoints)-2; i++ {

        fmt.Println(strconv.Itoa(i))
        filename := "/tmp/dat"+strconv.Itoa(i)+".jpg"
        extractedImage := data[jpgStartPoints[i]*bufferSize : jpgStartPoints[i+1]*bufferSize-1]
        err := ioutil.WriteFile(filename, extractedImage, 0644)
        if err != nil {
            panic(err)
        }

    }

    fmt.Println("Images Detected :", j)
    fmt.Println(len(jpgStartPoints))
}

// List of possible Improvements
// 
// Am i missing the final picture?
// Is there a smarter way to read and detect the startpoints?
// Is there a smarter way to denote the space between the startpoints?
