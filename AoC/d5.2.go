package AoC

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func D5_2() {
	dat, err := ioutil.ReadFile("inputd5.1")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	var lines = strings.Split(string(dat), "\n")

	var i int
	for i = 0; i < len(lines); i++ {
		fmt.Printf("%s\n", lines[i])
	}

	// check winning card
	//fmt.Printf("winning card %d\n", winner)

	//fmt.Printf("Winning Score = %d", scoreWinner(cards[winner].c, calledNumbers))
	//fmt.Printf("%d %d\n", gamma, epsilon)
	//var result int64 = gamma * epsilon
	//fmt.Printf("Result: %d\n", result)
}
