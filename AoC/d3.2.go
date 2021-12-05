package AoC

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func D3_2() {

	dat, err := ioutil.ReadFile("inputd3.1")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	var lines = strings.Split(string(dat), "\n")

	// convert string array to int array
	var bits = len(lines[0])
	var bitStringArray []int
	var i int
	for i = 0; i < len(lines); i++ {
		//		//fmt.Printf("%s\n", lines[i])
		var val int64
		val, _ = strconv.ParseInt(lines[i], 2, bits+1)
		bitStringArray = append(bitStringArray, int(val))
	}
	fmt.Printf("%d\n", bitStringArray[0])

	var oxGenRate int = 0
	var cO2ScrubRate int = 0

	oxGenRate = findCommonBitStringValue(bitStringArray, bits, true)
	cO2ScrubRate = findCommonBitStringValue(bitStringArray, bits, false)

	fmt.Printf("%d %d\n", oxGenRate, cO2ScrubRate)
	var result int = oxGenRate * cO2ScrubRate
	fmt.Printf("Result: %d\n", result)
}

func findCommonBitStringValue(bitString []int, bits int, mostCommon bool) int {

	var ones []int
	var zeros []int
	var result []int = bitString
	var b int
	for b = 1; b <= bits; b++ {
		var bit int16 = 1 << (bits - b)
		fmt.Printf("%012b\n", bit)
		var i int
		for i = 0; i < len(result); i++ {
			fmt.Printf("%012b\n", result[i])

			if int16(result[i])&bit > 0 {
				ones = append(ones, result[i])
			} else {
				zeros = append(zeros, result[i])
			}

		}
		fmt.Printf("%d\n", len(ones))
		fmt.Printf("%d\n", len(zeros))
		// set result to common target
		if len(ones) >= len(zeros) {
			if mostCommon {
				result = ones
			} else {
				result = zeros
			}
		} else {
			if mostCommon {
				result = zeros
			} else {
				result = ones
			}
		}
		ones = nil
		zeros = nil

		if len(result) == 1 {
			break
		}
	}

	fmt.Printf("%d\n", len(result))
	return result[0]

}
