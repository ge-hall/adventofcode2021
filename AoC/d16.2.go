package AoC

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func D16_2() {
	dat, err := ioutil.ReadFile("inputd16")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	var data = strings.Split(string(dat), "\n")
	// read data into bit array converting hex into binary
	for _, line := range data {
		fmt.Printf("%s\n", line)
		dst := make([]byte, hex.DecodedLen(len(line)))
		numBytes, _ := hex.Decode(dst, []byte(line))
		fmt.Printf("length of array is %d\n", numBytes)
		fmt.Printf("array is %v\n\n", dst)
		// dst is byte array
		if len(dst) < 4 {
			dst = append(dst, uint8(0))
		}
		var bitPointer = 0

		// parse packets
		versionSum := 0
		versionSum = processPacket(dst, &bitPointer, 0)

		fmt.Printf("result = %d", versionSum)

	}

}
func processPacket(dst []byte, bitPtr *int, length int) int {
	//var packet int
	versionSum := 0
	version := getNextBits(dst, 3, *bitPtr)
	//if version == 0 {
	//	return 0
	//} // extra bits
	versionSum += int(version)
	*bitPtr += 3
	length -= 3
	fmt.Printf("reading version %d packet.\n", version)

	typeID := getNextBits(dst, 3, *bitPtr)
	*bitPtr += 3
	length -= 3
	fmt.Printf("TypeID:%d\n", typeID)

	if typeID == 4 {
		// rules:
		// read groups until msb = 0
		var msb = 1
		for msb == 1 {
			group := getNextBits(dst, 5, *bitPtr)
			*bitPtr += 5
			msb = group & 10000 >> 4
			fmt.Printf(":group:%b\n", group)
		}

	} else {
		lengthType := getNextBits(dst, 1, *bitPtr)
		*bitPtr += 1
		fmt.Printf("Operator packet with lengthType:%b\n", lengthType)
		var packetCount int
		var packetLength int = 11

		if lengthType == 0 {
			packetLength = getNextBits(dst, 15, *bitPtr)
			*bitPtr += 15
			fmt.Printf(":packetLength:%d\n\n\n", packetLength)
			for remaining := packetLength; remaining > 0; {
				curPtr := *bitPtr
				versionSum += processPacket(dst, bitPtr, packetLength)
				remaining -= *bitPtr - curPtr
			}

		} else {
			packetCount = getNextBits(dst, 11, *bitPtr)
			*bitPtr += 11
			fmt.Printf(":Operator Packet contains packetCount:%d\n\n\n", packetCount)
			// packet count only applies to lengthType 1
			for i := 0; i < int(packetCount); i++ {
				versionSum += processPacket(dst, bitPtr, 0)
			}
		}

	}
	return versionSum
}
