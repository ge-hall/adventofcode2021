package AoC

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

type Point struct {
	i int
	j int
	v int
}

var graphspace = make([][]int, 0)
var risk [][]int
var sptSet [][]bool

var heads = []Point{{0, 0, 0}}

func D15_1() {
	dat, err := ioutil.ReadFile("inputd15")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	var data = strings.Split(string(dat), "\n")
	// read data into graph as addressable matrix

	for _, line := range data {
		fmt.Printf("%s\n", line)
		newLine := make([]int, 0)
		for _, c := range line {
			v, _ := strconv.ParseInt(string(c), 10, 8)
			newLine = append(newLine, int(v))
		}
		graphspace = append(graphspace, newLine)

	}
	// implement Dijkstra without adjacency matrix or priority queue perhaps
	// not planning on finding a library, might sneak a peak in an old data structure text.
	fmt.Printf("%v\n", graphspace)
	risk = make([][]int, len(graphspace))
	sptSet = make([][]bool, len(graphspace))
	for i, _ := range risk {
		risk[i] = make([]int, len(graphspace[0]))
		sptSet[i] = make([]bool, len(graphspace[0]))
		for j, _ := range risk[i] {
			risk[i][j] = math.MaxInt64
			sptSet[i][j] = false
		}
	}
	fmt.Printf("%v\n", risk)
	fmt.Printf("%v\n", sptSet)
	risk[0][0] = 0
	totalRisk := 0
	sptSet[0][0] = true
	// iterate until complete with break
	// dev with short iteration control

	for c := 0; c < 3; c++ {
		// find lowest risk path from current heads
		for _, point := range heads {
			i := point.i
			j := point.j
			fmt.Printf("%d, %d,[%d]\n", i, j, graphspace[i][j])
			if i == len(graphspace)-1 && j == len(graphspace[0])-1 {
				break
			}
			// min risk
			// check neighbour with lowest risk
			var t, r, b, l, minVal int = math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64
			if checkBounds(graphspace, i+1, j) && !sptSet[i+1][j] {

				b = getVal(graphspace, i+1, j)
				if b < minVal {
					minVal = b
				}
			}
			if checkBounds(graphspace, i, j-1) && !sptSet[i][j-1] {
				l = getVal(graphspace, i, j-1)

				if l < minVal {
					minVal = l
				}
			}
			if checkBounds(graphspace, i-1, j) && !sptSet[i-1][j] {

				t = getVal(graphspace, i-1, j)

				if t < minVal {
					minVal = t
				}
			}
			if checkBounds(graphspace, i, j+1) && !sptSet[i][j+1] {

				r = getVal(graphspace, i, j+1)

				if r < minVal {
					minVal = r
				}
			}
			// need to move this check up so it compares each minimum for each head.
			// since we are storing the minimum in the head Point we can have a second enumeration to get teh minimum to follow.
			// if not in sptSet && risk < risk
			// set risk
			if t == minVal {
				fmt.Printf("t is min %d\n", t)
				heads = append(heads, Point{i - 1, j, minVal + point.v})
				//risk[i-1][j] = minVal + risk[i][j]
				setVisited(i, j)

			}
			if r == minVal {
				fmt.Printf("r is min %d\n", r)
				heads = append(heads, Point{i, j + 1, minVal + point.v})
				//risk[i][j+1] = minVal + risk[i][j]
				setVisited(i, j)
			}
			if b == minVal {
				fmt.Printf("b is min %d\n", b)
				heads = append(heads, Point{i + 1, j, minVal + point.v})
				//risk[i+1][j] = minVal + risk[i][j]
				setVisited(i, j)
			}
			if l == minVal {
				fmt.Printf("l is min %d\n", l)
				heads = append(heads, Point{i, j - 1, minVal + point.v})
				//risk[i][j-1] = minVal + risk[i][j]
				setVisited(i, j)
			}

		}
		fmt.Printf("content of heads: %v\n", heads)
		fmt.Printf("risk =: %v\n", risk)
		// check heads for minimum value
		var minHead *Point
		var winners []Point
		for _, head := range heads {
			if minHead == nil {
				minHead = &head
			} else if minHead.v <= head.v {
				winners = append(winners, head)

			}

		}
		for _, winner := range winners {
			risk[winner.i][winner.j] = winner.v
			setVisited(winner.i, winner.j)
		}

	}
	fmt.Printf("TotalRisk is %d", totalRisk)

}
func setVisited(i int, j int) {
	sptSet[i][j] = true
	removeHead(i, j)
}
func removeHead(i int, j int) {
	// find position
	removeIndex := -1
	for hi, p := range heads {
		if p.i == i && p.j == j {
			removeIndex = hi
		}
	}
	if removeIndex == -1 {
		return
	}
	heads = append(heads[0:removeIndex], heads[removeIndex+1:len(heads)]...)
}
func checkBounds(space [][]int, i int, j int) bool {
	return i >= 0 && j >= 0 && i < len(space) && j < len(space[0])
}

func getVal(space [][]int, i int, j int) int {
	if !checkBounds(space, i, j) {
		return math.MaxInt64
	}
	return space[i][j]
}
