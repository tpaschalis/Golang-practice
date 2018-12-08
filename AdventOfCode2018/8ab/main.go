package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var Metadata []int

type node struct {
	children []*node
	metadata []int
}

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
	root, _ := parse(tree, Metadata)
	fmt.Println(sumMetadata(Metadata))
	fmt.Println(root.value())

}

func parse(data []int, metadata []int) (*node, int) {
	idx := 2
	n := &node{}

	for i := data[0]; i > 0; i-- {
		child, addMore := parse(data[idx:], metadata)
		idx += addMore
		n.children = append(n.children, child)
	}
	currMd := data[idx : idx+data[1]]
	n.metadata = data[idx : idx+data[1]]
	for i := range currMd {
		Metadata = append(Metadata, currMd[i])
	}
	idx += data[1]
	return n, idx
}

// For this part I cheated a bit :(
// took a hint from github.com/Alex-Wauters
func (n *node) value() (s int) {
	if len(n.children) == 0 {
		for _, meta := range n.metadata {
			s += meta
		}
		return s
	}
	for _, meta := range n.metadata {
		if meta > len(n.children) {
			s += 0
		} else {
			s += n.children[meta-1].value()
		}
	}
	return s
}

func sumMetadata(in []int) int {
	res := 0
	for i := range in { res += in[i] }
	return res
}
