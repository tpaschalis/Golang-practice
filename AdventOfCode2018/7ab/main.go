package main

import (
	"fmt"
	"sort"
)

var Alphabet string

func main() {
	var err error
	var order []string
	// Data are pairs of 'steps' and 'prerequisites'
	// links is used for pt1, and coop for pt2
	link := make(map[string][]string)
	coop := make(map[string][]string)
	Alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	// Initialize the map, because the first available steps will not have a 'prerequisite'
	for i := range Alphabet {
		link[Alphabet[i:i+1]] = []string{}
		coop[Alphabet[i:i+1]] = []string{}
	}

	for err == nil {
		var s, p string
		_, err = fmt.Scanf("Step %s must be finished before step %s can begin.\n", &p, &s)
		link[s] = append(link[s], p)
		coop[s] = append(coop[s], p)
	}

	fmt.Println(Ikea(link, order))

	elves := make(map[string]int)
	fmt.Println(factory(coop, elves, 0))
}

func Ikea(d map[string][]string, ord []string) []string {
	cand := []string{}
	if len(d) == 1 { return ord	}

	// Which steps are candidates for execution?
	for k, v := range d {
		if len(v) == 0 {
			cand = append(cand, k)
		}
	}

	// Pick the first one alphabetically
	sort.Strings(cand)
	currStep := cand[0]
	ord = append(ord, currStep)
	delete(d, currStep)

	// Then, execute it!
	for k, v := range d {
		d[k] = remove(v, currStep)
	}
	return Ikea(d, ord)
}

func factory(d map[string][]string, workers map[string]int, t int) int {
	cand := []string{}
	if len(d) == 1 && len(workers) == 0 {
		return t // We're all done!
	}

	// Which steps can the workers work on?
	for k, v := range d {
		if len(v) == 0 {
			cand = append(cand, k)
		}
	}
	sort.Strings(cand)

	// Are there any workers that will finish on the next second?
	for j, u := range workers {
		workers[j]--
		if u == 1 { // Step will finish on the next second
			cand = remove(cand, j)
			delete(workers, j)
			delete(d, j)
			for k, _ := range d {
				d[k] = remove(d[k], j)
			}
		}
	}

	// Are there empty workers to assign a job to?
	for i := 0; i < 5-len(workers); i++ {
		if i < len(cand) {
			currStep := cand[i]
			if workers[currStep] == 0 {
				workers[currStep] = timeReq(currStep)
				delete(d, currStep)
			}
		}
	}

	t = t + 1
	return factory(d, workers, t)
}

func remove(slice []string, s string) []string {
	for i, v := range slice {
		if v == s {	return append(slice[:i], slice[i+1:]...) }
	}
	return slice
}

func timeReq(s string) int {
	// Time needed to assemble a step
	var Alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for i := range Alphabet {
		if Alphabet[i:i+1] == s { return i + 60	}
	}
	return -1
}
