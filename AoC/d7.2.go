package AoC

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

func D7_2() {
	dat, err := ioutil.ReadFile("inputd7.1")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	var lines = strings.Split(string(dat), ",")
	var pos []int
	// read values into pos array
	for _, l := range lines {
		fmt.Printf("%s\n", l)
		var v, _ = strconv.ParseInt(l, 10, 64)
		pos = append(pos, int(v))
	}
	// simple parse each value and sum the fuel
	// get max value
	var maxValue int = 0
	for _, v := range pos {
		if v > maxValue {
			maxValue = v
		}
	}
	var fuel []int64 = make([]int64, maxValue+1, maxValue+1)
	for p, _ := range fuel {
		for _, m := range pos {
			//if fuel[p] > 0 { continue}
			fmt.Printf("p=%d, m=%d = %d\n", p, m, getN(int64(math.Abs(float64(m-p)))))

			fuel[p] += int64(getN(int64(math.Abs(float64(m - p)))))
		}
		fmt.Printf("Fuel for Pos %d = %d\n", p, fuel[p])
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
func getN(d int64) int64 {
	//fmt.Printf("%d\n", d)

	if d <= 1 {
		return d
	}
	return d + getN(d-1)
}
