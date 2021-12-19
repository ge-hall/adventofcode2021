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

		cs.AddCave(caves[1])
		// add connection
		cs.ConnectCave(cs.AddCave(caves[0]), cs.AddCave(caves[1]))

	}

	cs.String()
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
		s += cs.caves[i].String() + " -> "
		near := cs.connections[*cs.caves[i]]
		for j := 0; j < len(near); j++ {
			s += near[j].String() + " "
		}
		s += "\n"
	}
	fmt.Println(s)

}
