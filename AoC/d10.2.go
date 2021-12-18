package AoC

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

var completions [][]string

func D10_2() {
	dat, err := ioutil.ReadFile("inputd10.1")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	var data = strings.Split(string(dat), "\n")
	var tokenScore = map[string]int{
		")": 1,
		"]": 2,
		"}": 3,
		">": 4,
	}
	var illegals []string

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
			nextIllegal(line)

		}
	}
	print("\n")
	fmt.Printf(" c%o\n", completions)

	var result int = 0

	// cal syntax completion scores
	// store and sort then take mid value
	var scores []int
	for _, c := range completions {
		var cresult int
		for _, t := range c {
			cresult = cresult*5 + tokenScore[t]
			//fmt.Printf("%d\n", result)
		}
		scores = append(scores, cresult)
	}
	sort.Ints(scores)
	result = scores[len(scores)/2]
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

func nextIllegal(line string) {
	var stack []string
	var last string
	var thisCompletions []string
	//parseLine
	for i := len(line) - 1; i >= 0; i-- {
		if len(stack) == 0 {
			last = string(line[i])
			if isRight(last) {
				stack = append(stack, last)
			} else {
				thisCompletions = append(thisCompletions, syntax[string(line[i])])
			}
			continue
		}
		// if illegal char

		// next char is illegal if
		// last is right and c is left and not matched
		if isRight(last) && isLeft(string(line[i])) && !isSyntaxPair(string(line[i]), last) {
			thisCompletions = append(thisCompletions, string(line[i]))
			fmt.Printf("Expected %s, but found %s\n", syntax[last], string(line[i]))
			// clear stack and keep going
			stack = nil
		}

		fmt.Printf("checking last %s, with %s\n", last, string(line[i]))
		// iff legal and right then push
		if isRight(string(line[i])) {

			last = string(line[i])
			stack = append(stack, string(line[i]))
			fmt.Printf("pushing c %s on to stack\n", string(line[i]))
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
	completions = append(completions, thisCompletions)
}
