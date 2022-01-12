package AoC

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type Bounds struct {
	x1 int64
	y1 int64
	x2 int64
	y2 int64
}

func D17_1() {
	dat, err := ioutil.ReadFile("inputd17")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	var data = strings.Split(string(dat), "\n")
	// read data into bit array converting hex into binary
	// target area: x=20..30, y=-10..-5
	for _, line := range data {
		fmt.Printf("%s\n", line)
		tokens := strings.Split(line, " ")
		xDim := tokens[2]
		yDim := tokens[3]
		fmt.Printf("x:%s\n", strings.ReplaceAll(strings.Split(xDim, "..")[0], "x=", ""))
		fmt.Printf("x2:%s\n", strings.ReplaceAll(strings.Split(xDim, "..")[1], ",", ""))
		fmt.Printf("y:%s\n", strings.ReplaceAll(strings.Split(yDim, "..")[0], "y=", ""))
		fmt.Printf("y2:%s\n", strings.ReplaceAll(strings.Split(yDim, "..")[1], ",", ""))
		x1, _ := strconv.ParseInt(strings.ReplaceAll(strings.Split(xDim, "..")[0], "x=", ""), 10, 32)
		x2, _ := strconv.ParseInt(strings.ReplaceAll(strings.Split(xDim, "..")[1], ",", ""), 10, 32)
		y1, _ := strconv.ParseInt(strings.ReplaceAll(strings.Split(yDim, "..")[0], "y=", ""), 10, 32)
		y2, _ := strconv.ParseInt(strings.ReplaceAll(strings.Split(yDim, "..")[1], ",", ""), 10, 32)
		bounds := Bounds{x1, y1, x2, y2}
		fmt.Printf("x1,y1[%d,%d] x2,y2[%d,%d]\n", x1, y1, x2, y2)

		stepCount := 1000

		var x int64 = 0
		var y int64 = 0
		var xVelocity int64 = 0
		var yVelocity int64 = 0
		var highest int64 = 0
		var hitCount int64 = 0
		for j := -50; j < 500; j++ {
			for k := -2000; k < 2000; k++ {
				var maxY int64 = 0
				xVelocity = int64(j)
				yVelocity = int64(k)
				x = 0
				y = 0
				fmt.Printf("Check %d %d\n", xVelocity, yVelocity)
				for i := 0; i < stepCount; i++ {
					x += xVelocity
					if xVelocity > 0 {
						xVelocity -= 1
					}
					if xVelocity < 0 {
						xVelocity += 1
					}
					y += yVelocity
					yVelocity--
					//fmt.Printf("Velocity %d,%d\n", xVelocity, yVelocity)
					withinBounds := checkTargetBounds(x, y, bounds)
					if withinBounds < 1 {
						if maxY < y {
							maxY = y
						}
					}
					if withinBounds == 0 {
						if maxY > highest {
							highest = maxY
						}
						hitCount++
						fmt.Printf("hit target with %d %d %d\n\n", j, k, maxY)
						break
					}
					if withinBounds == 1 {
						break
					}

				}
			}
		}

		fmt.Printf("Result:%d %d\n", highest, hitCount)
	}

}

func checkTargetBounds(x int64, y int64, b Bounds) int {
	//fmt.Printf("checking x%d, y%d bounds %d %d %d %d\n", x, y, b.x1, b.x2, b.y1, b.y2)
	if x >= b.x1 && x <= b.x2 && y >= b.y1 && y <= b.y2 {
		print("within bounds\n")
		return 0
	}
	if x > b.x2 || y < b.y1 {
		//print("past target\n")
		return 1
	}
	return -1
}

//.............#....#............
//.......#..............#........
//...............................
//S........................#.....
//...............................
//...............................
//...........................#...
//...............................
//....................TTTTTTTTTTT
//....................TTTTTTTTTTT
//....................TTTTTTTT#TT
//....................TTTTTTTTTTT
//....................TTTTTTTTTTT
//....................TTTTTTTTTTT
