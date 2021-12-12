package AoC

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func D8_1() {
	dat, err := ioutil.ReadFile("inputd8.1")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	var entries = strings.Split(string(dat), "\n")
	//var signals [] string
	var outputs []string
	// read values into pos array
	for _, e := range entries {
		fmt.Printf("%s\n\n", e)
		var line = strings.Split(e, "|")
		outputs = append(outputs, line[1])

	}
	// process outputs and count unique segment patterns 1, 4, 7, 8
	var uniqueCount int = 0
	for _, o := range outputs {
		// get current outputs
		outs := strings.Split(o, " ")
		for _, digit := range outs {
			fmt.Printf("digit:%s\n", digit)
			if len(digit) == 2 || len(digit) == 3 || len(digit) == 4 || len(digit) == 7 {
				uniqueCount++
			}
		}
		fmt.Printf("uniqueCount:%s\n\n", uniqueCount)

	}
	var result int = uniqueCount
	fmt.Printf("Result: %d\n", result)
}
