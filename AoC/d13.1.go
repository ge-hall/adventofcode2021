package AoC

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type Fold struct {
	direction string
	distance  int
}

func D13_1() {
	dat, err := ioutil.ReadFile("inputd13")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	var data = strings.Split(string(dat), "\n")
	var folds []Fold
	var plotting = true
	var x []int
	xMax := 0
	var y []int
	yMax := 0
	for _, line := range data {
		fmt.Printf("%s\n", line)
		if line == "" {
			print("empty line\n")
			plotting = false
			continue
		}
		if !plotting {
			// read instructions
			instr := strings.Split(line, " ")
			fold := strings.Split(instr[2], "=")
			distance, _ := strconv.ParseInt(fold[1], 10, 32)
			folds = append(folds, Fold{fold[0], int(distance)})
			continue
		}

		coord := strings.Split(line, ",")
		xv, _ := strconv.ParseInt(coord[0], 10, 32)
		if int(xv) > xMax {
			xMax = int(xv)
		}
		x = append(x, int(xv))

		yv, _ := strconv.ParseInt(coord[1], 10, 32)
		if int(yv) > yMax {
			yMax = int(yv)
		}
		y = append(y, int(yv))

	}
	paper := make([][]string, yMax+1)
	// add dimensions
	for y := 0; y <= yMax; y++ {
		paper[y] = make([]string, xMax+1)
	}
	for _, r := range paper {
		for i, _ := range r {
			r[i] = " "
		}
	}
	// print plotting
	for i := 0; i < len(x); i++ {
		paper[y[i]][x[i]] = "#"
	}
	//fmt.Printf("%v\n", paper)
	fmt.Printf("%d max = %d\n", yMax, xMax)

	for i, _ := range folds {
		for j, r := range paper {
			for k, c := range r {
				if folds[i].direction == "y" && j == folds[i].distance {
					print("- ")
				} else if folds[i].direction == "x" && k == folds[i].distance {
					print("|")
				} else {
					fmt.Printf("%s ", c)
				}
			}
			print("\n")
		}
		fold(paper, folds[i].distance, folds[i].direction)
		if folds[i].direction == "y" {
			paper = paper[:folds[i].distance]
		} else {
			for r, _ := range paper {
				paper[r] = paper[r][:folds[i].distance]
			}
		}
		// count plots
		var count int = 0
		for _, col := range paper {
			for _, p := range col {
				if p == "#" {
					count++
				}
			}
		}
		fmt.Printf("count = %d\n", count)
	}
	// count plots
	var count int = 0
	for _, col := range paper {
		for _, p := range col {
			if p == "#" {
				count++
			}
		}
	}
	for _, r := range paper {
		for _, c := range r {
			fmt.Printf("%s ", c)
		}
		print("\n")
	}
	//fmt.Printf("%v\n", paper)
	//fmt.Printf("%v fold = %d\n", x, folds[0])
	//fmt.Printf("%v max = %d\n", y, yMax)
	print("===============\n\n")
	print("===============\n\n")
	print(count)
}
func fold(paper [][]string, f int, dir string) {
	if dir == "y" {
		for fa, fb := 0, 1; fa < f; {
			for x := 0; x < len(paper[fa]); x++ {
				if paper[len(paper)-fb][x] == "#" {
					paper[fa][x] = paper[len(paper)-fb][x]
				}
			}
			fa++
			fb++
		}
	} else {
		for fa, fb := 0, 1; fa < f; {
			for y := 0; y < len(paper); y++ {
				if paper[y][len(paper[y])-fb] == "#" {
					paper[y][fa] = paper[y][len(paper[y])-fb]
				}
			}
			fa++
			fb++
		}
	}

}
