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
	i   int
	j   int
	v   int
	src *Point
}

var graphspace = make([][]int, 0)
var risk [][]int
var sptSet [][]bool

var nexts []Point
var heads = []Point{{0, 0, 0, nil}}
var prevHeads []Point

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
			risk[i][j] = 0 //math.MaxInt64
			sptSet[i][j] = false
		}
	}
	//fmt.Printf("%v\n", risk)
	//fmt.Printf("%v\n", sptSet)
	risk[0][0] = 0
	//totalRisk := 0
	sptSet[0][0] = true
	// iterate until complete with break
	// dev with short iteration control

	for c := 0; ; c++ {
		fmt.Printf("-%d-[%d]\n", c, countVisited())
		//fmt.Printf("content of heads: %v\n", heads)
		//fmt.Printf("content of prevheads: %v\n", prevHeads)
		// find lowest risk path from current heads
		for _, point := range heads {
			//var skip = false
			//for _, ph := range prevHeads {
			//	if point.i == ph.i && point.j == ph.j {
			//		skip = true
			//	}
			//}
			//if skip {
			//	continue
			//}
			i := point.i
			j := point.j
			//fmt.Printf("%d, %d,[%d]\n", i, j, graphspace[i][j])
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
			if minVal == math.MaxInt64 {
				continue
			}
			// need to move this check up so it compares each minimum for each head.
			// since we are storing the minimum in the head Point we can have a second enumeration to get teh minimum to follow.
			// if not in sptSet && risk < risk
			// set risk
			// TODO: a lot of appending will go on here so can save some time by updating existing points
			// e.g. graphspace is [][]Point and Point v is updated
			// nexts becomes an array of *Point and just contains pointers into graphspace
			if t == minVal {
				//fmt.Printf("t is min %d\n", t)
				if !nextExists(i-1, j, minVal+point.v, &point) {
					nexts = append(nexts, Point{i - 1, j, minVal + point.v, &point})
				}
				//risk[i-1][j] = minVal + risk[i][j]
				//setVisited(i, j)

			}
			if r == minVal {
				//fmt.Printf("r is min %d\n", r)
				if !nextExists(i, j+1, minVal+point.v, &point) {
					nexts = append(nexts, Point{i, j + 1, minVal + point.v, &point})
				}
				//risk[i][j+1] = minVal + risk[i][j]
				//setVisited(i, j)
			}
			if b == minVal {
				//fmt.Printf("b is min %d\n", b)
				if !nextExists(i+1, j, minVal+point.v, &point) {

					nexts = append(nexts, Point{i + 1, j, minVal + point.v, &point})
				}
				//risk[i+1][j] = minVal + risk[i][j]
				//setVisited(i, j)
			}
			if l == minVal {
				//fmt.Printf("l is min %d\n", l)
				if !nextExists(i, j-1, minVal+point.v, &point) {
					nexts = append(nexts, Point{i, j - 1, minVal + point.v, &point})
				}
				//risk[i][j-1] = minVal + risk[i][j]
				//setVisited(i, j)
			}

		}
		//fmt.Printf("content of nexts: %v\n", nexts)

		// check for minimum nexts value
		// convert to a head and save it's head to spt
		if len(nexts) == 0 {
			continue
		}
		var minNext Point = nexts[0]
		for _, next := range nexts {
			if next.v <= minNext.v {

				minNext = next
			}
		}
		if minNext.v == math.MaxInt64 {
			break
		}
		//fmt.Printf("minNext: %o, %x\n", minNext, minNext.src)
		//fmt.Printf("nexts=: %v\n", nexts)
		// we could have a draw so need to check each next and change state accordingly
		// set current state so we only update nexts on the diff
		prevHeads = nil
		for _, h := range heads {
			prevHeads = append(prevHeads, h)
		}
		for _, winner := range nexts {
			if winner.v == minNext.v {
				risk[winner.i][winner.j] = winner.v
				//setVisited(winner.i, winner.j)
				setVisited(winner.i, winner.j)
				//sptSet[winner.i][winner.j] = true
				heads = append(heads, winner)
				//removeNext(winner.i, winner.j)
			}
		}

		//fmt.Printf("len(nexts) %d\n", len(nexts))

		// we should remove heads that no longer need to be checked
		// these have all neighbours in spt

		reduceHeads()

		//for _, r := range risk {
		//	fmt.Printf("risk =: %v\n", r)
		//}
		//for _, s := range sptSet {
		//	fmt.Printf("spt =: %v\n", s)
		//}
		// clear any in spt from nexts
		for _, n := range nexts {
			if sptSet[n.i][n.j] {
				removeNext(n.i, n.j)
			}
		}
		//fmt.Printf("nexts=: %v\n", nexts)
		//nexts = nil

		if risk[len(risk)-1][len(risk[0])-1] > 0 {
			break
		}

	}
	fmt.Printf("TotalRisk is %d", risk[len(risk)-1][len(risk[0])-1])

}

func nextExists(i int, j int, v int, p *Point) bool {
	for _, n := range nexts {
		if n.i == i && n.j == j && n.v == v && n.src == p {
			return true
		}
	}
	return false

}
func setVisited(i int, j int) {
	sptSet[i][j] = true
	//removeHead(i, j)
}
func countVisited() int {
	var count int = 0
	for _, r := range sptSet {
		for _, n := range r {
			if n {
				count++
			}
		}
	}
	return count
}
func reduceHeads() {
	for _, h := range heads {
		//fmt.Printf("checkign head to remove:%o\n", h)
		remove := true
		if checkBoundsBool(sptSet, h.i-1, h.j) && !sptSet[h.i-1][h.j] {
			remove = false
		}
		if checkBoundsBool(sptSet, h.i+1, h.j) && !sptSet[h.i+1][h.j] {
			remove = false
		}
		if checkBoundsBool(sptSet, h.i, h.j-1) && !sptSet[h.i][h.j-1] {
			remove = false
		}
		if checkBoundsBool(sptSet, h.i, h.j+1) && !sptSet[h.i][h.j+1] {
			remove = false
		}
		if remove {
			removeHead(h.i, h.j)
		}

	}
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
	// remove nexts that have this head set as src
	for _, n := range nexts {
		if (*n.src).i == heads[removeIndex].i && (*n.src).j == heads[removeIndex].j {
			removeNext(n.i, n.j)
		}
	}
	heads = append(heads[0:removeIndex], heads[removeIndex+1:len(heads)]...)
}

func removeNext(i int, j int) {
	// find position
	removeIndex := -1
	for hi, p := range nexts {
		if p.i == i && p.j == j {
			removeIndex = hi
		}
	}
	if removeIndex == -1 {
		return
	}
	nexts = append(nexts[0:removeIndex], nexts[removeIndex+1:len(nexts)]...)
}
func checkBounds(space [][]int, i int, j int) bool {
	return i >= 0 && j >= 0 && i < len(space) && j < len(space[0])
}

func checkBoundsBool(space [][]bool, i int, j int) bool {
	return i >= 0 && j >= 0 && i < len(space) && j < len(space[0])
}
func getVal(space [][]int, i int, j int) int {
	if !checkBounds(space, i, j) {
		return math.MaxInt64
	}
	return space[i][j]
}
