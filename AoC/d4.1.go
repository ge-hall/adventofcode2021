package AoC

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type Card struct {
	c [][]string
}

func D4_1() {
	dat, err := ioutil.ReadFile("inputd4.1")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	var lines = strings.Split(string(dat), "\n")

	var i int
	for i = 0; i < len(lines); i++ {
		fmt.Printf("%s\n", lines[i])
	}

	// read bingo numbers
	var bingo []string = strings.Split(lines[0], ",")
	fmt.Printf("%s\n", bingo)
	// read cards
	var cards []Card
	for i = 2; i < len(lines); i += 6 {

		if lines[i] == "" {
			continue
		}
		// read next 5 lines into card
		var card = make([][]string, 5)
		var l int
		for l = 0; l < 5; l++ {
			fmt.Printf("|%s|\n", lines[i+l])
			// read line and split into current row
			var arr = strings.Split(lines[i+l], " ")
			for _, e := range arr {
				if e != "" {
					card[l] = append(card[l], e)
				}
			}
			fmt.Printf("length of inserted array: %d", len(card[l]))
			// on 5th row add card to cards
			if l == 4 {
				cards = append(cards, Card{c: card})
			}
		}

	}
	fmt.Printf("%s\n", cards)

	// call bingo numbers and check cards each round
	var round int
	var calledNumbers []string
	var winner int = -1
	for round = 0; round < len(bingo); round++ {
		// call number
		calledNumbers = append(calledNumbers, bingo[round])
		fmt.Printf("called numbers: %s||%s", calledNumbers, bingo[round])
		// check each card

		var card int
		for card = 0; card < len(cards); card++ {
			fmt.Printf("%d\n\n", len(cards[card].c))
			var r int
			var c int
			//for r = 0; r < len(cards[card].c); r ++ {
			//for c = 0 ; c < len(cards[card].c[r]); c++{
			//	fmt.Printf("%s ", cards[card].c[r][c])
			//for _, r := range cards[card].c {
			for r = 0; r < len(cards[card].c); r++ {
				var checked int = 0
				//for _, c := range r {
				for c = 0; c < len(cards[card].c[r]); c++ {
					fmt.Printf("%s ", cards[card].c[r][c])
					//fmt.Printf("[%d,%d]%s ", r, c, cards[card].c[r][c])
					if contains(calledNumbers, cards[card].c[r][c]) {
						checked++
						fmt.Printf("found match: %d", cards[card].c[r][c])
					}
				}
				print("||\n")
				if checked == 5 {
					fmt.Printf("*******************WINNER**********************\n")
					winner = card

				}
			}
			//var r int
			//var c int
			r = 0
			c = 0
			print("==================\n")
			fmt.Printf("for c = 0 c < %d\n", len(cards[card].c))
			for c = 0; c < len(cards[card].c); c++ {
				var checked int = 0
				for r = 0; r < len(cards[card].c); r++ {

					fmt.Printf("[%d,%d]%s ", r, c, cards[card].c[r][c])
					if contains(calledNumbers, cards[card].c[r][c]) {
						checked++
					}
				}
				print("||\n")
				if checked == 5 {
					fmt.Printf("*******************WINNER**********************\n")
					winner = card

				}
			}
			if winner == card {
				break
			}

		}
		if winner >= 0 {
			break
		}

	}
	// check winning card
	fmt.Printf("winning card %d\n", winner)

	fmt.Printf("Winning Score = %d", scoreWinner(cards[winner].c, calledNumbers))
	//fmt.Printf("%d %d\n", gamma, epsilon)
	//var result int64 = gamma * epsilon
	//fmt.Printf("Result: %d\n", result)
}

func scoreWinner(card [][]string, calledNumbers []string) int64 {
	var r int
	var c int
	var score int64

	for r = 0; r < len(card); r++ {
		//for _, c := range r {
		for c = 0; c < len(card[r]); c++ {
			if !contains(calledNumbers, card[r][c]) {
				var v, _ = strconv.ParseInt(card[r][c], 10, 64)
				score += v
			}
		}
	}
	var lastCalled = calledNumbers[len(calledNumbers)-1]
	var v, _ = strconv.ParseInt(lastCalled, 10, 64)
	return score * v
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
