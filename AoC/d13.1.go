package AoC

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func D13_1() {
	dat, err := ioutil.ReadFile("inputd13")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	var data = strings.Split(string(dat), "\n")
	var cs CaveSystem
	for _, line := range data {
		fmt.Printf("%s\n", line)
		// split into from and to
		caves := strings.Split(line, "-")
		// add both caves

		// add connection
		cs.ConnectCave(cs.AddCave(caves[0]), cs.AddCave(caves[1]))
		cs.ConnectCave(cs.AddCave(caves[1]), cs.AddCave(caves[0]))

	}
	print("===============\n\n")
	cs.String()
	print("===============\n\n")
	cs.Traverse()
	print(pathCount)
}
