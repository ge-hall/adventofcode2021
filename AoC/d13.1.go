package AoC

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func D13_1() {
	dat, err := ioutil.ReadFile("inputd13")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	var data = strings.Split(string(dat), "\n")
	//var paper [][]int
	var plotting = true
	var x []int
	xMax := 0
	var y []int
	yMax := 0
	for _, line := range data {
		fmt.Printf("%s\n", line)
		if line == "" {
			print("empty line\n")
			plotting = false
			continue
		}
		if !plotting {
			// read instructions
			continue
		}

		coord := strings.Split(line, ",")
		xv, _ := strconv.ParseInt(coord[0], 10, 32)
		if int(xv) > xMax {
			xMax = int(xv)
		}
		x = append(x, int(xv))

		yv, _ := strconv.ParseInt(coord[1], 10, 32)
		if int(yv) > yMax {
			yMax = int(yv)
		}
		y = append(y, int(yv))

	}

	fmt.Printf("%v max = %d\n", x, xMax)
	fmt.Printf("%v max = %d\n", y, yMax)
	print("===============\n\n")
	print("===============\n\n")
	print(pathCount)
}
