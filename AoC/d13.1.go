package AoC

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func D13_1() {
	dat, err := ioutil.ReadFile("inputd13")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	var data = strings.Split(string(dat), "\n")
	//var paper [][]int
	//var plotting = true
	for _, line := range data {
		fmt.Printf("%s\n", line)
		if line == "" {
			print("empty line")
		}

	}
	print("===============\n\n")
	print("===============\n\n")
	print(pathCount)
}
