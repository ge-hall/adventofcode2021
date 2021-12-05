package AoC

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func D4_1() {
	dat, err := ioutil.ReadFile("inputd4.1")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	var lines = strings.Split(string(dat), "\n")

	var i int
	for i = 0; i < len(lines); i++ {
		fmt.Print(lines[i])
	}

	//fmt.Printf("%d %d\n", gamma, epsilon)
	//var result int64 = gamma * epsilon
	//fmt.Printf("Result: %d\n", result)
}
