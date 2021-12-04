package AoC

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func D3_1() {
	dat, err := ioutil.ReadFile("inputd3.1")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	var lines = strings.Split(string(dat), "\n")

	var gamma int64 = 0
	var epsilon int64 = 0

	// for each bit
	var b int
	var bitsize = len(lines[0])
	for b = 1; b <= bitsize; b++ {
		var mask int16 = 1 << (bitsize - b)
		var bitcount int = 0
		var i int
		for i = 0; i < len(lines); i++ {
			//fmt.Printf("%s\n", lines[i])
			var val int64
			val, _ = strconv.ParseInt(lines[i], 2, bitsize+1)

			if int16(val)&mask > 0 {
				bitcount++
			}

		}
		fmt.Printf("1 %d bitcount %d, mask %d	 \n", b, bitcount, mask)
		// if current bit is most commonly 1 set it in gamma
		if bitcount > len(lines)/2 {
			gamma += int64(mask)
		} else {
			epsilon += int64(mask)
		}
	}
	fmt.Printf("%d %d\n", gamma, epsilon)
	var result int64 = gamma * epsilon
	fmt.Printf("Result: %d\n", result)
}
