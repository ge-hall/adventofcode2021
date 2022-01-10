package AoC

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"math"
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
		versionSum = processPacket(dst, &bitPointer)

		fmt.Printf("result = %d", versionSum)

	}

}
func processPacket(dst []byte, bitPtr *int) int {
	//var packet int
	sum := 0
	version := getNextBits(dst, 3, *bitPtr)
	*bitPtr += 3
	fmt.Printf("reading version %d packet.\n", version)

	typeID := getNextBits(dst, 3, *bitPtr)
	*bitPtr += 3
	fmt.Printf("TypeID:%d\n", typeID)

	if typeID == 4 {
		// rules:
		// read groups until msb = 0
		var val int = 0
		group := getNextBits(dst, 5, *bitPtr)
		*bitPtr += 5
		val = int(group & 0b01111)
		var msb = group & 10000 >> 4
		fmt.Printf(":group:%b, val %b\n", group, val)
		for msb == 1 {
			group := getNextBits(dst, 5, *bitPtr)
			*bitPtr += 5
			msb = group & 10000 >> 4
			val = val << 4

			//fmt.Printf(":group:%b,group:%b, val %012b\n", group, group&0b01111, val)
			val = val | int(group&0b01111)
			//fmt.Printf(":group:%b, val %012b\n", group, val)
		}
		fmt.Printf("Literal Value: %d(%08b)\n", val, val)
		sum = val

	} else {
		lengthType := getNextBits(dst, 1, *bitPtr)
		*bitPtr += 1
		fmt.Printf("Operator packet with lengthType:%b\n", lengthType)
		var packetCount int
		var packetLength int = 11

		if lengthType == 0 {
			packetLength = getNextBits(dst, 15, *bitPtr)
			*bitPtr += 15
			var val = 0
			var minVal = math.MaxInt64
			var maxVal = 0
			var cval int = 0
			var pval int = -1
			fmt.Printf(":packetLength:%d\n\n\n", packetLength)
			for remaining := packetLength; remaining > 0; {
				curPtr := *bitPtr
				if typeID == 0 {
					val += processPacket(dst, bitPtr)
				} else if typeID == 1 {
					if val == 0 {
						val = 1
					}
					val *= processPacket(dst, bitPtr)
				} else if typeID == 2 {
					cval = processPacket(dst, bitPtr)
					if cval < minVal {
						minVal = cval
					}
				} else if typeID == 3 {
					cval = processPacket(dst, bitPtr)
					if cval > maxVal {
						maxVal = cval
					}

				} else if typeID == 5 {
					cval = processPacket(dst, bitPtr)
					if cval < pval {
						val = 1
					}
					pval = cval
				} else if typeID == 6 {
					cval = processPacket(dst, bitPtr)
					if cval > pval {
						val = 1
					}
					pval = cval
				} else if typeID == 7 {
					cval = processPacket(dst, bitPtr)
					if cval == pval {
						val = 1
					}
					pval = cval
				}
				remaining -= *bitPtr - curPtr
			}
			if typeID == 2 {
				val = minVal
			}
			if typeID == 3 {
				val = maxVal
			}
			sum = val

		} else {
			packetCount = getNextBits(dst, 11, *bitPtr)
			*bitPtr += 11
			fmt.Printf(":Operator Packet contains packetCount:%d\n\n\n", packetCount)

			var val = 0
			var minVal = math.MaxInt64
			var maxVal = 0
			var cval int = 0
			var pval int = -1
			// packet count only applies to lengthType 1
			for i := 0; i < int(packetCount); i++ {

				if typeID == 0 {
					val += processPacket(dst, bitPtr)
				} else if typeID == 1 {
					if val == 0 {
						val = 1
					}
					val *= processPacket(dst, bitPtr)
				} else if typeID == 2 {
					cval = processPacket(dst, bitPtr)
					if cval < minVal {
						minVal = cval
					}
				} else if typeID == 3 {
					cval = processPacket(dst, bitPtr)
					if cval > maxVal {
						maxVal = cval
					}

				} else if typeID == 5 {
					cval = processPacket(dst, bitPtr)
					if cval < pval {
						val = 1
					}
					pval = cval
				} else if typeID == 6 {
					cval = processPacket(dst, bitPtr)
					if cval > pval {
						val = 1
					}
					pval = cval
				} else if typeID == 7 {
					cval = processPacket(dst, bitPtr)
					if cval == pval {
						val = 1
					}
					pval = cval
				}
			}
			if typeID == 2 {
				val = minVal
			}
			if typeID == 3 {
				val = maxVal
			}
			sum = val

		}

	}
	return sum
}
