package AoC

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

func D9_2() {
	dat, err := ioutil.ReadFile("inputd9.1")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	var data = strings.Split(string(dat), "\n")
	heightMap := make([][]int64, len(data))
	checkMap := make([][]int64, len(data))
	for i := range heightMap {
		heightMap[i] = make([]int64, 100)
		checkMap[i] = make([]int64, 100)
	}
	for i, line := range data {
		fmt.Printf("%s\n", line)
		for j := range line {
			heightMap[i][j], _ = strconv.ParseInt(string(line[j]), 10, 64)
		}
	}
	print("\n")

	var lowPoints []int64
	var basins []int
	for i := range heightMap {
		for j := range heightMap[i] {
			//fmt.Printf("%d", heightMap[i][j])
			if check(heightMap, i, j) {
				lowPoints = append(lowPoints, heightMap[i][j])

				basins = append(basins, countNeighbours(heightMap, checkMap, i, j))
			}
		}
		print("\n")
	}

	print("\n")
	fmt.Printf("%o\n", basins)

	sort.Ints(basins)
	l := len(basins)
	var result int = (basins[l-1]) * (basins[l-2]) * (basins[l-3])
	fmt.Printf("Result: %d\n", result)
}
func countNeighbours(hm [][]int64, cm [][]int64, i int, j int) int {
	fmt.Printf("count: %d, %d\n", i, j)

	if i < 0 || j < 0 || i >= len(hm) || j >= len(hm[0]) || cm[i][j] == -1 || hm[i][j] == 9 {
		return 0
	}
	cm[i][j] = -1
	fmt.Printf("val= %d count: %d, %d\n", hm[i][j], i, j)
	return 1 + countNeighbours(hm, cm, i-1, j) +
		countNeighbours(hm, cm, i+1, j) +
		countNeighbours(hm, cm, i, j-1) +
		countNeighbours(hm, cm, i, j+1)

}
func basinSize(hm [][]int64, i int, j int) int {
	//var cur int64 = hm[i][j]
	fmt.Printf("hm[][] %d, %d\n", i, j)
	var size = 1
	// above

	for x := i - 1; x >= 0; x-- {
		fmt.Printf("above basin %d\n", hm[x][j])
		if hm[x][j] != 9 {
			size++
		} else {
			break
		}
	}
	// left
	for x := j - 1; x >= 0; x-- {
		fmt.Printf("left basin %d\n", hm[i][x])
		if hm[i][x] != 9 {
			size++
		} else {
			break
		}
	}
	// below
	for x := i + 1; x < len(hm); x++ {
		fmt.Printf("below basin %d\n", hm[x][j])
		if hm[x][j] != 9 {
			size++
		} else {
			break
		}
	}

	// right
	for x := j + 1; x < len(hm[i]); x++ {
		fmt.Printf("right basin %d\n", hm[i][x])
		if hm[i][x] != 9 {
			size++
		} else {
			break
		}
	}
	fmt.Printf("======== %d\n", size)
	return size
}
