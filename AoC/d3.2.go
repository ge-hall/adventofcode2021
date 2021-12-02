package AoC

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func D3_2() {
	dat, err := ioutil.ReadFile("inputd3.1")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

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
	}

	var result int64 = distance * depth

	fmt.Printf("Result: %d\n", result)
}
