package AoC

import (
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
		readPacket(dst, bitPointer)

		fmt.Printf("result = %d", versionSum)

	}

}
func readPacket(dst []byte, bitPtr int) int {
	var packet int
	versionSum := 0
	version := getNextBits(dst, 3, bitPtr)
	if version == 0 {
		return 0
	} // extra bits
	versionSum += int(version)
	bitPtr += 3
	fmt.Printf("reading version %d packet.\n", version)

	typeID := getNextBits(dst, 3, bitPtr)
	bitPtr += 3
	fmt.Printf("TTT:%03b\n", typeID)

	if typeID == 4 {
		for i := 0; i < 3; i++ {
			packet = getNextBits(dst, 5, bitPtr)
			bitPtr += 5
			fmt.Printf(":packet:%b\n", packet)
		}

	} else {
		lengthType := getNextBits(dst, 1, bitPtr)
		bitPtr += 1
		fmt.Printf("lengthType:%b\n", lengthType)
		var packetCount int
		var packetLength int = 11

		if lengthType == 0 {
			packetLength = getNextBits(dst, 15, bitPtr)
			bitPtr += 15
			fmt.Printf(":packetLength:%d\n\n\n", packetLength)
			// read 11 bit chunks then remainder
			for packetBits := 0; packetBits < packetLength; packetBits += 11 {
				if packetLength-packetBits > 22 {
					packet = getNextBits(dst, 11, bitPtr)
					fmt.Printf(":--packet:%d\n", packet)
					bitPtr += 11
				} else {

					packet = getNextBits(dst, packetLength-packetBits, bitPtr)
					bitPtr += packetLength - packetBits
					fmt.Printf(":--packet:%d\n", packet)
					break
				}
			}
		} else {
			packetCount = getNextBits(dst, 11, bitPtr)
			bitPtr += 11
			fmt.Printf(":packetCount:%d\n\n\n", packetCount)
			// packet count only applies to lengthType 1
			for i := 0; i < int(packetCount); i++ {
				packet = getNextBits(dst, 11, bitPtr)
				bitPtr += 11
				fmt.Printf(":packet:%d\n", packet)
			}
		}

	}
	return versionSum
}
func getNextBits(bytes []byte, bits int, bitPtr int) int {
	fmt.Printf("%v, %d, %d[%08b]%d\n", bytes, bits, bitPtr/8, bytes[bitPtr/8], bytes[bitPtr/8])
	var result int
	// for each bit
	for b := 0; b < bits; b++ {
		// get currentByte and shift right
		currByte := bytes[bitPtr/8] >> (7 - bitPtr%8)
		//>> (8 - bitPtr%8)
		fmt.Printf("currByte:%08b - %d\n", currByte, 7-bitPtr%8)
		result = result << 1
		result = result | int(currByte&1)

		bitPtr++

	}
	return result

}
