package AoC

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

var syntax = map[string]string{
	"[": "]",
	"(": ")",
	"{": "}",
	"<": ">",
}

func D10_1() {
	dat, err := ioutil.ReadFile("inputd10.1")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	var data = strings.Split(string(dat), "\n")
	var tokenScore = map[string]int{
		")": 3,
		"]": 57,
		"}": 1197,
		">": 25137,
	}
	var illegals []string

	for _, line := range data {
		fmt.Printf("%s\n", line)
		var stack []string
		var last string
		//parseLine
		for _, c := range line {
			if len(stack) == 0 {
				last = string(c)
				stack = append(stack, string(c))
				continue
			}
			// if illegal char

			// next char is illegal if
			// last is left and c is right and not matched
			if isLeft(last) && isRight(string(c)) && !isSyntaxPair(last, string(c)) {
				illegals = append(illegals, string(c))
				fmt.Printf("Expected %s, but found %s\n", syntax[last], string(c))
				break
			}
			fmt.Printf("checking last %s, with %s\n", last, string(c))
			// iff legal and left then push
			if isLeft(string(c)) {

				last = string(c)
				stack = append(stack, string(c))
				fmt.Printf("pushing c %s on to stack\n", string(c))
			} else {
				// must be right and legal so pop and set last to new top of stack
				var popped = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if len(stack) > 0 {
					last = stack[len(stack)-1]
				}
				fmt.Printf("popped  %s off stack\n", popped)
				fmt.Printf("stack = %o\n", stack)
			}

		}
		print("============\n")
	}
	print("\n")
	print("\n")

	var result int = 0

	// cal syntax error score
	for _, t := range illegals {
		result += tokenScore[t]
	}
	fmt.Printf("Result: %d\n", result)
}
func isSyntaxPair(l string, r string) bool {
	if val, ok := syntax[l]; ok {
		if val == r {
			return true
		}
	}
	return false
}
func isLeft(t string) bool {
	for k, _ := range syntax {
		if k == t {
			return true
		}
	}
	return false
}
func isRight(t string) bool {
	for _, v := range syntax {
		if v == t {
			return true
		}
	}
	return false
}
