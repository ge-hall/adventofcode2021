package AoC

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strings"
)

func D8_2() {
	dat, err := ioutil.ReadFile("inputd8.1")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	var entries = strings.Split(string(dat), "\n")
	var signals []string
	var outputs []string
	// read values into pos array
	for _, e := range entries {
		fmt.Printf("%s\n\n", e)
		var line = strings.Split(e, "|")
		signals = append(signals, line[0])
		outputs = append(outputs, line[1])

	}
	//{"acedgfb": 8, "cdfbe": 5, "gcdfa": 2, "fbcad": 3, "dab": 7, "cefabd": 9, "cdfgeb": 6, "eafb": 4, "cagedb": 0, "ab": 1}

	// process outputs and count unique segment patterns 1, 4, 7, 8
	var uniqueCount int = 0
	for i, o := range outputs {
		// build map
		var m = make(map[string]int)
		var wires = strings.Split(signals[i], " ")
		var wireOne string
		var wireNine string
		var wireFour string
		var looseWires []string

		// process fixed size
		for _, wire := range wires {
			if len(wire) == 2 {
				m[wire] = 1
				wireOne = wire
			} else if len(wire) == 3 {
				m[wire] = 7
			} else if len(wire) == 4 {
				m[wire] = 4
				wireFour = wire
			} else if len(wire) == 7 {
				m[wire] = 8
			}
		}
		//fmt.Printf("%o - %d \n", wires, wireOne)
		//fmt.Printf("subChar %s, %s, %s\n",wireFour, wireOne, subtractChars(wireFour, wireOne))
		//process variable size
		for _, wire := range wires {
			if len(wire) == 6 {
				if containsChars(wire, wireFour) {
					m[wire] = 9
					wireNine = wire
				} else if containsChars(wire, subtractChars(wireFour, wireOne)) && !containsChars(wire, wireOne) {
					m[wire] = 6
				} else {
					m[wire] = 0
				}
			} else if len(wire) == 5 {
				if containsChars(wire, wireOne) {
					m[wire] = 3
				} else if len(wireNine) > 0 && containsChars(wireNine, wire) {
					m[wire] = 5
				} else if len(wireNine) > 0 {
					m[wire] = 2
				} else {
					looseWires = append(looseWires, wire)
				}
			}

		}
		//fmt.Printf("%o == %o\n", m, looseWires)
		for _, wire := range looseWires {
			if len(wire) == 5 {
				if containsChars(wire, wireOne) {
					m[wire] = 3
				} else if len(wireNine) > 0 && containsChars(wireNine, wire) {
					m[wire] = 5
				} else if len(wireNine) > 0 {
					m[wire] = 2
				}
			}
		}
		fmt.Printf("%s \n", wires)
		for k, v := range m {
			fmt.Printf("%s, %d \n", k, v)
		}

		// get current outputs
		outs := strings.Split(o, " ")
		for i, digit := range outs {
			uniqueCount += int(math.Pow10(4-i)) * decodeFromMap(digit, m)
			//if len(digit) == 2 {
			//	uniqueCount += int(math.Pow10(4-i)) * 1
			//} else if len(digit) == 4 {
			//	uniqueCount += int(math.Pow10(4-i)) * 4
			//} else if len(digit) == 3 {
			//	uniqueCount += int(math.Pow10(4-i)) * 7
			//} else if len(digit) == 7 {
			//	uniqueCount += int(math.Pow10(4-i)) * 8
			//} else if len(digit) == 0 {
			//	continue
			//} else {
			//	// check anagram of map
			//
			//	uniqueCount += int(math.Pow10(4-i)) * decodeFromMap(digit, m)
			//}
			//fmt.Printf("p, %4.0f digit:%s:%d\n", math.Pow10(4-i), digit, uniqueCount)
		}
		//fmt.Printf("uniqueCount:%s\n\n", uniqueCount)

	}
	var result int = uniqueCount
	fmt.Printf("Result: %d\n", result)

	fmt.Printf("aegbdf containsChars dcbef %s", containsChars("aegbdf", "dbcfg"))
	fmt.Printf("aegbdf containsChars dcbef %s", containsChars("aegbdf", "dcbef"))
}
func decodeFromMap(digit string, m map[string]int) int {
	// for each m value where len(digit == len(key) check each char exists
	for k, v := range m {
		//fmt.Printf("Map Key = %s with value = %d for digit %s\n", k, v, digit)
		var ccount int = 0
		for _, c := range digit {
			//fmt.Printf("i%d %s:%s|\n", i, string(c), string(digit[i]))
			for _, sc := range k {
				if c == sc {
					ccount++
				}
			}
		}
		if ccount == len(digit) && ccount == len(k) {
			fmt.Printf("ret= %d\n", v)
			return v
		}
	}
	return 0
}
func containsChars(str string, substr string) bool {
	for _, c := range substr {
		var found = false
		for _, s := range str {
			if c == s {
				found = true
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func subtractChars(lop string, rop string) string {
	var diff string = ""
	for _, c := range rop {
		if !strings.Contains(lop, string(c)) {
			diff += string(c)
		}
	}
	return diff
}
