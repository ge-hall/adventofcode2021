package AoC

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

func D7_1() {
	dat, err := ioutil.ReadFile("inputd7.1")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	var lines = strings.Split(string(dat), ",")
	var pos []int64
	// read values into pos array
	for _, l := range lines {
		fmt.Printf("%s\n", l)
		var v, _ = strconv.ParseInt(l, 10, 64)
		pos = append(pos, v)
	}
	// simple parse each value and sum the fuel
	// get max value
	var maxValue int64 = 0
	for _, v := range pos {
		if v > maxValue {
			maxValue = v
		}
	}
	var fuel []int64 = make([]int64, maxValue+1, maxValue+1)
	for _, m := range pos {
		if fuel[m] > 0 {
			continue
		}
		for _, n := range pos {
			fmt.Printf("m=%d, n=%d = %d\n", m, n, int64(math.Abs(float64(m-n))))
			fuel[m] += int64(math.Abs(float64(m - n)))
		}
	}

	// get min fuel
	var minFuel int64 = math.MaxInt64
	for _, f := range fuel {
		fmt.Printf("%d, %d\n", f, minFuel)
		if f > 0 && f < minFuel {
			minFuel = f
		}
	}

	var result int64 = minFuel
	fmt.Printf("Result: %d\n", result)
}
