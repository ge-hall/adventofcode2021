package AoC

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func D16_1() {
	dat, err := ioutil.ReadFile("inputd16")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	var data = strings.Split(string(dat), "\n")
	// read data into bit array converting hex into binary
	for _, line := range data {
		fmt.Printf("%s\n", line)

	}
	// implement Dijkstra

}
