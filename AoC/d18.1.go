package AoC

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type TreeNode struct {
	left  *TreeNode
	right *TreeNode
	value int
}
type SyntaxTree struct {
	top   *TreeNode
	depth int // keep track to save traverse to derive
}

func (n TreeNode) AddNode(value string, left bool) {

	// check if we are adding Leaf
	val, err := strconv.ParseInt(value, 10, 64)
	if err == nil {
		fmt.Printf("create leaf with %d\n", val)
		n.value = int(val)
	} else {
		// not leaf so get parts
		leftPart, rightPart := getParts(value)
		// if there is a part then add new Node to hold it.
		if leftPart != "" {
			n.left = &TreeNode{}
			n.left.AddNode(leftPart, true)
		}
		if rightPart != "" {
			n.right = &TreeNode{}
			n.right.AddNode(leftPart, false)
		}
	}

}

func (t SyntaxTree) buildFromString(line string) {

	t.top = &TreeNode{}
	left, right := getParts(line)
	t.top.AddNode(string(left), true)
	t.top.AddNode(string(right), false)
}

func getParts(expression string) (left string, right string) {
	left = ""
	right = ""

	expression = expression[1 : len(expression)-1]
	// strips current level brackets
	fmt.Printf("Stripped outer brackets for top of tree %s\n", expression)
	start := string(expression[0])
	end := string(expression[len(expression)-1])
	fmt.Printf("start, end: %s, %s\n", start, end)
	_, startErr := strconv.ParseInt(start, 10, 32)
	_, endErr := strconv.ParseInt(end, 10, 32)
	fmt.Printf("parsing ends %o, %o\n", startErr, endErr)
	if startErr == nil {
		// start is a number as left  and everything else is right
		left = start
		right = string(expression[2:len(expression)])
		fmt.Printf("start is a number %s, %d\n", start, right)
	} else if endErr == nil {
		// end is a number and everything to left is left
		left = string(expression[0 : len(expression)-2])
		fmt.Printf("end is a number %s, %d\n", left, end)
		right = end
	} else {
		// we need to find the middle and split
	}
	return left, right

}

func D18_1() {
	dat, err := ioutil.ReadFile("inputd18")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	var data = strings.Split(string(dat), "\n")
	var tree SyntaxTree = SyntaxTree{}
	for _, line := range data {
		fmt.Printf("%s\n", line)

		fmt.Printf("Tree:%o\n", tree)
		tree.buildFromString(line)
		fmt.Printf("Tree:%o\n", tree)

	}

}
