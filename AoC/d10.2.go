package AoC

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func D10_2() {
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
	var completions []string
	for _, line := range data {
		fmt.Printf("%s\n", line)
		var stack []string
		var last string
		var illegal bool = false
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
				illegal = true
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
		if !illegal {
			// score completion tokens.
			// for each right character from stack
			for i := len(stack) - 1; i >= 0; {

				var current = stack[i]
				// right or left
				if isLeft(current) {
					// add right to completions
					completions = append(completions, syntax[current])
					// step left
					i--
				}
				// if left we add to incomplete
				var match string = ""
				// for right char find left
				for j := i; j < 0; j-- {
					// if next char is right, find left  then hand back to i
					var next = stack[j]
					if isRight(next) {
						match = getLeft(next)
					} else {
						// we have a left so check for match
						if match != "" && next == match {
							match = ""
						}
					}
					i--
				}
			}
		}
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
func getLeft(right string) string {
	for k, v := range syntax {
		if v == right {
			return k
		}
	}
	return "error"
}
