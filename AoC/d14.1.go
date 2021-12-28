package AoC

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

func D14_1() {
	dat, err := ioutil.ReadFile("inputd14")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	var data = strings.Split(string(dat), "\n")

	var template string = ""
	var insertionRules map[string]string = make(map[string]string)
	var readTemplate = false
	for _, line := range data {
		fmt.Printf("%s\n", line)
		if !readTemplate {
			// read template
			template = line
			readTemplate = true
		} else if line == "" {
		} else {
			// insertionRule
			rule := strings.Split(line, " -> ")
			insertionRules[rule[0]] = rule[1]
		}

	}

	// build polymerisation
	// read pairs into map
	pairs := make(map[string]int)
	newPairs := make(map[string]int)
	for a, b := 0, 1; b < len(template); {
		pairs[string(template[a])+string(template[b])]++
		a++
		b++
	}
	fmt.Printf("pairs:%v\n", pairs)

	for i := 0; i < 40; i++ {
		for k, v := range pairs {
			newPairs[k] = v
		}

		for k, v1 := range pairs {
			if v1 <= 0 {
				continue
			}
			fmt.Printf("testing pair %s\n", k)
			for rule, v := range insertionRules {

				if rule == k {
					fmt.Printf(" k = %s, v = %s, rule = %s, quant=%d\n", k, v, rule, v1)
					newPairs[k] -= v1
					newPairs[string(k[0])+v] += v1
					newPairs[v+string(k[1])] += v1
					break
				}
			}

		}
		for k, v := range newPairs {
			if v < 0 {
				v = 0
			}
			pairs[k] = v

		}

		fmt.Printf("newPairs %v\n", newPairs)

	}
	print("completed polymerisation\n")
	// count elements, sort and find max - min
	//elements := make(map[string]int)
	//
	//for c := 0; c < len(newPolymer); c++ {
	//	elements[string(newPolymer[c])] = elements[string(newPolymer[c])] + 1
	//}
	// count left element of pair +1 if we are counting last element of template
	elements := make(map[string]int)
	elements[string(template[len(template)-1])] = 1
	values := make([]int, 0)
	for k, v := range pairs {

		elements[string(k[0])] += v

	}
	for _, v := range elements {
		values = append(values, v)
	}
	sort.Ints(values)

	////fmt.Printf("newPolymer %s with len = %d\n", newPolymer, len(newPolymer))
	fmt.Printf("result %d - %d = %d\n%v", values[len(values)-1], values[0], values[len(values)-1]-values[0], values)

}
