package main

import "fmt"

type Link struct {
	i    int
	p, l *Link
}

// Initially I only completed Part One.
// This Part Two is completely and shamelessly stolen from this guy :
// https://github.com/seankhliao
// Check his stuff out, he's awesome in Go

func main() {
	players := make([]int, 464)
	n := 0
	p := &Link{}

	p.p, p.l = p, p
	pi, pl := p, p

	for n <= 7173000 {
		n++
		if n%23 == 0 {
			p = p.p.p.p.p.p.p.p
			pi, pl = p.p, p.l
			players[n%len(players)] += n + p.i
			pi.l, pl.p = pl, pi
			p = pl
		} else {
			p = p.l
			pi, pl = p, p.l
			p = &Link{i: n}
			pi.l, p.l = p, pl
			pl.p, p.p = p, pi
		}
	}
	highScore := 0
	for _, p := range players {
		if p > highScore { highScore = p }
	}
	fmt.Println("High Score :", highScore)
}
