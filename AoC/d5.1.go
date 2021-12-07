package AoC

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func D5_1() {
	dat, err := ioutil.ReadFile("inputd5.1")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	var lines = strings.Split(string(dat), "\n")

	//var map = [][]int := make([][]int, 0, 1000)
	//print(map)
	// for each lineSpec
	var i int
	for i = 0; i < len(lines); i++ {
		fmt.Printf("%s\n", lines[i])

		var cx, x1, x2, cy, y1, y2 int64

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
		fmt.Printf("x1[%d], y1[%d] => x2[%d], y2[%d]\n", x1, y1, x2, y2)

		// use nested for to traverse line
		for cx = x1; cx < x2; cx++ {
			for cy = y1; cy < y2; cy++ {
				fmt.Printf("cxy[%d,%d]", cx, cy)
			}
		}
		// get value at each point and increment
	}

	//fmt.Printf("Winning Score = %d", scoreWinner(cards[winner].c, calledNumbers))
	//fmt.Printf("%d %d\n", gamma, epsilon)
	//var result int64 = gamma * epsilon
	//fmt.Printf("Result: %d\n", result)
}
