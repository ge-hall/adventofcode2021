package AoC

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func D6_2() {
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
	var days int = 256
	// require fish counter array
	var fishCount [9]int
	for _, f := range lanternFish {
		fishCount[f]++
	}
	fmt.Printf("%o\n", fishCount)
	for i = 0; i < days; i++ {
		SoBCount := fishCount
		for fi, _ := range fishCount {
			if fi == 0 {
				fishCount[8] = SoBCount[0]
				SoBCount[7] += SoBCount[0] // using pos 7 here to update new fish reset as next step will move 7 => 6
				fishCount[0] = 0
			} else {
				if SoBCount[fi] == 0 {
					fishCount[fi-1] = 0
				} else {
					fishCount[fi-1] = SoBCount[fi]
				}
			}

		}
		fmt.Printf("%o\n", fishCount)
	}

	var sum int = 0
	for _, f := range fishCount {
		sum += f
	}
	fmt.Printf("%d", sum)
	var result int = len(lanternFish)
	fmt.Printf("Result: %d\n", result)
}
