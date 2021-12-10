package AoC

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func D6_1() {
	dat, err := ioutil.ReadFile("inputd6.1")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	var lines = strings.Split(string(dat), "\n")

	// read lanternfish
	var lanternFish []int64 = make([]int64, 0, 10000)
	var i int
	for i = 0; i < len(lines); i++ {
		var fish []string = strings.Split(lines[i], ",")
		for _, f := range fish {
			var fint, _ = strconv.ParseInt(f, 10, 64)
			lanternFish = append(lanternFish, fint)
		}
		fmt.Printf("%s\n", lines[i])

	}
	var days int = 200
	// This is too slow for Part 2 so will move code across to new d2 file
	for i = 0; i < days; i++ {
		for lfi, lf := range lanternFish {
			if lf == 0 {
				// reset and spawn
				lanternFish[lfi] = 6
				var lfCap, lfLen int = cap(lanternFish), len(lanternFish)
				if lfCap == lfLen {
					target := make([]int64, lfLen, lfCap+10000)
					copy(target, lanternFish)
					lanternFish = target
				} else {
					lanternFish = append(lanternFish, 8)
				}
			} else {
				lanternFish[lfi]--
			}
		}

	}

	var result int = len(lanternFish)
	fmt.Printf("Result: %d\n", result)
}
