package AoC

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

var octopuses [][]int

func D11_1() {
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
	flashCount := 0
	steps := 100
	for i := 0; i < steps; i++ {
		flashCount += runStep()
	}
	//fmt.Printf("%o\n", octopuses)

	print("\n")
	print("\n")

	var result int = flashCount

	fmt.Printf("Result: %d\n", result)
}
func runStep() int {
	fmt.Printf("Start step with len %d\n", len(octopuses))

	printCave()
	// increment energy
	for r := 0; r < len(octopuses); r++ {
		for c := 0; c < len(octopuses[0]); c++ {
			octopuses[r][c]++
		}
	}
	print("\n")
	// check for flashes
	printCave()
	totalFlash := 0
	fCount := 1
	for fCount > 0 {
		fCount = 0
		for r := 0; r < len(octopuses); r++ {
			for c := 0; c < len(octopuses[0]); c++ {
				if octopuses[r][c] > 9 {
					fCount++
					flash(r, c)
				}

			}
		}
		totalFlash += fCount
		printCave()
	}
	return totalFlash
}
func flash(r int, c int) {
	octopuses[r][c] = 0
	rCount := len(octopuses)
	cCount := len(octopuses[0])
	// n
	if r > 0 && octopuses[r-1][c] != 0 {
		octopuses[r-1][c]++

	}
	//ne

	if r > 0 && c < cCount-1 && octopuses[r-1][c+1] != 0 {
		octopuses[r-1][c+1]++

	}
	//e
	if c < cCount-1 && octopuses[r][c+1] != 0 {
		octopuses[r][c+1]++

	}
	//se
	if r < rCount-1 && c < cCount-1 && octopuses[r+1][c+1] != 0 {
		octopuses[r+1][c+1]++

	}
	//s
	if r < rCount-1 && octopuses[r+1][c] != 0 {
		octopuses[r+1][c]++

	}
	//sw
	if r < rCount-1 && c > 0 && octopuses[r+1][c-1] != 0 {
		octopuses[r+1][c-1]++

	}
	//w
	if c > 0 && octopuses[r][c-1] != 0 {
		octopuses[r][c-1]++

	}
	//nw
	if r > 0 && c > 0 && octopuses[r-1][c-1] != 0 {
		octopuses[r-1][c-1]++

	}
}
func printCave() {
	for r := 0; r < len(octopuses); r++ {
		for c := 0; c < len(octopuses[0]); c++ {
			fmt.Printf("%d, ", octopuses[r][c])
		}
		print("\n")
	}
	print("\n")
}
