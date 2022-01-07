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
		numBytes, _ := hex.Decode(dst, []byte(line))
		fmt.Printf("length of array is %d\n", numBytes)
		fmt.Printf("array is %v\n", dst)
		// dst is byte array
		if len(dst) < 4 {
			dst = append(dst, uint8(0))
		}
		var bitPointer = 0
		var bytePointer = 0
		word := binary.BigEndian.Uint32(dst[bytePointer : bytePointer+4])
		// get 3 bits
		fmt.Printf("AAA:%03b\n", word)

		// parse packets
		versionSum := 0
		for {
			version := getNextBits(dst, 3, bitPointer)
			versionSum += int(version)
			bitPointer += 3
			fmt.Printf("VVV:%03b\n", version)
			typeID := getNextBits(dst, 3, bitPointer)
			bitPointer += 3
			fmt.Printf("TTT:%03b\n", typeID)
			if typeID == 4 {

			} else {
				lengthType := getNextBits(dst, 1, bitPointer)
				bitPointer += 1
				fmt.Printf("lengthType:%b\n", lengthType)
				var packetCount uint16
				if lengthType == 0 {
					packetCount = getNextBits(dst, 15, bitPointer)
					bitPointer += 15
					fmt.Printf(":packetCount:%b\n", packetCount)
				} else {
					packetCount = getNextBits(dst, 11, bitPointer)
					bitPointer += 11
					fmt.Printf(":packetCount:%b\n", packetCount)
				}
				for i := 0; i < int(packetCount); i++ {
					packet := getNextBits(dst, 11, bitPointer)
					bitPointer += 11
					fmt.Printf(":packet:%b\n", packet)
				}

			}
			if len(dst)*8-bitPointer < 8 {
				break
			}

		}
		fmt.Printf("result = %d", versionSum)

	}

}

func getNextBits(bytes []byte, bits int, bitPtr int) uint16 {
	fmt.Printf("%v, %d, %d[%08b]\n", bytes, bits, bitPtr, bytes[bitPtr/8])
	var result uint16
	// for each bit
	for b := 0; b < bits; b++ {
		// get currentByte and shift right
		currByte := bytes[bitPtr/8] >> (7 - bitPtr%8)
		//>> (8 - bitPtr%8)
		fmt.Printf("currByte:%08b - %d\n", currByte, 7-bitPtr%8)
		result = result << 1
		result = result | uint16(currByte&1)

		bitPtr++

	}
	return result

}
