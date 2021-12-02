package AoC

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func D2_1() {
	dat, err := ioutil.ReadFile("inputd2.2")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	//fmt.Print(string(dat))

	var lines = strings.Split(string(dat), "\n")
	var depths = make([]int64, len(lines))
	var i int
	for i = 0; i< len(lines); i++ {
		depths[i], _ = strconv.ParseInt(lines[i], 10, 0)
	}

	fmt.Printf("list %q\n",lines)
	var result int = 0

	for i = 4; i < len(depths); i++ {

		var l, h int64
		l = depths[i-4]+depths[i-3]+depths[i-2]
		h = depths[i-3]+depths[i-2]+depths[i-1]
		if l < h {
			result++
			fmt.Printf("%d < %d - increase, count %d\n", l, h, result)

		} else{
			fmt.Printf("%d < %d - decrease, count %d\n", l, h, result)
		}

	}
	fmt.Printf("Result: %d\n", result)
}

