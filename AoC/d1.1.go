package AoC

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func D1_1() {
	dat, err := ioutil.ReadFile("inputd1")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	var lines = strings.Split(string(dat), "\n")
	fmt.Printf("list %q\n", lines)
	var result int = 0
	var i int
	var records = len(lines)
	fmt.Printf("Number of lines = %d\n", records)
	for i = 0; i < records; i++ {

		if i == 0 {
			continue
		}
		var l, h int64
		l, _ = strconv.ParseInt(lines[i-1], 10, 16)
		h, _ = strconv.ParseInt(lines[i], 10, 16)
		if l < h {
			result++
			fmt.Printf("%d < %d - increase, count %d\n", l, h, result)

		} else {
			fmt.Printf("%d < %d - decrease, count %d\n", l, h, result)
		}

	}
	fmt.Printf("Result: %d\n", result)
}
