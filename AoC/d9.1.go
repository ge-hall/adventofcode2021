package AoC

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func D9_1() {
	dat, err := ioutil.ReadFile("inputd9.1")
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

	var lowPoints []int64
	for i := range heightMap {
		for j := range heightMap[i] {
			fmt.Printf("%d", heightMap[i][j])
			if check(heightMap, i, j) {
				lowPoints = append(lowPoints, heightMap[i][j])
			}
		}
		print("\n")
	}

	print("\n")

	fmt.Printf("%o\n", lowPoints)

	var result int = 0
	for _, v := range lowPoints {
		result += int(v) + 1
	}
	fmt.Printf("Result: %d\n", result)
}
func check(hm [][]int64, i int, j int) bool {
	var cur int64 = hm[i][j]
	// above
	if (i > 0 && hm[i-1][j] <= cur) ||
		// left
		(j > 0 && hm[i][j-1] <= cur) ||
		// below
		(i < len(hm)-1 && hm[i+1][j] <= cur) ||
		// right
		(j < len(hm[i])-1 && hm[i][j+1] <= cur) {
		return false
	}
	return true
}
