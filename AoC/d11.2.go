package AoC

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func D11_2() {
	dat, err := ioutil.ReadFile("inputd11")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	var data = strings.Split(string(dat), "\n")

	for _, line := range data {
		//fmt.Printf("%s\n", line)
		var row []int
		for i := 0; i < len(line); i++ {
			v, _ := strconv.ParseInt(string(line[i]), 10, strconv.IntSize)
			row = append(row, int(v))

		}
		octopuses = append(octopuses, row)
	}
	breakStep := 0
	steps := 2000
	for i := 0; i < steps; i++ {
		runStep()
		if flashSync() {
			fmt.Printf("breaking on step %d\n", i)
			breakStep = i + 1
			printCave()
			break
		}
		//fmt.Printf("After step %d\n", i+1)
		//printCave()
	}
	//fmt.Printf("%o\n", octopuses)

	print("\n")
	print("\n")

	var result int = breakStep

	fmt.Printf("Result: %d\n", result)
}
func flashSync() bool {
	for r := 0; r < len(octopuses); r++ {
		for c := 0; c < len(octopuses[0]); c++ {
			if octopuses[r][c] != 0 {
				return false
			}
		}
	}
	return true
}
