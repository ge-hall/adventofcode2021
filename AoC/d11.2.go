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
