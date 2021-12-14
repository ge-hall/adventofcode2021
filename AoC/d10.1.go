package AoC

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func D10_1() {
	dat, err := ioutil.ReadFile("inputd10.1")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	var data = strings.Split(string(dat), "\n")
	heightMap := make([][]int64, len(data))
	for i := range heightMap {
		heightMap[i] = make([]int64, 100)
	}
	for i, line := range data {
		fmt.Printf("%s\n", line)
		for j := range line {
			heightMap[i][j], _ = strconv.ParseInt(string(line[j]), 10, 64)
		}
	}
	print("\n")

	print("\n")

	var result int = 0

	fmt.Printf("Result: %d\n", result)
}
