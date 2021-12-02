package AoC

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func D2_2() {
	dat, err := ioutil.ReadFile("inputd2.1")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	//fmt.Print(string(dat))

	var lines = strings.Split(string(dat), "\n")

	var depth int64 = 0
	var distance int64 = 0
	var aim int64 = 0

	var i int
	for i = 0; i < len(lines); i++ {
		fmt.Printf("%s\n", lines[i])
		var values = strings.Split(string(lines[i]), " ")
		fmt.Printf("dir:%s, dist:%s\n", values[0], values[1])
		if values[0] == "forward" {
			var d int64
			d, _ = strconv.ParseInt(values[1], 10, 0)
			distance += d
			depth += aim * d
		}
		if values[0] == "down" {
			var d int64
			d, _ = strconv.ParseInt(values[1], 10, 0)
			aim += d
		}
		if values[0] == "up" {
			var d int64
			d, _ = strconv.ParseInt(values[1], 10, 0)
			aim -= d
		}
		//append(distance, 1)

		//depths[i], _ = strconv.ParseInt(lines[i], 10, 0)
	}

	//fmt.Printf("list %q\n",lines)
	var result int64 = distance * depth

	//for i = 4; i < len(depths); i++ {
	//
	//	var l, h int64
	//	l = depths[i-4]+depths[i-3]+depths[i-2]
	//	h = depths[i-3]+depths[i-2]+depths[i-1]
	//	if l < h {
	//		result++
	//		fmt.Printf("%d < %d - increase, count %d\n", l, h, result)
	//
	//	} else{
	//		fmt.Printf("%d < %d - decrease, count %d\n", l, h, result)
	//	}
	//
	//}
	fmt.Printf("Result: %d\n", result)
}
