package AoC

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func D9_1() {
	dat, err := ioutil.ReadFile("inputd9.1")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	var heightMap = strings.Split(string(dat), "\n")
	for _, heights := range heightMap {
		fmt.Printf("%s\n", heights)
	}

	var result int = 0
	fmt.Printf("Result: %d\n", result)
}
