package main

import (
	"fmt"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

type Obj struct {
	rx, ry, px, py int
}

func main() {

	flares := []Obj{}
	var err error

	for err == nil {
		c := Obj{}
		_, err = fmt.Scanf("position=<%d, %d> velocity=<%d,  %d>\n", &c.rx, &c.ry, &c.px, &c.py)
		flares = append(flares, c)
	}
	flares = flares[:len(flares)-1]     // The final one is empty, out of the way!

	message := make([]Obj, len(flares))
	min, t0 := 1000000, 0
	for t := 1; t < 20000; t++ {
		timestep(flares)
		if predictMessage(flares) < min {
			min, t0 = predictMessage(flares), t
			copy(message, flares)
		}
	}


	pl, err := plot.New()
	if err != nil {
		panic(err)
	}

	pts := make(plotter.XYs, len(flares))
	for i := range message {
		pts[i].X = float64(message[i].rx)
		pts[i].Y = float64(message[i].ry)
	}
	s, err := plotter.NewScatter(pts)

	pl.Add(s)
	if err := pl.Save(30*vg.Inch, 30*vg.Inch, "points.png"); err != nil {
		panic(err)
	}
    fmt.Println("Open `points.png` using mspaint, eog, or whatever you wish to see the letters...")
	fmt.Println("After so many seconds, the best score I saw was on second :", t0)
}

func timestep(f []Obj) {
    // One second passes, and our flares move!
	for i := range f {
		f[i].rx += f[i].px
		f[i].ry += f[i].py
	}
	return
}

func minmax(f []Obj) (int, int, int, int) {
	xmax, xmin, ymax, ymin := 0, 0, 0, 0
	for i := range f {
		if f[i].rx > xmax { xmax = f[i].rx }
		if f[i].rx < xmin {	xmin = f[i].rx }
		if f[i].ry > ymax {	ymax = f[i].ry }
		if f[i].ry < ymin {	ymin = f[i].ry }
	}
	return xmax, xmin, ymax, ymin
}

func predictMessage(f []Obj) int {
	// Find when the letters take the least area
	x2, x1, y2, y1 := minmax(f)
	return (x2 - x1) * (y2 - y1)
}
