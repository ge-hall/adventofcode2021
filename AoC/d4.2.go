package AoC

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func D4_2() {
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
	var pastWinners []int
	var lastCalled string
	for round = 0; round < len(bingo); round++ {
		// call number
		calledNumbers = append(calledNumbers, bingo[round])

		// check each card

		var card int
		for card = 0; card < len(cards); card++ {
			//fmt.Printf("%d , card number %d\n\n", len(cards[card].c), card)
			// if pastwinner skip
			var skip = false
			for _, w := range pastWinners {
				if w == card {
					skip = true
				}
			}
			if skip {
				continue
			}
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
					//fmt.Printf("%s ", cards[card].c[r][c])
					//fmt.Printf("[%d,%d]%s ", r, c, cards[card].c[r][c])
					if contains(calledNumbers, cards[card].c[r][c]) {
						checked++
						//fmt.Printf("found match: %d", cards[card].c[r][c])
					}
				}
				//print("||\n")
				if checked == 5 {
					printCard(cards[card].c, r, -1)
					winner = card
					lastCalled = calledNumbers[len(calledNumbers)-1]
					fmt.Printf("*******************WINNER %d [%s]**********************\n", card, lastCalled)
					fmt.Printf("called numbers: %s\nlast called|%s|\n%d\n\n", calledNumbers, bingo[round], pastWinners)
					pastWinners = append(pastWinners, card)
					break

				}
			}
			if winner == card {
				continue
			}

			//var r int
			//var c int
			r = 0
			c = 0
			//print("==================\n")
			//fmt.Printf("for c = 0 c < %d\n", len(cards[card].c))
			for c = 0; c < len(cards[card].c); c++ {
				var checked int = 0
				for r = 0; r < len(cards[card].c); r++ {

					//fmt.Printf("[%d,%d]%s ", r, c, cards[card].c[r][c])
					if contains(calledNumbers, cards[card].c[r][c]) {
						checked++
					}
				}
				//print("||\n")
				if checked == 5 {

					printCard(cards[card].c, -1, c)
					winner = card
					lastCalled = calledNumbers[len(calledNumbers)-1]
					fmt.Printf("*******************WINNER %d [%s]**********************\n", card, lastCalled)
					fmt.Printf("called numbers: %s\nlast called|%s|\n%d\n\n", calledNumbers, bingo[round], pastWinners)
					pastWinners = append(pastWinners, card)
					break
				}
			}

		}
		//if winner >= 0 {
		//	break
		//}

	}
	// check winning card
	fmt.Printf("winning card %d\n", winner)

	fmt.Printf("Winning Score = %d", scoreLastWinner(cards[winner].c, calledNumbers, lastCalled))
	//fmt.Printf("%d %d\n", gamma, epsilon)
	//var result int64 = gamma * epsilon
	//fmt.Printf("Result: %d\n", result)
}

func scoreLastWinner(card [][]string, calledNumbers []string, lastCalled string) int64 {
	var r int
	var c int
	var score int64
	fmt.Printf("LastCalled: %s", lastCalled)
	// trim array to lastCalled
	var calledNumbersToWin []string
	for _, e := range calledNumbers {
		calledNumbersToWin = append(calledNumbersToWin, e)
		if e == lastCalled {
			break
		}

	}

	for r = 0; r < len(card); r++ {
		//for _, c := range r {
		for c = 0; c < len(card[r]); c++ {
			if !contains(calledNumbersToWin, card[r][c]) {
				var v, _ = strconv.ParseInt(card[r][c], 10, 64)
				score += v
			}
		}
	}

	var v, _ = strconv.ParseInt(lastCalled, 10, 64)
	return score * v
}

func printCard(card [][]string, row int, col int) {
	var r int
	var c int

	for r = 0; r < len(card); r++ {
		if r == row {
			print("|")
		}
		for c = 0; c < len(card[r]); c++ {
			if c == col {
				print("|")
			}
			fmt.Printf("%s ", card[r][c])
			if c == col {
				print("|")
			}
			//fmt.Printf("[%d,%d]%s ", r, c, cards[card].c[r][c])
		}
		if r == row {
			print("|")
		}
		print("\n\n")
	}
}
