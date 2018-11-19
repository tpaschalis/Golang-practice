package main

import (
    "fmt"
    "golang.org/x/image/bmp"
    _"io"
    "os"
    "reflect"
    "image"
    "image/color"
    "image/png"
)

func main() {

    if len(os.Args) != 2 {
        fmt.Println("Expecting exactly one argument.")
        fmt.Println("Usage : go run main.go clue.bmp")
        os.Exit(1)
    }

    filename := os.Args[1]
    file, err := os.Open(filename)
    if err != nil {
        panic(err)
    }
    fmt.Println(file)
    img, err := bmp.Decode(file)
    if err != nil {
        panic(err)
    }
    fmt.Println(img)
    fmt.Println(reflect.TypeOf(img),"\n")

    size := img.Bounds().Size()
    m := image.NewRGBA(image.Rect(0, 0, size.X, size.Y))
	for x := 0; x < size.X; x++ {
	    for y := 0; y < size.Y; y++ {
            r,g,b,a := img.At(x,y).RGBA()
            _ = a
	        color := color.RGBA{
                uint8(r/255),
                uint8(g/255),
                uint8(b/255),
                255}
	        m.Set(x, y, color)
	    }
	}
    f, _ := os.OpenFile("result.png", os.O_WRONLY|os.O_CREATE, 0600)
    defer f.Close()
    png.Encode(f, m)
}


// IT WAS 1) PROFESSOR PLUM 2) WITH THE CANDLESTICK 3) IN THE LIBRARY
//
// List of possible improvements
