package AoC

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func D12_1() {
	dat, err := ioutil.ReadFile("inputd12")
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

	}
	print("===============\n\n")
	cs.String()
	print("===============\n\n")
	cs.Traverse()
}

type CaveSystem struct {
	caves       []*Cave
	connections map[Cave][]*Cave
}
type Cave struct {
	label    string
	big      bool
	entrance bool
	exit     bool
}

func (c *Cave) String() string {
	return fmt.Sprint(c.label)
}
func (cs *CaveSystem) AddCave(label string) *Cave {
	// if already exists return
	for i := 0; i < len(cs.caves); i++ {
		if cs.caves[i].label == label {
			return cs.caves[i]
		}
	}
	cave := &Cave{label, strings.ToUpper(label) == label, label == "start", label == "end"}
	cs.caves = append(cs.caves, cave)
	return cave
}
func (cs *CaveSystem) ConnectCave(c1, c2 *Cave) {
	if cs.connections == nil {
		cs.connections = make(map[Cave][]*Cave)
	}
	cs.connections[*c1] = append(cs.connections[*c1], c2)
}

func (cs *CaveSystem) String() {
	s := ""
	for i := 0; i < len(cs.caves); i++ {
		fmt.Printf("cave:%o\n", cs.caves[i])
		if !cs.caves[i].entrance {
			continue
		}
		s += cs.caves[i].String() + " -> "
		near := cs.connections[*cs.caves[i]]
		for j := 0; j < len(near); j++ {
			s += near[j].String() + " "
		}
		s += "\n"
	}
	fmt.Println(s)

}
func (cs *CaveSystem) Traverse() {
	// get Starts
	for i := 0; i < len(cs.caves); i++ {

		if !cs.caves[i].entrance {
			continue
		}
		var path []string
		cs.FindPaths(*cs.caves[i], path)

	}
}
func (cs *CaveSystem) FindPaths(cave Cave, path []string) {
	path = append(path, cave.label)
	if cave.exit {

		fmt.Printf(",%s", path)
		return
	}

	for c := 0; c < len(cs.connections[cave]); c++ {
		cs.FindPaths(*cs.connections[cave][c], path)
	}
}
