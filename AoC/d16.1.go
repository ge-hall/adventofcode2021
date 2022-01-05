package AoC

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func D16_1() {
	dat, err := ioutil.ReadFile("inputd16")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	var data = strings.Split(string(dat), "\n")
	// read data into bit array converting hex into binary
	for _, line := range data {
		fmt.Printf("%s\n", line)
		dst := make([]byte, hex.DecodedLen(len(line)))
		hex.Decode(dst, []byte(line))
		// dst is byte array
		dst = append(dst, uint8(0))
		b := binary.BigEndian.Uint32(dst[0:])

		fmt.Printf("%b\n", b)
		var version uint8 = uint8(b >> 29)
		var typeID uint8 = uint8(b << 3 >> 29)
		var aGroup uint8 = uint8(b << 5 >> 26)
		var bGroup uint8 = uint8(b << 11 >> 27)
		var cGroup uint8 = uint8(b << 16 >> 27)
		fmt.Printf("VVV=%b\n", version)
		fmt.Printf("TTT=%b\n", typeID)
		fmt.Printf("AAAAA=%08b\n", aGroup)
		fmt.Printf("BBBBB=%08b\n", bGroup)
		fmt.Printf("CCCCC=%08b\n", cGroup)
		// need to construct binary value of 15 bits by building 5 bits at a time and shift left
		// or bitwise & with mask
		//v := binary.BigEndian.Uint32([]byte{0, aGroup, bGroup, cGroup})o
		var v1 uint16 = uint16(aGroup) << 9
		var v2 uint16 = uint16(bGroup) << 4
		var v uint16 = v1 | v2 | uint16(cGroup)
		if typeID == 4 {
			fmt.Printf("Type 4 literal value composed of last 4 bits of each group\n")
			v1 = uint16(aGroup<<4) << 4
			v2 = uint16(bGroup << 4)
			v = v1 | v2 | uint16(cGroup)
		}

		fmt.Printf("v1 value= %015b => %d\n", v1, v1)
		fmt.Printf("v2 value= %015b => %d\n", v2, v2)
		fmt.Printf("value= %015b => %d\n", v, v)

	}
	// implement Dijkstra

}
