package AoC

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type Vector struct {
	x1 int64
	y1 int64
	x2 int64
	y2 int64
}

func D5_1() {
	dat, err := ioutil.ReadFile("inputd5.1")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	var lines = strings.Split(string(dat), "\n")

	// read vectors
	var vectors []Vector
	var maxX, maxY int64
	var i int
	for i = 0; i < len(lines); i++ {
		fmt.Printf("%s\n", lines[i])

		var x1, x2, y1, y2 int64

		var vec, points []string

		vec = strings.Split(lines[i], "->")
		points = strings.Split(vec[0], ",")
		//fmt.Printf("|%s|\n",points[1])
		x1, _ = strconv.ParseInt(strings.Trim(points[0], " "), 10, 64)
		y1, _ = strconv.ParseInt(strings.Trim(points[1], " "), 10, 64)
		//fmt.Printf("%d\n",y1)
		points = strings.Split(vec[1], ",")
		x2, _ = strconv.ParseInt(strings.Trim(points[0], " "), 10, 64)
		y2, _ = strconv.ParseInt(strings.Trim(points[1], " "), 10, 64)

		vectors = append(vectors, Vector{x1, x2, y1, y2})
		fmt.Printf("x1[%d], y1[%d] => x2[%d], y2[%d]\n", x1, y1, x2, y2)
	}

	// set maxXY
	for _, v := range vectors {
		if v.x1 > maxX {
			maxX = v.x1
		}
		if v.x2 > maxX {
			maxX = v.x2
		}
		if v.y1 > maxY {
			maxY = v.y1
		}
		if v.y2 > maxY {
			maxY = v.y2
		}
	}

	// create ventMap maxXY
	ventMap := make([][]int64, maxX+1)
	for r := range ventMap {
		ventMap[r] = make([]int64, maxY+1)
	}

	// plot each vector
	for _, v := range vectors {
		var cx, minx, maxx, cy, miny, maxy int64 // use nested for to traverse line
		// get maxX and maxY to use one traversal but will instead
		if v.x1 > v.x2 {
			maxx = v.x1
			minx = v.x2
		} else {
			maxx = v.x2
			minx = v.x1
		}
		if v.y1 > v.y2 {
			maxy = v.y1
			miny = v.y2
		} else {
			maxy = v.y2
			miny = v.y1
		}
		fmt.Printf("v= %s\n", v)
		if v.x1 == v.x2 || v.y1 == v.y2 {
			for cx = minx; cx <= maxx; cx++ {
				for cy = miny; cy <= maxy; cy++ {
					//fmt.Printf("cxy[%d,%d]", cx, cy)

					// check if position
					ventMap[cx][cy]++
					printMap(ventMap)

				}
			}
		}
		print("\n")
		// get value at each point and increment
	}
	var result int = 0
	for c := range ventMap {
		for r := range ventMap[c] {
			if ventMap[r][c] >= 2 {
				result++
			}
		}
	}

	//fmt.Printf("Winning Score = %d", scoreWinner(cards[winner].c, calledNumbers))
	//fmt.Printf("%d %d\n", gamma, epsilon)
	//var result int64 = gamma * epsilon
	fmt.Printf("Result: %d\n", result)
}
func printMap(vm [][]int64) {

	for c := range vm {
		for r := range vm[c] {
			fmt.Printf("%d", vm[r][c])
		}
		print("\n")
	}
	print("=========================\n")

}
